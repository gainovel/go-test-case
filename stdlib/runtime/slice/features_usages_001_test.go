/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024/1/9 18:10:28 星期二
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : stdlib/runtime/slice/features_usages_001_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package slice

import (
	"testing"

	commonprint "gainovel.com/go/testcase/tools/common/print"
)

var (
	myfmt = commonprint.MyFmt
)

func TestName_2024_01_09_18_10_28(t *testing.T) {

	// panic的情况：①索引超出范围。
	// go test -run TestName_2024_01_09_18_10_28/index_out_of_range
	// Result👉panic: runtime error: index out of range [1] with length 1
	t.Run("index out of range", func(t *testing.T) {
		var (
			s1 []int
		)
		s1 = make([]int, 0, 5)
		s1 = append(s1, 1)
		myfmt.VarInitPrintln(`var (
	s1 []int
)`)
		myfmt.ColorDescPrintln("s1 = make([]int, 0, 5);s1 = append(s1, 1)")
		myfmt.KeyValuePrintln("s1", s1, "s1[0]", s1[0])
		myfmt.KeyValuePrintln("s1[1]", s1[1])
	})

	// 多个slice可能共享同一底层数组。
	// 简单表达式(s1[low:high])，简单表达式切片并未限制新的slice的容量，也就是说通过内置函数append()添加新元素时有覆盖原数组或者原slice元素的风险
	// go test -run TestName_2024_01_09_18_10_28/share_array
	t.Run("share array", func(t *testing.T) {
		var (
			s1   []int
			s1_1 []int
			s2   []int
			s2_1 []int
		)
		s1 = []int{1, 2, 3, 4, 5}
		s2 = []int{1, 2, 3, 4, 5}
		s1_1 = s1[1:3]
		s2_1 = s2[1:3:3]
		myfmt.VarInitPrintln(`var (
	s1   []int
	s1_1 []int
	s2   []int
	s2_1 []int
)`)
		myfmt.ColorDescPrintln("s1 = []int{1, 2, 3, 4, 5};s1_1 = s1[1:3]", "s2 = []int{1, 2, 3, 4, 5};s2_1 = s2[1:3:3]")
		myfmt.KeyValuePrintln("s1", s1, "len(s1)", len(s1), "cap(s1)", cap(s1),
			"s1_1", s1_1, "len(s1_1)", len(s1_1), "cap(s1_1)", cap(s1_1))
		myfmt.KeyValuePrintln(
			"s2", s2, "len(s2)", len(s2), "cap(s2)", cap(s2),
			"s2_1", s2_1, "len(s2_1)", len(s2_1), "cap(s2_1)", cap(s2_1))
		s1_1 = append(s1_1, 100)
		s2_1 = append(s2_1, 100)
		myfmt.ColorDescPrintln("s1_1 = append(s1_1, 100)", "s2_1 = append(s2_1, 100)")
		myfmt.KeyValuePrintln("s1", s1, "len(s1)", len(s1), "cap(s1)", cap(s1),
			"s1_1", s1_1, "len(s1_1)", len(s1_1), "cap(s1_1)", cap(s1_1))
		myfmt.KeyValuePrintln(
			"s2", s2, "len(s2)", len(s2), "cap(s2)", cap(s2),
			"s2_1", s2_1, "len(s2_1)", len(s2_1), "cap(s2_1)", cap(s2_1))
	})

	// go test -run TestName_2024_01_09_18_10_28/slice_copy
	t.Run("slice copy", func(t *testing.T) {
		var (
			s1 []int
			s2 []int
		)
		s1 = []int{1, 2, 3, 4, 5}
		s2 = []int{100, 100, 100}
		myfmt.VarInitPrintln(`var (
	s1 []int
	s2 []int
)`)
		myfmt.ColorDescPrintln("s1 = []int{1, 2, 3, 4, 5}", "s2 = []int{100, 100, 100}")
		myfmt.KeyValuePrintln("s1", s1, "s2", s2)
		copy(s2, s1)
		myfmt.ColorDescPrintln("copy(s2, s1)")
		myfmt.KeyValuePrintln("s1", s1, "s2", s2)
		s1 = []int{1, 2, 3, 4, 5}
		s2 = []int{100, 100, 100}
		myfmt.ColorDescPrintln("s1 = []int{1, 2, 3, 4, 5}", "s2 = []int{100, 100, 100}")
		myfmt.KeyValuePrintln("s1", s1, "s2", s2)
		copy(s1, s2)
		myfmt.ColorDescPrintln("copy(s1, s2)")
		myfmt.KeyValuePrintln("s1", s1, "s2", s2)
	})

	// slice扩容规则。①slice扩容时有几个关键的值需要提前说明一下👉oldCap：扩容前容量、oldLen：扩容前元素个数、cap：扩容所需最小容量、newCap：预估容量②Go1.15扩容规则👉如果oldCap(扩容前的容量)翻倍之后还是小于cap(扩容所需最小容量)，那么newCap(预估容量)就等于cap(扩容所需最小容量)，如果不满足第一条，而且oldLen(扩容前元素个数)小于1024，那么newCap(预估容量)=oldCap(扩容前的容量)*2，如果不满足第一条，而且oldLen(扩容前元素个数)大于等于1024，那就循环扩容四分之一，直到大于等于所需最小容量③Go1.16扩容规则👉Go1.16中有了些变化，和1024比较的不再是oldLen(扩容前元素个数)，而是oldCap(扩容前的容量)；Go1.18扩容规则👉到了Go1.18时，又改成不和1024比较了，而是和256比较；并且扩容的增量也有所变化，除了每次扩容1/4，还得加上256*3/4
	// 预估容量不一定为最终申请的容量；Go的内存管理，申请内存时都有一定的规格(8,16,32,48…)，e.g. int类型slice，预估容量为5，64位操作系统，需要申请40字节，但内存规格中不存在40，Go的内存管理会帮我们匹配到足够大、且最接近的规格48，最终申请的容量是6。
	// go test -run TestName_2024_01_09_18_10_28/expansion_experiment
	t.Run("expansion experiment", func(t *testing.T) {
		var (
			s1 []int
			s2 []int
		)
		s1 = make([]int, 0, 1)
		s2 = make([]int, 0, 1)
		myfmt.VarInitPrintln(`var (
	s1 []int
	s2 []int
)`)
		myfmt.ColorDescPrintln("s1 = make([]int, 0, 1)", "s2 = make([]int, 0, 1)")
		myfmt.KeyValuePrintln("s1", s1, "len(s1)", len(s1), "cap(s1)", cap(s1))
		myfmt.KeyValuePrintln("s2", s2, "len(s2)", len(s2), "cap(s2)", cap(s2))
		s1 = append(s1, 1, 2, 3, 4)
		s2 = append(s2, 1, 2, 3, 4, 5)
		myfmt.ColorDescPrintln("s1 = append(s1, 1, 2, 3, 4)", "s2 = append(s2, 1, 2, 3, 4, 5)")
		myfmt.KeyValuePrintln("s1", s1, "len(s1)", len(s1), "cap(s1)", cap(s1))
		myfmt.KeyValuePrintln("s2", s2, "len(s2)", len(s2), "cap(s2)", cap(s2))
	})
}
