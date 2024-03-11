/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-02-21 16:40:23 星期三
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : thirdparty/github.com/github.com_rabbitmq_amqp091-g/tools/conn.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package amqp091tools

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func GetConn(url string) *amqp.Connection {
	var (
		conn *amqp.Connection
		err  error
	)
	conn, err = amqp.Dial(url)
	FailOnError(err, "Failed to connect to RabbitMQ")
	return conn
}

func GetChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	return ch
}

func CloseConn(conn *amqp.Connection) {
	conn.Close()
}

func CloseChannel(ch *amqp.Channel) {
	ch.Close()
}
