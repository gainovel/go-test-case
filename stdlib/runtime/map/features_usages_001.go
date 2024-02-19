/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024/1/10 11:56:06 星期三
 * @ProductName  : GoLand
 * @PrjectName   : go-examples
 * @File         : examples/stdlib/runtime/map/features_usages_001.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package map000

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"sync"
	"time"
)

func ConcurrentMapReads() {
	var (
		m1                   map[int]int
		wg                   sync.WaitGroup
		n                    int
		targetKey, targetVal int
	)
	n = 1000
	m1 = make(map[int]int, n)

	// 先随机添加一些元素值
	for i := 0; i < n; i++ {
		key := rand.Intn(i+1000) + 100
		val := rand.Intn(key) + 88
		m1[key] = val
	}

	targetKey = 100
	targetVal = 10000
	m1[targetKey] = targetVal
	// 开启1000个协程并发读
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(j int) {
			defer wg.Done()
			color.New(color.FgHiMagenta).Printf("goroutine %d read start...\n", j)
			time.Sleep(time.Microsecond * 10)
			val := m1[targetKey]
			fmt.Printf("groutine %s get targetKey(%d), targetVal(%d)\n", color.New(color.FgHiCyan).Sprintf("%d", j), targetKey, val)
		}(i)
	}
	wg.Wait()
}
func ConcurrentMapWrites() {
	var (
		m1                   map[int]int
		wg                   sync.WaitGroup
		n                    int
		targetKey, targetVal int
	)
	n = 1000
	m1 = make(map[int]int, n)

	// 先随机添加一些元素值
	for i := 0; i < n; i++ {
		key := rand.Intn(i+1000) + 100
		val := rand.Intn(key) + 88
		m1[key] = val
	}

	targetKey = 100
	targetVal = 10000

	// 开启1000个协程并发写
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(j int) {
			defer wg.Done()
			color.New(color.FgHiMagenta).Printf("goroutine %d write start...\n", j)
			//fmt.Printf("groutine %s write targetKey(%d), targetVal(%d)\n", color.New(color.FgHiCyan).Sprintf("%d", j), targetKey, targetVal)
			time.Sleep(time.Microsecond * 10)
			m1[targetKey] = targetVal
		}(i)
	}
	wg.Wait()
}
