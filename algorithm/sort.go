package algorithm

import (
	"fmt"
)

func TestSort() {
	nums := []int{5, 3, 15, 0}
	res := SelectionSort(nums)
	fmt.Println(res)

	nums1 := []int{5, 3, 15, 0}
	res1 := BubbleSort(nums1)
	fmt.Println(res1)

	nums2 := []int{5, 3, 15, 0}
	res2 := InsertSort(nums2)
	fmt.Println(res2)

	numsx := RandNValidator(10, 50)
	fmt.Println(SelectionSort(numsx))
	fmt.Println(BubbleSort(numsx))
	fmt.Println(InsertSort(numsx))
}

func SelectionSort(nums []int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ { // [0, n-2]
		minIndex := i                // 假定最小值下标
		for j := i + 1; j < n; j++ { // [i+1, n-1]  第一次是 [1, n-1]
			if nums[j] < nums[minIndex] { // [i+1, n-1] 中寻找最小值对应的下标
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
	return nums
}

func BubbleSort(nums []int) []int {
	n := len(nums)
	for end := n - 1; end > 0; end-- { // [0, n-1]
		for j := 0; j < end; j++ { // [0, n-2]
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func InsertSort(nums []int) []int {
	n := len(nums)

	for i := 0; i < n; i++ { // [0, n-1]
		for j := i - 1; j >= 0 && nums[j] > nums[j+1]; j-- { // [..., i-1]
			nums[j], nums[j+1] = nums[j+1], nums[j]
		}
	}

	return nums
}
