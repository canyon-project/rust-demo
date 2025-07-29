package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -1, -2, -3},
		{"mixed numbers", -5, 10, 5},
		{"zero", 0, 0, 0},
		{"with zero", 5, 0, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestAddFloat(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive floats", 2.5, 3.5, 6.0},
		{"negative floats", -1.5, -2.5, -4.0},
		{"mixed floats", -2.5, 5.0, 2.5},
		{"zero", 0.0, 0.0, 0.0},
		{"with zero", 3.14, 0.0, 3.14},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AddFloat(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("AddFloat(%.2f, %.2f) = %.2f; want %.2f", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// 基准测试
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(100, 200)
	}
}

func BenchmarkAddFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddFloat(3.14, 2.86)
	}
}
