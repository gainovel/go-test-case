/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024/1/10 17:38:02 星期三
 * @ProductName  : GoLand
 * @PrjectName   : go-examples
 * @File         : examples/stdlib/runtime/string/features_usages_001_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package string

import (
	"fmt"
	commonconsts "gainovel.com/go/testcase/consts/common"
	commontools "gainovel.com/go/testcase/tools/common"
	"testing"
)

func TestName_2024_01_10_17_38_02(t *testing.T) {
	// go test -run TestName_2024_01_10_17_38_02/string_1
	t.Run("string 1", func(t *testing.T) {
		var (
			str1  string
			bytes []byte
			runes []rune
		)
		str1 = "hello 中国！"
		// ['h','e','l','l','o',   ' ',         '中',          '国',           '！'   ]
		// ['h','e','l','l','o',   ' ',         '中',          '国',           '！'   ]
		// [104,101,108,108,111,    32,     228,184,173    229,155,189    239,188,129] []byte len:15
		// [104,101,108,108,111,    32,        20013,        22269,          65281   ] []rune len:9
		//fmt.Println(len(str1))
		//fmt.Println(len([]rune(str1)))

		commontools.PrintDownHand("str1")
		commontools.Print88HorizontalLine()
		commontools.Println(true, "str1", "hello 中国！")
		commontools.Print88HorizontalLine()

		// []byte
		commontools.PrintDownHand("[]byte(str1)")
		commontools.Print88HorizontalLine()
		fmt.Printf("%sfor (str1): ", commonconsts.VERTICAL1SPACE1)
		for i := 0; i < len(str1); i++ {
			fmt.Printf("%d ", str1[i])
		}
		fmt.Println()

		fmt.Printf("%sfor []byte(str1): ", commonconsts.VERTICAL1SPACE1)
		bytes = []byte(str1)
		for i := 0; i < len(bytes); i++ {
			fmt.Printf("%d ", bytes[i])
		}
		fmt.Println()
		commontools.Print88HorizontalLine()

		// []rune
		commontools.PrintDownHand("[]rune(str1)")
		commontools.Print88HorizontalLine()
		fmt.Printf("%sfor-range str1: ", commonconsts.VERTICAL1SPACE1)
		for _, v := range str1 {
			fmt.Printf("%d ", v)
		}
		fmt.Println()

		fmt.Printf("%sfor []rune(str1): ", commonconsts.VERTICAL1SPACE1)
		runes = []rune(str1)
		for i := 0; i < len(runes); i++ {
			fmt.Printf("%d ", runes[i])
		}
		fmt.Println()
		commontools.Print88HorizontalLine()

	})
	// go test -run TestName_2024_01_10_17_38_02/string_2
	t.Run("string 2", func(t *testing.T) {
		var (
			str1  string
			bytes []byte
			runes []rune
		)
		str1 = "hello 中国！"
		// ['h','e','l','l','o',   ' ',         '中',          '国',           '！'   ]
		// ['h','e','l','l','o',   ' ',         '中',          '国',           '！'   ]
		// [104,101,108,108,111,    32,     228,184,173    229,155,189    239,188,129] []byte len:15
		// [104,101,108,108,111,    32,        20013,        22269,          65281   ] []rune len:9
		//fmt.Println(len(str1))
		//fmt.Println(len([]rune(str1)))

		commontools.PrintDownHand("str1")
		commontools.Print88HorizontalLine()
		commontools.Println(true, "str1", "hello 中国！", "len(str1)", len(str1))
		commontools.Print88HorizontalLine()

		// []byte
		bytes = []byte(str1)
		commontools.PrintDownHand("[]byte(str1)")
		commontools.Print88HorizontalLine()
		commontools.Println(true, "[]byte(str1)", bytes, "len(bytes)", len(bytes))
		commontools.Print88HorizontalLine()

		// []rune
		runes = []rune(str1)
		commontools.PrintDownHand("[]rune(str1)")
		commontools.Print88HorizontalLine()
		commontools.Println(true, "[]rune(str1)", runes, "len(runes)", len(runes))
		commontools.Print88HorizontalLine()

	})
	t.Run("string 3", func(t *testing.T) {
		var (
			str1  string
			bytes []byte
			runes []rune
		)
		str1 = "hello 中国！"
		// ['h','e','l','l','o',   ' ',         '中',          '国',           '！'   ]
		// ['h','e','l','l','o',   ' ',         '中',          '国',           '！'   ]
		// [104,101,108,108,111,    32,     228,184,173    229,155,189    239,188,129] []byte len:15
		// [104,101,108,108,111,    32,        20013,        22269,          65281   ] []rune len:9
		//fmt.Println(len(str1))
		//fmt.Println(len([]rune(str1)))
		commontools.PrintAll(true, "str1", "origin string", "", "str1", str1, "len(str1)", len(str1))
		// []byte
		bytes = []byte(str1)
		commontools.PrintAll(true, "[]byte(str1)", "string to []byte", "string中的字符的存储形式就是[]byte", "[]byte(str1)", bytes, "len(bytes)", len(bytes))
		// []rune
		runes = []rune(str1)
		commontools.PrintAll(true, "[]rune(str1)", "string to []rune", "", "[]rune(str1)", runes, "len(rune)", len(runes))
	})
	// go test -run TestName_2024_01_10_17_38_02/string_4
	t.Run("string 4", func(t *testing.T) {
		var (
			str1  string
			bytes []byte
			runes []rune
		)
		str1 = "hello 中国！"
		// ['h','e','l','l','o',   ' ',         '中',          '国',           '！'   ]
		// ['h','e','l','l','o',   ' ',         '中',          '国',           '！'   ]
		// [104,101,108,108,111,    32,     228,184,173    229,155,189    239,188,129] []byte len:15
		// [104,101,108,108,111,    32,        20013,        22269,          65281   ] []rune len:9
		// fmt.Println(len(str1))
		// fmt.Println(len([]rune(str1)))
		bytes = []byte(str1)
		runes = []rune(str1)
		myfmt.ColorDescPrintln("字符串底层数组是[]byte类型")
		fmt.Println()
		for i := 0; i < len(str1); i++ {

		}
		myfmt.ColorDescPrintln("str1")
		myfmt.DumpPrintln(str1)
		myfmt.ColorDescPrintln("string👉[]byte")
		myfmt.ColorDescPrintln("[]byte(str1)")
		myfmt.DumpPrintln(bytes)
		myfmt.ColorDescPrintln("string👉[]rune")
		myfmt.ColorDescPrintln("[]rune(str1)")
		myfmt.DumpPrintln(runes)
	})

}
