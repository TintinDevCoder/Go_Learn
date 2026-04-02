package main

import "fmt"

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// 求切片最大值
func Max[T Number](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}

	max := slice[0]
	for _, value := range slice[1:] {
		if value > max {
			max = value
		}
	}
	return max
}

// 求切片最小值
func Min[T Number](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}

	min := slice[0]
	for _, value := range slice[1:] {
		if value < min {
			min = value
		}
	}
	return min
}

// 求切片平均值
func Average[T Number](slice []T) float64 {
	if len(slice) == 0 {
		return 0
	}

	var sum T
	for _, value := range slice {
		sum += value
	}
	return float64(sum) / float64(len(slice))
}

// 使用示例
func main() {
	ints := []int{1, 5, 3, 9, 2}
	floats := []float64{1.1, 5.5, 3.3, 9.9, 2.2}

	fmt.Printf("Max int: %d\n", Max(ints))         // 输出: 9
	fmt.Printf("Min float: %.1f\n", Min(floats))   // 输出: 1.1
	fmt.Printf("Average: %.2f\n", Average(floats)) // 输出: 4.40
}
