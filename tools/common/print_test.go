/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024/1/9 19:06:13 星期二
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : examples/tools/common/print_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package commontools

import (
	"fmt"
	"testing"
)

func TestName_2024_01_09_19_06_13(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		var (
			name  string
			name1 string
			int1  int
			str1  string
			intS1 []int
			strS1 []string
		)
		name = "case1"
		name1 = "case1"
		int1 = 1
		str1 = "hello"
		intS1 = []int{1, 2, 3}
		strS1 = []string{"hello", "world", "example"}
		fmt.Println(intS1)
		Println(true, name, name1, int1, str1, name1, strS1)
	})
	// go test -run TestName_2024_01_09_19_06_13/case2
	t.Run("case2", func(t *testing.T) {
		var (
			name  string
			name1 string
			int1  int
			str1  string
			strS1 []string
		)
		name = "case1"
		name1 = "case1"
		int1 = 1
		str1 = "hello"
		strS1 = []string{"hello", "world", "example"}
		PrintAll(true, "testDown", name, name1, int1, str1, name1, strS1)
	})
}
