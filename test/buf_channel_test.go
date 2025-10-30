package test

import (
	"fmt"
	"testing"
)

func TestBufChannel(t *testing.T) {
	// 创建缓冲区大小为2的整型通道
	ch := make(chan int, 2)

	// 向缓冲通道发送两个值（不会立即阻塞）
	ch <- 1
	ch <- 2

	// 从通道接收并打印值
	fmt.Println(<-ch) // 输出: 1
	fmt.Println(<-ch) // 输出: 2
}
