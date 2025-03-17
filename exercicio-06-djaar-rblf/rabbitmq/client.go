package rabbitmq

import (
	"encoding/json"
	"exercicio-06-djaar-rblf/shared"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Client(invocations int, a [][]int, b [][]int) {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672")

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		panic(err)
	}

	defer ch.Close()

	responseQueue, err := ch.QueueDeclare(
		"matrix-multiplier-response-queue",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	msgs, err := ch.Consume(
		responseQueue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	correlationIDs := make(map[string]bool)

	for i := 0; i < invocations; i++ {
		msgRequest := shared.Request{
			Operation: "Mul",
			A:         a,
			B:         b,
		}

		msgRequestBytes, err := json.Marshal(msgRequest)

		if err != nil {
			panic(err)
		}

		correlationId, err := shared.GenerateRandomString(32)

		if err != nil {
			panic(err)
		}

		correlationIDs[correlationId] = true

		startTime := time.Now()

		err = ch.Publish(
			"",
			"matrix-multiplier-request-queue",
			false,
			false,
			amqp.Publishing{
				ContentType:   "text/plain",
				CorrelationId: correlationId,
				ReplyTo:       responseQueue.Name,
				Body:          msgRequestBytes,
			},
		)

		if err != nil {
			panic(err)
		}

		m := <-msgs

		if _, ok := correlationIDs[m.CorrelationId]; ok {
			elapsedTime := float64(time.Since(startTime).Nanoseconds()) / 1000000
			shared.WriteRTTValue("/app/data/rabbitmq-results.txt", elapsedTime)

			msgResponse := shared.Reply{}
			err = json.Unmarshal(m.Body, &msgResponse)

			if err != nil {
				panic(err)
			}

			delete(correlationIDs, m.CorrelationId)
		}
	}
}
