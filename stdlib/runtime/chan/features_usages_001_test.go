/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024/1/9 14:59:43 星期二
 * @ProductName  : GoLand
 * @PrjectName   : go-examples
 * @File         : examples/stdlib/runtime/chan/features_usages_001_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package chan000

import (
	"fmt"
	"testing"
	"time"
)

func TestName_2024_01_09_14_59_43(t *testing.T) {
	t.Run("chan crud", func(t *testing.T) {
		var (
			ch1 chan int
			val int
			ok  bool
		)

		// 内置函数make()进行初始化，可创建无缓冲管道和有缓冲管道
		ch1 = make(chan int, 2)

		// 向管道ch1中写数据1
		ch1 <- 1

		// len(ch1)：缓冲区中元素个数 | cap(ch)：缓冲区容量
		fmt.Println(len(ch1), cap(ch1))

		// 向管道ch1中写数据1
		ch1 <- 2
		fmt.Println(len(ch1), cap(ch1))

		// 从管道中读取数据，存入ch1中
		val = <-ch1
		fmt.Println(val)
		fmt.Println(len(ch1), cap(ch1))

		// 使用逗号模式从管道中读取数据，ok为bool类型，表示是否读取到数据
		val, ok = <-ch1
		fmt.Println(val, ok)
		fmt.Println(len(ch1), cap(ch1))

		// 管道关闭&管道缓冲区中无数据 ⇌ ok为false
		// 管道关闭后读取管道需通过ok判断本次是否读取到数据，以防返回零值，误操作
		ch1 <- 1
		close(ch1)
		val, ok = <-ch1
		fmt.Println(val, ok)
		val, ok = <-ch1
		fmt.Println(val, ok)
		fmt.Println(len(ch1), cap(ch1))
	})
	// Result:
	// 1 2
	// 2 2
	// 1
	// 1 2
	// 2 true
	// 0 2
	// 1 true
	// 0 false
	// 0 2
	t.Run("write to a closed chan", func(t *testing.T) {
		var (
			ch1 chan int
		)
		ch1 = make(chan int, 1)
		close(ch1)
		ch1 <- 1
	})
	t.Run("close a closed chan", func(t *testing.T) {
		var (
			ch1 chan int
		)
		ch1 = make(chan int, 1)
		close(ch1)
		close(ch1)
	})
	t.Run("read from a nil chan", func(t *testing.T) {
		var (
			ch1 chan int
		)
		fmt.Println(ch1 == nil)
		go func() {
			time.Sleep(time.Second * 5)
		}()
		fmt.Println("main goroutine blocked..,")
		<-ch1
	})
	t.Run("read from a empty chan", func(t *testing.T) {
		var (
			ch1 chan int
		)
		ch1 = make(chan int, 2)
		go func() {
			time.Sleep(time.Second * 5)
		}()
		fmt.Println("main goroutine blocked..,")
		<-ch1
	})
	t.Run("write to a full chan", func(t *testing.T) {
		var (
			ch1 chan int
		)
		ch1 = make(chan int, 2)
		ch1 <- 1
		ch1 <- 2
		go func() {
			time.Sleep(time.Second * 5)
		}()
		fmt.Println("main goroutine blocked..,")
		ch1 <- 3
	})
	t.Run("write to a nil chan", func(t *testing.T) {
		var (
			ch1 chan int
		)
		fmt.Println(ch1 == nil)
		go func() {
			time.Sleep(time.Second * 5)
		}()
		fmt.Println("main goroutine blocked..,")
		ch1 <- 1
	})
}
