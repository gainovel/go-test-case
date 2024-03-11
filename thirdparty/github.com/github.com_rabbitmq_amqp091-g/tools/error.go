/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-02-21 16:32:15 星期三
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : thirdparty/github.com/github.com_rabbitmq_amqp091-g/tools/error.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package amqp091tools

import (
	"log"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
