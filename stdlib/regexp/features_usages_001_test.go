/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-02-05 14:35:07 星期一
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : stdlib/regexp/features_usages_001_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package regexp000

import (
	"fmt"
	"regexp"
	"testing"
)

func TestName_2024_02_05_14_35_07(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		pattern := "go(lang)?"
		text := "I love golang programming language. Go is awesome!"

		matched, _ := regexp.MatchString(pattern, text)
		fmt.Println(matched)

		match := regexp.MustCompile(pattern)
		fmt.Println(match.FindString(text))

		matches := match.FindAllString(text, -1)
		fmt.Println(matches)

		replaced := match.ReplaceAllString(text, "Golang")
		fmt.Println(replaced)
	})
}
