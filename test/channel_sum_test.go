package test

import (
	"fmt"
	"testing"
)

// 利用goroutine和channel来计算求和Demo
func TestChannelSum(t *testing.T) {
	s := []int{7, 2, 8, -9, 4, 0}
	// 创建channel
	c := make(chan int)
	// 计算数组前一半元素之和
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Printf("channelSub(), x=%d, y=%d, x+y=%d\n", x, y, x+y)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}
