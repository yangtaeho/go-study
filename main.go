package main

import (
    "fmt"
    "os"
)

var print = fmt.Println

func main() {
    print("Hello World")

    // 일시 정지를 위해
    var b []byte = make([]byte, 1)
    os.Stdin.Read(b)
} 