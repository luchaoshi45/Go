package algorithm

import (
	"fmt"
)

// ScheduleTasks 函数接收任务列表和冷却时间 n，返回完成所有任务所需的总时间片次数
func ScheduleTasks(tasks []byte, n int) int {
	// 使用 map 记录每个任务的剩余数量
	taskCount := make(map[byte]int)
	for _, task := range tasks {
		taskCount[task]++
	}

	res := 0
	dn := 0
	for len(taskCount) > 0 {
		ok := len(taskCount) - n - 1
		if ok >= 0 {
			res += len(taskCount)
		} else {
			res += n + 1
		}
		dn = 0
		for k, v := range taskCount {
			if v == 1 {
				dn++
				delete(taskCount, k)
			} else {
				taskCount[k]--
			}
		}
	}

	return res - n - 1 + dn
}

func TestScheduleTasks() {
	tasks := []byte{'A', 'A', 'A', 'A', 'A', 'B', 'C', 'C', 'C', 'C', 'C'}
	n := 2
	totalTime := ScheduleTasks(tasks, n)
	fmt.Println("完成所有任务所需的总时间片次数:", totalTime)
}
