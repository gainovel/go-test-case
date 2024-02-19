/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024/1/9 18:10:28 星期二
 * @ProductName  : GoLand
 * @PrjectName   : go-examples
 * @File         : examples/stdlib/runtime/slice/features_usages_001_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package slice

import (
	"fmt"
	commonconsts "gainovel.com/go/testcase/consts/common"
	commontools "gainovel.com/go/testcase/tools/common"
	"github.com/fatih/color"
	"sync"
	"testing"
	"time"
)

func TestName_2024_01_09_18_10_28(t *testing.T) {
	t.Run("slice can't use val,ok ", func(t *testing.T) {
		//var (
		//	s1 []int
		//)
		//s1 = make([]int, 0, 5)
		//s1 = append(s1, 1)
		//val, ok := s1[0]
	})
	t.Run("index out of range", func(t *testing.T) {
		var (
			s1 []int
		)
		s1 = make([]int, 0, 5)
		s1 = append(s1, 1)
		fmt.Println(s1[0])
		fmt.Println(s1[1])
	})
	// Result👉panic: runtime error: index out of range [1] with length 1
	t.Run("nil slice", func(t *testing.T) {
		var (
			s1 []int
		)
		fmt.Println(s1[0])
	})
	// Result👉panic: runtime error: index out of range [0] with length 0
	t.Run("simple expression slice operation", func(t *testing.T) {
		var (
			n    int
			arr1 [6]int
			s1   []int
		)
		n = 6
		s1 = make([]int, n)
		for i := 0; i < n; i++ {
			arr1[i] = i + 1
			s1[i] = i + 10
		}

		commontools.Print88HorizontalLine()
		commontools.Println(false, "原数组(arr1)")
		commontools.Println(true, "arr1", arr1, "len(arr1)", len(arr1), "cap(arr1)", cap(arr1))
		commontools.Println(false, "原slice(s1)")
		commontools.Println(true, "s1", s1, "len(s1)", len(s1), "cap(s1)", cap(s1))
		commontools.Print88HorizontalLine()
		commontools.PrintDownHand("n=6")
		// 简单表达式切片
		// 通过对数组或其它slice进行切片(`arr[1:len(arr)-1] | s1[1:len(arr)-1]`)可以获得一个新的slice
		// 新的slice会和数组或原slice共享底层数组
		// 简单表达式切片并未限制新的slice的容量，也就是说通过append()添加新元素时有覆盖原数组或者原slice元素的风险
		commontools.Print88HorizontalLine()
		s2 := arr1[1 : n-2]
		commontools.Println(false, "arr1切片(arr1[1:n-2])得到的slice(s2)")
		commontools.Println(true, "s2", s2, "len(s2)", len(s2), "cap(s2)", cap(s2))
		s3 := s1[1 : n-2]
		commontools.Println(false, "s1切片(s1[1:n-2])得到的slice(s3)")
		commontools.Println(true, "s3", s3, "len(s3)", len(s3), "cap(s3)", cap(s3))
		commontools.Print88HorizontalLine()
		commontools.PrintDownHand("")

		commontools.Print88HorizontalLine()
		commontools.Println(false, "s2和s3通过append()添加新元素100")
		s2 = append(s2, 100)
		s3 = append(s3, 100)
		commontools.Println(false, "原数组(arr1)：")
		commontools.Println(true, "arr1", arr1, "len(arr1)", len(arr1), "cap(arr1)", cap(arr1))
		commontools.Println(false, "原slice(s1)：")
		commontools.Println(true, "s1", s1, "len(s1)", len(s1), "cap(s1)", cap(s1))
		fmt.Println(commonconsts.VERTICAL1SPACE1)
		commontools.Println(false, "s2 = append(s2, 100)")
		commontools.Println(true, "s2", s2, "len(s2)", len(s2), "cap(s2)", cap(s2))
		commontools.Println(false, "s3 = append(s3, 100)")
		commontools.Println(true, "s3", s3, "len(s3)", len(s3), "cap(s3)", cap(s3))
		commontools.Print88HorizontalLine()
		commontools.PrintDownHand("")

		commontools.Print88HorizontalLine()
		commontools.Println(false, "s2和s3通过append()添加新元素101, 102, 103")
		s2 = append(s2, 101, 102, 103)
		s3 = append(s3, 101, 102, 103)
		commontools.Println(false, "原数组(arr1)：")
		commontools.Println(true, "arr1", arr1, "len(arr1)", len(arr1), "cap(arr1)", cap(arr1))
		commontools.Println(false, "原slice(s1)：")
		commontools.Println(true, "s1", s1, "len(s1)", len(s1), "cap(s1)", cap(s1))
		fmt.Println(commonconsts.VERTICAL1SPACE1)
		commontools.Println(false, "s2 = append(s2, 101, 102, 103)")
		commontools.Println(true, "s2", s2, "len(s2)", len(s2), "cap(s2)", cap(s2))
		commontools.Println(false, "s3 = append(s3, 101, 102, 103)")
		commontools.Println(true, "s3", s3, "len(s3)", len(s3), "cap(s3)", cap(s3))
		commontools.Print88HorizontalLine()
	})
	t.Run("expand expression slice operation", func(t *testing.T) {
		var (
			n    int
			arr1 [6]int
			s1   []int
		)
		n = 6
		s1 = make([]int, n)
		for i := 0; i < n; i++ {
			arr1[i] = i + 1
			s1[i] = i + 10
		}

		commontools.Print88HorizontalLine()
		commontools.Println(false, "原数组(arr1)")
		commontools.Println(true, "arr1", arr1, "len(arr1)", len(arr1), "cap(arr1)", cap(arr1))
		commontools.Println(false, "原slice(s1)")
		commontools.Println(true, "s1", s1, "len(s1)", len(s1), "cap(s1)", cap(s1))
		commontools.Print88HorizontalLine()
		commontools.PrintDownHand("n=6")
		// 简单表达式切片
		// 通过对数组或其它slice进行切片(`arr[1:len(arr)-1] | s1[1:len(arr)-1]`)可以获得一个新的slice
		// 新的slice会和数组或原slice共享底层数组
		// 简单表达式切片并未限制新的slice的容量，也就是说通过append()添加新元素时有覆盖原数组或者原slice元素的风险
		commontools.Print88HorizontalLine()
		s2 := arr1[1 : n-2 : n-2]
		commontools.Println(false, "arr1切片(arr1[1:n-2])得到的slice(s2)")
		commontools.Println(true, "s2", s2, "len(s2)", len(s2), "cap(s2)", cap(s2))
		s3 := s1[1 : n-2 : n-2]
		commontools.Println(false, "s1切片(s1[1:n-2])得到的slice(s3)")
		commontools.Println(true, "s3", s3, "len(s3)", len(s3), "cap(s3)", cap(s3))
		commontools.Print88HorizontalLine()
		commontools.PrintDownHand("")

		commontools.Print88HorizontalLine()
		commontools.Println(false, "s2和s3通过append()添加新元素100")
		s2 = append(s2, 100)
		s3 = append(s3, 100)
		commontools.Println(false, "原数组(arr1)：")
		commontools.Println(true, "arr1", arr1, "len(arr1)", len(arr1), "cap(arr1)", cap(arr1))
		commontools.Println(false, "原slice(s1)：")
		commontools.Println(true, "s1", s1, "len(s1)", len(s1), "cap(s1)", cap(s1))
		fmt.Println(commonconsts.VERTICAL1SPACE1)
		commontools.Println(false, "s2 = append(s2, 100)")
		commontools.Println(true, "s2", s2, "len(s2)", len(s2), "cap(s2)", cap(s2))
		commontools.Println(false, "s3 = append(s3, 100)")
		commontools.Println(true, "s3", s3, "len(s3)", len(s3), "cap(s3)", cap(s3))
		commontools.Print88HorizontalLine()
		commontools.PrintDownHand("")
	})
	t.Run("", func(t *testing.T) {
		var (
			n    int
			arr1 [18]int
			wg   sync.WaitGroup
		)
		n = 10000
		wg.Add(n)
		for i := 0; i < n; i++ {
			go func(j int) {
				defer wg.Done()
				color.New(color.FgHiMagenta).Printf("goroutine %d start...\n", j)
				s1 := arr1[1:12]
				time.Sleep(time.Microsecond * 10)
				s1[8]++
			}(i)
		}
		wg.Wait()
		fmt.Println(arr1)
	})
}
