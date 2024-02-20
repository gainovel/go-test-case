/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024/1/9 15:54:07 星期二
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : examples/tools/common/time_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package commontools

import (
	"fmt"
	"testing"
)

func TestName_2024_01_09_15_54_07(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		fmt.Println(GenerateCurFormatTime())
	})
}
