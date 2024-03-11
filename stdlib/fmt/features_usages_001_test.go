/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-01-14 16:03:35 星期日
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : stdlib/fmt/features_usages_001_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package fmt

import (
	"fmt"
	"testing"

	commonprint "github.com/gainovel/testcase/tools/common/print"
)

var (
	myfmt = commonprint.MyFmt
)

func TestName_2024_01_14_16_03_35(t *testing.T) {
	// go test -run TestName_2024_01_14_16_03_35/MyFmt
	t.Run("MyFmt", func(t *testing.T) {
		var (
			s1 []string
		)
		s1 = []string{"hello", "world", "china"}
		myfmt.ColorDescPrintln("for-range迭代时，打印i,v的地址，看是否共享迭代变量：")
		fmt.Println()
		for i, v := range s1 {
			myfmt.ColorDescPrintln(fmt.Sprintf("第%2d次迭代", i+1))
			myfmt.KeyValuePrintln("i address", fmt.Sprintf("%p", &i), "v address", fmt.Sprintf("%p", &v))
		}
	})
	t.Run("", func(t *testing.T) {
		var (
			str1 any
		)
		str1 = "hello world"
		fmt.Println(str1.(int))
	})
	t.Run("", func(t *testing.T) {
		var (
			str2 any
		)
		fmt.Println(str2.(int))
		fmt.Println("last")
	})
}
