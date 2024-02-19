/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024/1/10 11:22:46 星期三
 * @ProductName  : GoLand
 * @PrjectName   : go-examples
 * @File         : examples/stdlib/runtime/map/features_usages_001_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package map000

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/fatih/color"
)

func TestName_2024_01_10_11_22_46(t *testing.T) {
	t.Run("panic: assignment to entry in nil map", func(t *testing.T) {
		var (
			m1 map[int]int
		)
		m1[1] = 2
	})
	t.Run("concurrent map writes 1", func(t *testing.T) {
		var (
			m1                   map[int]int
			wg                   sync.WaitGroup
			n                    int
			targetKey, targetVal int
		)
		m1 = make(map[int]int)
		n = 1000
		for i := 0; i < n; i++ {
			key := rand.Intn(i+1000) + 100
			val := rand.Intn(key) + 88
			m1[key] = val
		}
		targetKey = 100
		targetVal = 10000
		wg.Add(n)
		for i := 0; i < n; i++ {
			go func(j int) {
				defer wg.Done()
				m1[targetKey] = targetVal
			}(i)
		}
		wg.Wait()
		fmt.Println(m1)
	})
	t.Run("concurrent map reads 1", func(t *testing.T) {
		var (
			m1 map[int]int
			wg sync.WaitGroup
			n  int
		)
		m1 = make(map[int]int)
		n = 1000
		for i := 0; i < n; i++ {
			key := rand.Intn(i+1000) + 100
			val := rand.Intn(key) + 88
			m1[key] = val
		}
		wg.Add(n)
		for i := 0; i < n; i++ {
			go func(j int) {
				defer wg.Done()
				color.New(color.FgHiMagenta).Printf("goroutin %d read start...\n", j)
				time.Sleep(time.Microsecond * 10)
				for k, v := range m1 {
					fmt.Printf("%d:%d ", k, v)
				}
				fmt.Println()
			}(i)
		}
		wg.Wait()
	})
	t.Run("concurrent map writes 2", func(t *testing.T) {
		// 使用windows terminal 在./cmd/main.go测试ConcurrentMapReads()
		ConcurrentMapReads()
	})
	t.Run("concurrent map reads 2", func(t *testing.T) {
		// 使用windows terminal 在./cmd/main.go测试ConcurrentMapWrites()
		ConcurrentMapWrites()
	})
}
