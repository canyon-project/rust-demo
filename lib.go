// Package rustaddffi provides Go bindings for Rust add functions via FFI
package main

/*
#cgo LDFLAGS: -L. -lrust_add_ffi
#include <stdlib.h>

// 声明Rust函数
int add(int a, int b);
double add_float(double a, double b);
*/
import "C"

// Add 调用Rust实现的整数加法函数
func Add(a, b int) int {
	return int(C.add(C.int(a), C.int(b)))
}

// AddFloat 调用Rust实现的浮点数加法函数
func AddFloat(a, b float64) float64 {
	return float64(C.add_float(C.double(a), C.double(b)))
}
