/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024/1/9 15:47:37 星期二
 * @ProductName  : GoLand
 * @PrjectName   : go-examples
 * @File         : examples/thirdparty/github.com/github.com_gin-gonic_gin/gin_quickstart_001_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package gin000

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestName_2024_01_09_15_47_37(t *testing.T) {
	t.Run("server listen on 0.0.0.0:8080", func(t *testing.T) {
		r := gin.Default()
		//r.Use(cors.New(cors.Config{
		//
		//	AllowCredentials: true,
		//	AllowHeaders:     []string{"Content-Type"},
		//	AllowOriginFunc: func(origin string) bool {
		//		if strings.HasPrefix(origin, "http://localhost") {
		//			return true
		//		}
		//		return strings.Contains(origin, "your_company.com")
		//	},
		//}))
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		r.Run() // 监听并在 0.0.0.0:8080 上启动服务
		fmt.Println()
	})
}
