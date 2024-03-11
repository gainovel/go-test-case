/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-01-15 11:16:48 星期一
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : stdlib/runtime/for/for_001_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package for000

import (
	"fmt"
	"testing"

	commonprint "github.com/gainovel/testcase/tools/common/print"
)

var (
	myfmt = commonprint.MyFmt
)

func TestName_2024_01_15_11_16_48(t *testing.T) {
	// go test -v -run TestName_2024_01_15_11_16_48/case1
	t.Run("case1", func(t *testing.T) {
		//	myfmt.DumpPrintln(`for i := 0; i < 3; i++ {
		//fmt.Printf("%p", &i)
		//}`)
		myfmt.ColorDescPrintln("打印for循环迭代索引变量的地址")
		for i := 0; i < 3; i++ {
			myfmt.ColorDescPrintln(fmt.Sprintf("第%2d次迭代", i+1))
			myfmt.KeyValuePrintln("i address", fmt.Sprintf("%p", &i))
		}
	})
}
