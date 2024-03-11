/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-02-21 16:29:27 星期三
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : thirdparty/github.com/github.com_rabbitmq_amqp091-g/features_usages_001_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package amqp091

import (
	"context"
	"log"
	"testing"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	amqp091tools "github.com/gainovel/testcase/thirdparty/github.com/github.com_rabbitmq_amqp091-g/tools"
)

var (
	url = "amqp://guest:guest@lenovoy7000:5672/"
)

func TestName_2024_02_21_16_29_27(t *testing.T) {
	t.Run("amqp.Dial", func(t *testing.T) {
		conn := amqp091tools.GetConn(url)
		defer conn.Close()
	})
	// go test -v -run TestName_2024_02_21_16_29_27/Hello_World_Producer
	t.Run("Hello World Producer", func(t *testing.T) {
		conn := amqp091tools.GetConn(url)
		defer conn.Close()
		ch := amqp091tools.GetChannel(conn)
		defer ch.Close()
		q, err := ch.QueueDeclare(
			"hello", // name
			false,   // durable
			false,   // delete when unused
			false,   // exclusive
			false,   // no-wait
			nil,     // arguments
		)
		amqp091tools.FailOnError(err, "Failed to declare a queue")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		body := "Hello World!"
		err = ch.PublishWithContext(ctx,
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		amqp091tools.FailOnError(err, "Failed to publish a message")
		log.Printf(" [x] Sent %s\n", body)
	})
	// go test -v -run TestName_2024_02_21_16_29_27/Hello_World_Consumer
	t.Run("Hello World Consumer", func(t *testing.T) {
		conn := amqp091tools.GetConn(url)
		defer conn.Close()
		ch := amqp091tools.GetChannel(conn)
		defer ch.Close()
		q, err := ch.QueueDeclare(
			"hello", // name
			false,   // durable
			false,   // delete when unused
			false,   // exclusive
			false,   // no-wait
			nil,     // arguments
		)
		amqp091tools.FailOnError(err, "Failed to declare a queue")
		msgs, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			true,   // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
		)
		amqp091tools.FailOnError(err, "Failed to register a consumer")

		var forever chan struct{}

		go func() {
			for d := range msgs {
				log.Printf("Received a message: %s", d.Body)
			}
		}()

		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
		<-forever
	})
}
