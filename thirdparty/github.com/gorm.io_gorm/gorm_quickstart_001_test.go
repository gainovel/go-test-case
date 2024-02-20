/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-01-17 15:19:18 星期三
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : examples/thirdparty/github.com/gorm.io_gorm/gorm_quickstart_001_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package gorm000

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"testing"
	"time"

	gormmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestName_2024_01_17_15_19_18(t *testing.T) {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123@tcp(192.168.0.109:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(gormmysql.Open(dsn), &gorm.Config{})
	t.Run("connecting_to_the_database", func(t *testing.T) {
		fmt.Println(db, err)
	})
	// go test -run TestName_2024_01_17_15_19_18/create_one
	t.Run("create one", func(t *testing.T) {
		db.AutoMigrate(&User{})
		user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
		result := db.Create(&user) // 通过数据的指针来创建
		myfmt.ColorDescPrintln("插入数据的返回值")
		// 返回插入数据的主键
		// 返回 error
		// 返回插入记录的条数
		myfmt.KeyValuePrintln("primary key", user.ID, "error", result.Error, "rows affected", result.RowsAffected)
	})
	t.Run("create two", func(t *testing.T) {
		users := []*User{
			&User{Name: "Jinzhu", Email: "xxx@qq.com", Age: 18, Birthday: time.Now()},
			&User{Name: "Jackson", Email: "xxx@qq.com", Age: 19, Birthday: time.Now()},
		}

		result := db.Create(users) // 传递切片以插入多行数据
		// 返回 error
		// 返回插入记录的条数
		err := result.Error
		if e, ok := err.(*mysql.MySQLError); ok {
			fmt.Println(e.Number)
		}
		myfmt.KeyValuePrintln("error", result.Error, "rows affected", result.RowsAffected)
	})
}
