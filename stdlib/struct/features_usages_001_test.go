/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-02-20 12:18:40 星期二
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : stdlib/struct/features_usages_001_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package struct000

import (
	"fmt"
	"testing"
)

type A struct {
	name string
}

func (a A) Name() string {
	a.name = "Hi! " + a.name
	return a.name
}

func TestName_2024_02_20_12_18_40(t *testing.T) {
	// 方法即method，Go语言支持为自定义类型实现方法，method在具体实现上与普通的函数并无不同，只不过会通过运行时栈多传递一个隐含的参数，这个隐含的参数就是所谓的接收者。
	// 以下代码展示了两种不同的写法，都能顺利通过编译并正常运行，实际上这两种写法会生成同样的机器码。
	// 第一种：a.Name()，这是我们惯用的写法，很方便；
	// 第二种：A.Name(a)，这种写法更底层也更严谨，要求所有的类型必须严格对应，否则是无法通过编译的。
	// 其实编译器会帮我们把第一种转换为第二种的形式，所以我们惯用的第一种写法只是“语法糖”，方便而已。
	t.Run("case1", func(t *testing.T) {
		a := A{name: "eggo"}
		// 1）编译器的语法糖，提供面向对象的语法
		fmt.Println(a.Name())
		// 2）更贴近真实实现的写法，和普通函数调用几乎没什么不同
		fmt.Println(A.Name(a))
	})
}
