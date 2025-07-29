# Rust Add FFI

[![CI](https://github.com/your-username/rust-add-ffi/workflows/CI/badge.svg)](https://github.com/your-username/rust-add-ffi/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/your-username/rust-add-ffi)](https://goreportcard.com/report/github.com/your-username/rust-add-ffi)
[![GoDoc](https://godoc.org/github.com/your-username/rust-add-ffi?status.svg)](https://godoc.org/github.com/your-username/rust-add-ffi)

这个项目演示了如何用Rust实现函数并通过FFI（Foreign Function Interface）让Go语言调用，并提供完整的CI/CD流程。

## 安装

```bash
go get github.com/your-username/rust-add-ffi
```

## 使用方法

```go
package main

import (
    "fmt"
    rustffi "github.com/your-username/rust-add-ffi"
)

func main() {
    // 整数加法
    result := rustffi.Add(10, 20)
    fmt.Printf("Add(10, 20) = %d\n", result)
    
    // 浮点数加法
    floatResult := rustffi.AddFloat(3.14, 2.86)
    fmt.Printf("AddFloat(3.14, 2.86) = %.2f\n", floatResult)
}
```

## 项目结构

```
.
├── .github/
│   └── workflows/
│       ├── ci.yml          # 持续集成
│       └── release.yml     # 发布流程
├── src/
│   └── lib.rs              # Rust库实现
├── Cargo.toml              # Rust项目配置
├── go.mod                  # Go模块配置
├── lib.go                  # Go FFI封装
├── add_test.go             # Go测试文件
├── main.go                 # 示例程序
├── build.sh                # 构建脚本
└── README.md               # 说明文档
```

## 开发

### 本地构建

```bash
# 构建Rust动态库
./build.sh

# 运行Go测试
go test -v

# 运行示例程序
go run main.go
```

### 发布新版本

1. 创建并推送标签：
```bash
git tag v1.0.0
git push origin v1.0.0
```

2. GitHub Action会自动：
   - 运行测试
   - 构建跨平台动态库
   - 创建GitHub Release
   - 上传构建产物

## API文档

### Add(a, b int) int
整数加法函数，调用Rust实现。

### AddFloat(a, b float64) float64
浮点数加法函数，调用Rust实现。

## 系统要求

- Go 1.20+
- Rust 1.70+
- CGO支持

## 许可证

MIT License