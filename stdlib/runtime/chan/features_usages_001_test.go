/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024/1/9 14:59:43 星期二
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : stdlib/runtime/chan/features_usages_001_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/
// qmemcodestart
package chan000

import (
	"fmt"
	commonprint "github.com/gainovel/testcase/tools/common/print"
	"testing"
	"time"
)

var (
	myfmt = commonprint.MyFmt
)

func sleepG() {
	time.Sleep(time.Hour)
}
func TestName_2024_01_09_14_59_43(t *testing.T) {
	// chan的基本用法
	// make 001/chan_crud -f Makefiles/stdlib/runtime/chan.mk
	t.Run("chan crud", func(t *testing.T) {
		var (
			ch1 chan int
			val int
			ok  bool
		)
		myfmt.VarInitPrintln(`var (
	ch1 chan int
	val int
	ok  bool
)`)
		// 内置函数make()进行初始化，可创建无缓冲管道和有缓冲管道
		ch1 = make(chan int, 2)

		// 向管道ch1中写数据1
		ch1 <- 1

		// len(ch1)：缓冲区中元素个数 | cap(ch)：缓冲区容量
		myfmt.ColorDescPrintln("ch1 = make(chan int, 2);ch1 <- 1")
		myfmt.KeyValuePrintln("len(ch1)", len(ch1), "cap(ch1)", cap(ch1))

		// 向管道ch1中写数据1
		ch1 <- 2
		myfmt.ColorDescPrintln("ch1 <- 2")
		myfmt.KeyValuePrintln("len(ch1)", len(ch1), "cap(ch1)", cap(ch1))

		// 从管道中读取数据，存入ch1中
		val = <-ch1
		myfmt.ColorDescPrintln("val = <-ch1")
		myfmt.KeyValuePrintln("len(ch1)", len(ch1), "cap(ch1)", cap(ch1), "val", val)

		// 使用逗号模式从管道中读取数据，ok为bool类型，表示是否读取到数据
		val, ok = <-ch1
		myfmt.ColorDescPrintln("val, ok = <-ch1")
		myfmt.KeyValuePrintln("len(ch1)", len(ch1), "cap(ch1)", cap(ch1), "val", val, "ok", ok)

		// 管道关闭&管道缓冲区中无数据 ⇌ ok为false
		// 管道关闭后读取管道需通过ok判断本次是否读取到数据，以防返回零值，误操作
		close(ch1)
		val, ok = <-ch1
		myfmt.ColorDescPrintln("val, ok = <-ch1")
		myfmt.KeyValuePrintln("len(ch1)", len(ch1), "cap(ch1)", cap(ch1), "val", val, "ok", ok)
	})

	// 协程读取管道时，阻塞的条件有3个：①管道无缓冲区②管道缓冲区中无数据③管道为nil。
	// make 001/1.read_no_buf -f Makefiles/stdlib/runtime/chan.mk
	t.Run("1.read no buf", func(t *testing.T) {
		go sleepG()
		var (
			ch1 chan int
		)
		ch1 = make(chan int)
		fmt.Println("管道ch1无缓冲区，读取阻塞...")
		<-ch1
		fmt.Println("before return...")
	})
	// make 001/2.read_no_data -f Makefiles/stdlib/runtime/chan.mk
	t.Run("2.read no data", func(t *testing.T) {
		go sleepG()
		var (
			ch1 chan int
		)
		ch1 = make(chan int, 2)
		fmt.Println("管道ch1缓冲区中无数据，读取阻塞...")
		<-ch1
		fmt.Println("before return...")
	})
	// make 001/3.read_nil_chan -f Makefiles/stdlib/runtime/chan.mk
	t.Run("3.read nil chan", func(t *testing.T) {
		go sleepG()
		var (
			ch1 chan int
		)
		fmt.Println("管道ch1为nil，读取阻塞...")
		<-ch1
		fmt.Println("before return...")
	})

	// 协程写入管道时，阻塞的条件有3个：①管道无缓冲区②管道缓冲区已满③管道为nil。
	// make 001/1.write_no_buf -f Makefiles/stdlib/runtime/chan.mk
	t.Run("1.write no buf", func(t *testing.T) {
		go sleepG()
		var (
			ch1 chan int
		)
		ch1 = make(chan int)
		fmt.Println("管道ch1无缓冲区，写入阻塞...")
		ch1 <- 1
		fmt.Println("before return...")
	})
	// make 001/2.write_full_data -f Makefiles/stdlib/runtime/chan.mk
	t.Run("2.write full data", func(t *testing.T) {
		go sleepG()
		var (
			ch1 chan int
		)
		ch1 = make(chan int, 2)
		ch1 <- 1
		ch1 <- 1
		fmt.Println("管道ch1缓冲区已满，写入阻塞...")
		ch1 <- 1
		fmt.Println("before return...")
	})
	// 001/3.write_nil_chan -f Makefiles/stdlib/runtime/chan.mk
	t.Run("3.write nil chan", func(t *testing.T) {
		go sleepG()
		var (
			ch1 chan int
		)
		fmt.Println("管道ch1为nil，写入阻塞...")
		ch1 <- 1
		fmt.Println("before return...")
	})

	// panic的情况有两种，①向关闭的管道写数据会触发panic②关闭已经关闭的管道

	// make 001/1.write_to_a_closed_chan -f Makefiles/stdlib/runtime/chan.mk
	// panic: send on closed channel
	t.Run("1.write to a closed chan", func(t *testing.T) {
		var (
			ch1 chan int
		)
		ch1 = make(chan int, 1)
		close(ch1)
		ch1 <- 1
	})
	// make 001/2.close_a_closed_chan -f Makefiles/stdlib/runtime/chan.mk
	// panic: close of closed channel
	t.Run("2.close a closed chan", func(t *testing.T) {
		var (
			ch1 chan int
		)
		ch1 = make(chan int, 1)
		close(ch1)
		close(ch1)
	})

	// 管道关闭后，仍然可以读取数据
	// 使用逗号模式(val,ok := <-ch1)从管道中读取数据，ok为bool类型，表示是否读取到数据
	// 管道关闭&管道缓冲区中无数据 ⇌ ok为false；管道关闭后读取管道需通过ok判断本次是否读取到数据，以防返回零值，误操作

	// make 001/read_from_a_close_chan -f Makefiles/stdlib/runtime/chan.mk
	t.Run("read from a close chan", func(t *testing.T) {
		var (
			ch1 chan int
			val int
			ok  bool
		)
		ch1 = make(chan int, 2)
		ch1 <- 1
		ch1 <- 2
		close(ch1)
		myfmt.VarInitPrintln(`var (
	ch1 chan int
	val int
	ok  bool
)`)
		myfmt.ColorDescPrintln("ch1 = make(chan int, 2);ch1 <- 1;ch1 <- 2;close(ch1)")
		myfmt.KeyValuePrintln("len(ch1)", len(ch1), "cap(ch1)", cap(ch1))
		val, ok = <-ch1
		myfmt.ColorDescPrintln("val, ok = <-ch1")
		myfmt.KeyValuePrintln("len(ch1)", len(ch1), "cap(ch1)", cap(ch1), "val", val, "ok", ok)
		val, ok = <-ch1
		myfmt.ColorDescPrintln("val, ok = <-ch1")
		myfmt.KeyValuePrintln("len(ch1)", len(ch1), "cap(ch1)", cap(ch1), "val", val, "ok", ok)
		val, ok = <-ch1
		myfmt.ColorDescPrintln("val, ok = <-ch1")
		myfmt.KeyValuePrintln("len(ch1)", len(ch1), "cap(ch1)", cap(ch1), "val", val, "ok", ok)
	})
}

