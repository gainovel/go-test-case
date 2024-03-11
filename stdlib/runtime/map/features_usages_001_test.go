/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024/1/10 11:22:46 星期三
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : stdlib/runtime/map/features_usages_001_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package map000

import (
	"testing"

	commonprint "github.com/gainovel/testcase/tools/common/print"
)

var (
	myfmt = commonprint.MyFmt
)

func TestName_2024_01_10_11_22_46(t *testing.T) {
	// 多协程同时写会报错concurrent map writes
	// go test -v -run TestName_2024_01_10_11_22_46/concurrent_map_write
	t.Run("concurrent map write", func(t *testing.T) {
		// 使用windows terminal 在./cmd/main.go测试ConcurrentMapWrites()
		ConcurrentMapWrites()
	})
	// panic的情况👉给nil map添加key-value
	// go test -v -run TestName_2024_01_10_11_22_46/assignment_to_entry_in_nil_map
	t.Run("assignment to entry in nil map", func(t *testing.T) {
		var (
			m1 map[int]int
		)
		m1[1] = 2
	})

	// map 的简单使用
	// 使用内置函数delete()进行删除
	// 查询map时，使用逗号模式(val,ok)获取值，避免操作零值，ok表示key是否存在
	// go test -v -run TestName_2024_01_10_11_22_46/map_crud
	t.Run("map crud", func(t *testing.T) {
		var (
			m1  map[int]int
			key int
			ok  bool
		)
		m1 = make(map[int]int)
		m1[1] = 101
		m1[2] = 102
		myfmt.VarInitPrintln(`var (
	m1 map[int]int
)`)
		myfmt.ColorDescPrintln("m1 = make(map[int]int)", "m1[1] = 101", "m1[2] = 102")
		myfmt.KeyValuePrintln("m1", m1, "len(m1)", len(m1))
		delete(m1, 1)
		myfmt.ColorDescPrintln("delete(m1, 1)")
		myfmt.KeyValuePrintln("m1", m1, "len(m1)", len(m1))
		key, ok = m1[1]
		myfmt.ColorDescPrintln("key,ok = m1[1]")
		myfmt.KeyValuePrintln("m1", m1, "len(m1)", len(m1), "key", key, "ok", ok)
		key, ok = m1[2]
		myfmt.ColorDescPrintln("key,ok = m1[2]")
		myfmt.KeyValuePrintln("m1", m1, "len(m1)", len(m1), "key", key, "ok", ok)
	})
}
