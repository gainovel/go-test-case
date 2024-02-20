/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-01-17 15:37:50 星期三
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : examples/thirdparty/github.com/gorm.io_gorm/user.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package gorm000

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"unique"`
	Age          uint8
	Birthday     time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
}

// TableName 指定表名
func (User) TableName() string {
	return "user"
}
