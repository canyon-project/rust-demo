package main

import (
	"fmt"
)

func main() {
	fmt.Println("Go调用Rust FFI示例")
	fmt.Println("==================")

	// 使用封装的函数
	a, b := 10, 20
	result := Add(a, b)
	fmt.Printf("Add(%d, %d) = %d\n", a, b, result)

	// 调用浮点数加法函数
	x, y := 3.14, 2.86
	floatResult := AddFloat(x, y)
	fmt.Printf("AddFloat(%.2f, %.2f) = %.2f\n", x, y, floatResult)

	// 更多测试用例
	fmt.Println("\n更多测试用例:")
	testCases := []struct {
		a, b int
	}{
		{1, 2},
		{-5, 10},
		{0, 0},
		{100, -50},
	}

	for _, tc := range testCases {
		result := Add(tc.a, tc.b)
		fmt.Printf("Add(%d, %d) = %d\n", tc.a, tc.b, result)
	}
}
