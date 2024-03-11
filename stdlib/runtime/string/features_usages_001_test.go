/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024/1/10 17:38:02 星期三
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : stdlib/runtime/string/features_usages_001_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package string

import (
	"testing"

	commonprint "github.com/gainovel/testcase/tools/common/print"
)

var (
	myfmt = commonprint.MyFmt
)

func TestName_2024_01_10_17_38_02(t *testing.T) {
	// go test -v -run TestName_2024_01_10_17_38_02/string_to_bytes_or_runes
	t.Run("string to bytes or runes", func(t *testing.T) {
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
		myfmt.VarInitPrintln(`var (
	str1  string
	bytes []byte
	runes []rune
)`)
		bytes = []byte(str1)
		runes = []rune(str1)
		myfmt.ColorDescPrintln("bytes = []byte(str1)", "runes = []rune(str1)")
		myfmt.KeyValuePrintln("str1", str1, "bytes", bytes, "runes", runes)
	})
}
