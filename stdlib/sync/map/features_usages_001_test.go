/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024/1/11 16:18:49 星期四
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : stdlib/sync/map/features_usages_001_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package syncmap000

import (
	"fmt"
	commontools "github.com/gainovel/testcase/tools/common"
	"sync"
	"testing"
)

func TestName_2024_01_11_16_18_49(t *testing.T) {
	//  go test -v -run TestName_2024_01_11_16_18_49/sync.Map_1
	t.Run("sync.Map 1", func(t *testing.T) {
		var (
			m1   sync.Map
			val  any
			ok   bool
			temp string
		)
		temp = "{}"
		fmt.Println()
		commontools.PrintAll(true, "var m1  sync.Map", "sync.Map m1 status", "", "sync.Map m1", temp)
		m1.Store("Jim", 80)
		m1.Store("Kevin", 90)
		m1.Store("Jane", 100)
		temp = ""
		m1.Range(func(key, value any) bool {
			temp += fmt.Sprintf("{%v:%v} ", key, value)
			return true
		})
		commontools.PrintAll(true, "add "+temp+"to sync.Map m1", "sync.Map m1 status", "", "sync.Map m1", temp)
		val, ok = m1.Load("Jim")
		commontools.PrintAll(true, "m1.Load(\"Jim\")", "Load Result", "", "val", val, "ok", ok)
		m1.Delete("Jim")
		temp = ""
		m1.Range(func(key, value any) bool {
			temp += fmt.Sprintf("{%v:%v} ", key, value)
			return true
		})
		commontools.PrintAll(true, "m1.Delete(\"Jim\")", "sync.Map m1 status", "", "sync.Map m1", temp)
	})
	// go test -v -run TestName_2024_01_11_16_18_49/sync.Map.LoadOrStore
	t.Run("sync.Map.LoadOrStore", func(t *testing.T) {
		var (
			m1     sync.Map
			actual any
			loaded bool
			temp   string
		)
		temp = "{}"
		fmt.Println()
		commontools.PrintAll(true, "var m1  sync.Map", "sync.Map m1 status", "", "sync.Map m1", temp)
		m1.Store("Jim", 80)
		m1.Store("Kevin", 90)
		m1.Store("Jane", 100)
		temp = ""
		m1.Range(func(key, value any) bool {
			temp += fmt.Sprintf("{%v:%v} ", key, value)
			return true
		})
		commontools.PrintAll(true, "add "+temp+"to sync.Map m1", "sync.Map m1 status", "", "sync.Map m1", temp)
		actual, loaded = m1.LoadOrStore("Jim", 81)
		commontools.PrintAll(true, "m1.LoadOrStore(\"Jim\", 81)", "LoadOrStore Result", "", "actual", actual, "loaded", loaded)
		m1.Delete("Jim")
		temp = ""
		m1.Range(func(key, value any) bool {
			temp += fmt.Sprintf("{%v:%v} ", key, value)
			return true
		})
		commontools.PrintAll(true, "m1.Delete(\"Jim\")", "sync.Map m1 status", "", "sync.Map m1", temp)
		actual, loaded = m1.LoadOrStore("Jim", 81)
		commontools.PrintAll(true, "m1.LoadOrStore(\"Jim\", 81)", "LoadOrStore Result", "", "actual", actual, "loaded", loaded)
	})
	// go test -v -run TestName_2024_01_11_16_18_49/sync.Map.LoadAndDelete
	t.Run("sync.Map.LoadAndDelete", func(t *testing.T) {
		var (
			m1     sync.Map
			value  any
			loaded bool
			temp   string
		)
		temp = "{}"
		fmt.Println()
		commontools.PrintAll(true, "var m1  sync.Map", "sync.Map m1 status", "", "sync.Map m1", temp)
		m1.Store("Jim", 80)
		m1.Store("Kevin", 90)
		m1.Store("Jane", 100)

		temp = ""
		m1.Range(func(key, value any) bool {
			temp += fmt.Sprintf("{%v:%v} ", key, value)
			return true
		})
		commontools.PrintAll(true, "add "+temp+"to sync.Map m1", "sync.Map m1 status", "", "sync.Map m1", temp)

		value, loaded = m1.LoadAndDelete("Jim")
		commontools.PrintAll(true, "m1.LoadAndDelete(\"Jim\")", "LoadAndDelete Result", "", "value", value, "loaded", loaded)

		temp = ""
		m1.Range(func(key, value any) bool {
			temp += fmt.Sprintf("{%v:%v} ", key, value)
			return true
		})
		commontools.PrintAll(true, "m1.Range", "sync.Map m1 status", "", "sync.Map m1", temp)

		value, loaded = m1.LoadAndDelete("Jim")
		commontools.PrintAll(true, "m1.LoadAndDelete(\"Jim\")", "LoadAndDelete Result", "", "value", value, "loaded", loaded)
	})
}
