package test

import (
	"fmt"
	"testing"
	"time"
)

// goroutine 示例
func TestHelloGoroutine(t *testing.T) {
	// 异步执行
	go sayHello()
	// 循环打印Main
	for range 5 {
		fmt.Println("Main")
		time.Sleep(100 * time.Millisecond)
	}
}

// 循环5次打印Hello
func sayHello() {
	for range 5 {
		fmt.Println("Hello")
		time.Sleep(100 * time.Millisecond)
	}
}
