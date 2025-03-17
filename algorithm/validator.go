package algorithm

import (
	"math/rand"
	"time"
)

// 验证器
func Validator(n, v int) []int { // 返回一个 n 个元素， 元素值在 1-v 范围的切片
	if n <= 0 {
		return []int{}
	}
	if v < 1 {
		return []int{}
	}

	// 创建一个局部随机数生成器
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	result := make([]int, n)
	for i := 0; i < n; i++ {
		// 使用局部随机数生成器生成 1 到 v 之间的随机数
		result[i] = rng.Intn(v) + 1
	}
	return result
}

// 随机数生成器
func RandN(n int) int { // 【1, n】
	// 创建一个局部随机数生成器
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return rng.Intn(n) + 1
}

// 随机数验证器
func RandNValidator(nMax, vMax int) []int { // 返回一个 n 个元素， 元素值在 1-v 范围的切片\
	n := RandN(nMax)
	v := RandN(vMax)

	if n <= 0 {
		return []int{}
	}
	if v < 1 {
		return []int{}
	}

	// 创建一个局部随机数生成器
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	result := make([]int, n)
	for i := 0; i < n; i++ {
		// 使用局部随机数生成器生成 1 到 v 之间的随机数
		result[i] = rng.Intn(v) + 1
	}
	return result
}