// qmemcodeend

// qmemoutputstart
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: make 001/chan_crud -f Makefiles/stdlib/runtime/chan.mk
//*  Test Result:
//*  👇 变量初始化：
//*  --------------------------------------------------------------------------------
//*  var (
//*          ch1 chan int
//*          val int
//*          ok  bool
//*  )
//*  --------------------------------------------------------------------------------
//*  👇
//*  ch1 = make(chan int, 2);ch1 <- 1
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |          len(ch1) | 1
//*  |          cap(ch1) | 2
//*  --------------------------------------------------------------------------------
//*  👇
//*  ch1 <- 2
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |          len(ch1) | 2
//*  |          cap(ch1) | 2
//*  --------------------------------------------------------------------------------
//*  👇
//*  val = <-ch1
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |          len(ch1) | 1
//*  |          cap(ch1) | 2
//*  |               val | 1
//*  --------------------------------------------------------------------------------
//*  👇
//*  val, ok = <-ch1
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |          len(ch1) | 0
//*  |          cap(ch1) | 2
//*  |               val | 2
//*  |                ok | true
//*  --------------------------------------------------------------------------------
//*  👇
//*  val, ok = <-ch1
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |          len(ch1) | 0
//*  |          cap(ch1) | 2
//*  |               val | 0
//*  |                ok | false
//*  --------------------------------------------------------------------------------
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: make 001/1.read_no_buf -f Makefiles/stdlib/runtime/chan.mk
//*  Test Result: 管道ch1无缓冲区，读取阻塞...
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: make 001/2.read_no_data -f Makefiles/stdlib/runtime/chan.mk
//*  Test Result: 管道ch1缓冲区中无数据，读取阻塞...
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: make 001/3.read_nil_chan -f Makefiles/stdlib/runtime/chan.mk
//*  Test Result: 管道ch1为nil，读取阻塞...
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: make 001/1.write_no_buf -f Makefiles/stdlib/runtime/chan.mk
//*  Test Result: 管道ch1无缓冲区，写入阻塞...
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: make 001/2.write_full_data -f Makefiles/stdlib/runtime/chan.mk
//*  Test Result: 管道ch1缓冲区已满，写入阻塞...
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: make 001/3.write_nil_chan -f Makefiles/stdlib/runtime/chan.mk
//*  Test Result: 管道ch1为nil，写入阻塞...
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: make 001/1.write_to_a_closed_chan -f Makefiles/stdlib/runtime/chan.mk
//*  Test Result: panic: send on closed channel
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: make 001/2.close_a_closed_chan -f Makefiles/stdlib/runtime/chan.mk
//*  Test Result: panic: close of closed channel
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: make 001/read_from_a_close_chan -f Makefiles/stdlib/runtime/chan.mk
//*  Test Result:
//*  👇 变量初始化：
//*  --------------------------------------------------------------------------------
//*  var (
//*          ch1 chan int
//*          val int
//*          ok  bool
//*  )
//*  --------------------------------------------------------------------------------
//*  👇
//*  ch1 = make(chan int, 2);ch1 <- 1;ch1 <- 2;close(ch1)
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |          len(ch1) | 2
//*  |          cap(ch1) | 2
//*  --------------------------------------------------------------------------------
//*  👇
//*  val, ok = <-ch1
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |          len(ch1) | 1
//*  |          cap(ch1) | 2
//*  |               val | 1
//*  |                ok | true
//*  --------------------------------------------------------------------------------
//*  👇
//*  val, ok = <-ch1
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |          len(ch1) | 0
//*  |          cap(ch1) | 2
//*  |               val | 2
//*  |                ok | true
//*  --------------------------------------------------------------------------------
//*  👇
//*  val, ok = <-ch1
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |          len(ch1) | 0
//*  |          cap(ch1) | 2
//*  |               val | 0
//*  |                ok | false
//*  --------------------------------------------------------------------------------
//**************************************************************************************
// qmemoutputend
