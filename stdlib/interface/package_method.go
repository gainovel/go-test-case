/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-02-20 00:16:44 星期二
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : stdlib/interface/package_method.go
 * @Version      : v0.1.0
 * @Description  : 包装方法探究
 **/

package interface000

type Ia interface {
	X() float64
}
type Point struct {
	x float64
}

func (p Point) X() float64 {
	return p.x
}

func (p *Point) SetX(x float64) {
	p.x = x
}
