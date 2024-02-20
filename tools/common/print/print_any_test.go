/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-01-14 16:10:49 星期日
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : examples/tools/common/print/print_any_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package commonprint

import (
	"fmt"
	"testing"
)

func TestName_2024_01_14_16_10_49(t *testing.T) {
	// go test -run TestName_2024_01_14_16_10_49/case1
	t.Run("case1", func(t *testing.T) {
		type Person struct {
			Name  string
			Age   int
			Email string
		}
		var (
			s1 []string
			m1 map[string]string
			p1 *Person
		)
		s1 = []string{"hello", "world", "sdfsa", "sdfsa", "sdfsa"}
		m1 = map[string]string{
			"sdfs":   "sdfsa",
			"sdfs2":  "sdfsa",
			"sdfs3":  "sdfsa",
			"sdfs4":  "sdfsa",
			"sdfs5":  "sdfsa",
			"sdfs6":  "sdfsa",
			"sdfs7":  "sdfsa",
			"sdfs8":  "sdfsa",
			"sdfs9":  "sdfsa",
			"sdfs10": "sdfsa",
		}
		p1 = &Person{
			Name:  "xiaomao",
			Age:   30,
			Email: "xxx@qq.com",
		}
		MyFmt.ColorDescPrintln("s1:")
		MyFmt.DumpPrintln(s1)
		MyFmt.ColorDescPrintln("m1:")
		MyFmt.DumpPrintln(m1)
		MyFmt.ColorDescPrintln("p1:")
		MyFmt.DumpPrintln(p1)
	})
	// go test -run TestName_2024_01_14_16_10_49/case1
	t.Run("case2", func(t *testing.T) {
		var (
			s1 []string
		)
		s1 = []string{"hello", "world", "china"}
		MyFmt.ColorDescPrintln("for-range迭代时，打印i,v的地址，看是否共享迭代变量：")
		fmt.Println()
		for i, v := range s1 {
			MyFmt.ColorDescPrintln(fmt.Sprintf("第%2d次迭代", i+1))
			MyFmt.KeyValuePrintln("i address", fmt.Sprintf("%p", &i), "v address", fmt.Sprintf("%p", &v))
		}
	})
}
