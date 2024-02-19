/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024/1/9 15:08:44 星期二
 * @ProductName  : GoLand
 * @PrjectName   : go-examples
 * @File         : examples/leetcode/interview_classic_150_questions_test.go
 * @Version      : v0.1.0
 * @Description  : website https://leetcode.cn/studyplan/top-interview-150/
 **/

package leetcode

import (
	"testing"
)

// merge
/**
 * @author: Gainovel
 * @description: 88.合并两个有序数组 https://leetcode.cn/problems/merge-sorted-array/description/
 * @date: 2024/1/9 15:09:08
 * @params:
 * @return:
 */

// 解法一 | 时间复杂度：O(m+n) 空间复杂度：O(m)
/*func merge(nums1 []int, m int, nums2 []int, n int) {
	if m+n == 0 {
		return
	}
	var (
		i, j, k int
		temp    = make([]int, m, m)
	)
	copy(temp, nums1[:m])
	for i < m && j < n {
		if temp[i] < nums2[j] {
			nums1[k] = temp[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
		k++
	}
	for i < m {
		nums1[k] = temp[i]
		i++
		k++
	}
	for j < n {
		nums1[k] = nums2[j]
		j++
		k++
	}
	fmt.Println(nums1)
}*/
// 解法二 | 时间复杂度：O(m+n) 空间复杂度：O(1)
func merge(nums1 []int, m int, nums2 []int, n int) {
	if m+n == 0 {
		return
	}
	var (
		i, j int //i,j分别指向nums1和nums2的最后一个元素
		tail int // tail执行nums1的末尾
	)
	i, j = m-1, n-1
	tail = len(nums1) - 1
	for i >= 0 && j >= 0 {
		if nums1[i] < nums2[j] {
			nums1[tail] = nums2[j]
			j--
		} else {
			nums1[tail] = nums1[i]
			i--
		}
		tail--
	}
	for i >= 0 {
		nums1[tail] = nums1[i]
		i--
		tail--
	}
	for j >= 0 {
		nums1[tail] = nums2[j]
		j--
		tail--
	}
}

func TestName_2024_01_09_15_08_44(t *testing.T) {
	t.Run("88.合并两个有序数组", func(t *testing.T) {
		tests := []struct {
			name  string
			m     int
			n     int
			nums1 []int
			nums2 []int
		}{
			{
				name:  "case1",
				m:     3,
				n:     3,
				nums1: []int{1, 2, 3, 0, 0, 0},
				nums2: []int{2, 5, 6},
			}, {
				name:  "case2",
				m:     1,
				n:     0,
				nums1: []int{1},
				nums2: []int{},
			}, {
				name:  "case3",
				m:     0,
				n:     1,
				nums1: []int{0},
				nums2: []int{1},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				merge(tt.nums1, tt.m, tt.nums2, tt.n)
				t.Log("Result: ", tt.nums1)
			})
		}
	})
}
