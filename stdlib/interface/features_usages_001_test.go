/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-02-20 00:13:58 星期二
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : stdlib/interface/features_usages_001_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package interface000

import (
	"fmt"
	"testing"
)

func TestName_2024_02_20_00_22_32(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		var (
			ia Ia
			p1 *Point
			p2 Point
		)
		p1 = &Point{x: 1.1}
		p2 = Point{x: 1.2}
		ia = p1
		fmt.Println(ia.X())
		ia = p2
		fmt.Println(ia.X())
	})
}
