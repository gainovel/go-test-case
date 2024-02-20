/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-01-17 16:41:27 星期三
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : thirdparty/golang.org/x/crypto/crypto_quickstart_001_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package crypto000

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestName_2024_01_17_16_41_27(t *testing.T) {
	// go test -run TestName_2024_01_17_16_41_27/case1
	t.Run("case1", func(t *testing.T) {
		pwd := []byte("123456")
		encrypted, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
		myfmt.ColorDescPrintln("encrypt加密\"123456\"结果：")
		myfmt.KeyValuePrintln(string(pwd), string(encrypted))
		err = bcrypt.CompareHashAndPassword(encrypted, pwd)
		require.NoError(t, err)
	})
}
