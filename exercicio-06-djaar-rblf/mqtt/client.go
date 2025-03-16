package mqtt

import (
	"encoding/json"
	"exercicio-06-djaar-rblf/shared"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Client(invocations int, a [][]int, b [][]int) {
	// Inicia conexão com o broker
	opts := mqtt.NewClientOptions().AddBroker("mqtt:3883")
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	defer client.Disconnect(1000)
	startTime := time.Now()
	for i := 0; i < invocations; i++ {
		startTime = time.Now()
		request := shared.Request{Operation: "mul", A: a, B: b, Client_id: opts.ClientID}

		// Enviar as matrizes para o broket
		payload, _ := json.Marshal(request)
		client.Publish("matrix/request", 1, false, payload)

		// Subescreve à resposta
		// Apenas as respostas com o nosso Client_id que interessam
		token := client.Subscribe("matrix/response", 1, func(client mqtt.Client, msg mqtt.Message) {
			var reply shared.Reply
			for {
				if err := json.Unmarshal(msg.Payload(), &reply); err != nil {
					log.Println("Erro ao decodificar resposta:", err)
					return
				}
				if reply.Client_id == opts.ClientID {
					break
				}
			}
			// Tempo em milisegundos mais preciso
			elapsedTime := float64(time.Since(startTime).Nanoseconds()) / 1000000
			shared.WriteRTTValue("/app/data/mqtt-results.txt", elapsedTime)
		})
		token.Wait()
	}
}
