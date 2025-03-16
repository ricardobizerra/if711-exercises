package mqtt

import (
	"encoding/json"
	"exercicio-06-djaar-rblf/matrix"
	"exercicio-06-djaar-rblf/shared"
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Server() {
	opts := mqtt.NewClientOptions().AddBroker("mqtt:3883")
	opts.SetClientID("server")
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	// Subscribe ao tópico de requisição
	client.Subscribe("matrix/request", 1, func(client mqtt.Client, msg mqtt.Message) {
		var request shared.Request
		if err := json.Unmarshal(msg.Payload(), &request); err != nil {
			log.Println("Erro ao decodificar JSON:", err)
			return
		}

		result := matrix.Multiply(request.A, request.B)

		// Publica o resultado no tópico de resposta
		resultJSON, _ := json.Marshal(shared.Reply{R: result, Client_id: request.Client_id})
		client.Publish("matrix/response", 1, false, resultJSON)
	})

	fmt.Println("Servidor MQTT rodando...")
	select {} // Mantém o programa rodando
}
