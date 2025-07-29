#!/bin/bash

# 构建Rust动态库
echo "构建Rust动态库..."
cargo build --release

# 根据操作系统复制动态库文件
if [[ "$OSTYPE" == "darwin"* ]]; then
    # macOS
    cp target/release/librust_add_ffi.dylib ./librust_add_ffi.dylib
    echo "已复制 librust_add_ffi.dylib"
elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
    # Linux
    cp target/release/librust_add_ffi.so ./librust_add_ffi.so
    echo "已复制 librust_add_ffi.so"
elif [[ "$OSTYPE" == "msys" ]] || [[ "$OSTYPE" == "win32" ]]; then
    # Windows
    cp target/release/rust_add_ffi.dll ./rust_add_ffi.dll
    echo "已复制 rust_add_ffi.dll"
fi

echo "构建完成！"