#!/bin/bash

set -e

echo "构建Rust库..."
./build.sh

echo "运行Rust测试..."
cargo test

echo "运行Go测试..."
if [[ "$OSTYPE" == "darwin"* ]]; then
    export DYLD_LIBRARY_PATH=$PWD:$DYLD_LIBRARY_PATH
elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
    export LD_LIBRARY_PATH=$PWD:$LD_LIBRARY_PATH
fi

go test -v -race -coverprofile=coverage.out ./...

echo "运行Go vet..."
go vet ./...

echo "检查Go代码格式..."
if [ "$(gofmt -s -l . | wc -l)" -gt 0 ]; then
    echo "代码格式不正确:"
    gofmt -s -l .
    exit 1
fi

echo "运行示例程序..."
go run .

echo "所有测试通过！"