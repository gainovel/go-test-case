/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024/1/5 14:39:32 星期五
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : examples/tools/common/line.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package commontools

import (
	commonconsts "github.com/gainovel/testcase/consts/common"
)

func GenerateStarLine(n int) string {
	var (
		bytes []byte
	)
	bytes = make([]byte, n)
	for i := 0; i < n; i++ {
		bytes[i] = commonconsts.STAR
	}
	return string(bytes)
}
func Generate88StarLine() string {
	return GenerateStarLine(88)
}
func Generate88HorizontalLine() string {
	return GenerateHorizontalLine(88)
}
func GenerateHorizontalLine(n int) string {
	var (
		bytes []byte
	)
	bytes = make([]byte, n)
	for i := 0; i < n; i++ {
		bytes[i] = commonconsts.HORIZONTAL
	}
	return string(bytes)
}
func GenerateNullLine(n int) string {
	var (
		bytes []byte
	)
	bytes = make([]byte, n)
	for i := 0; i < n; i++ {
		bytes[i] = commonconsts.NEWLINE
	}
	return string(bytes)
}
