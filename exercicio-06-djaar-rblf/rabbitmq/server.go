package rabbitmq

import (
	"encoding/json"
	"exercicio-06-djaar-rblf/shared"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Server() {
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

	q, err := ch.QueueDeclare(
		"matrix-multiplier-request-queue",
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
		q.Name,
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

	for d := range msgs {
		msg := shared.Request{}
		err := json.Unmarshal(d.Body, &msg)

		if err != nil {
			panic(err)
		}

		response := shared.Reply{}
		err = new(MatrixService).Multiply(msg, &response)

		if err != nil {
			panic(err)
		}

		responseBytes, err := json.Marshal(response)

		if err != nil {
			panic(err)
		}

		err = ch.Publish(
			"",
			d.ReplyTo,
			false,
			false,
			amqp.Publishing{
				ContentType:   "text/plain",
				CorrelationId: d.CorrelationId,
				Body:          responseBytes,
			},
		)

		if err != nil {
			panic(err)
		}
	}
}
