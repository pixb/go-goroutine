package main

import (
	"fmt"
	"time"

	"github.com/spf13/pflag"
)

var name string

func init() {
	pflag.StringVarP(&name, "name", "n", "hello", `goroutine demo
hello: hello goroutine demo
channelsum: channel sum number demo
bufchannel: channel buffer demo`)
}

func sayHello() {
	for range 5 {
		fmt.Println("Hello")
		time.Sleep(100 * time.Millisecond)
	}
}

func helloGoroutine() {
	fmt.Println("helloGoroutine() run...")
	go sayHello()
	for range 5 {
		fmt.Println("Main")
		time.Sleep(100 * time.Millisecond)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

// 利用goroutine和channel来计算求和Demo
func channelSum() {
	fmt.Println("channelSum(), Run...")
	s := []int{7, 2, 8, -9, 4, 0}
	// 创建channel
	c := make(chan int)
	// 计算数组前一半元素之和
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Printf("channelSub(), x=%d, y=%d, x+y=%d\n", x, y, x+y)
}

func bufChannel() {
	// 创建缓冲区大小为2的整型通道
	ch := make(chan int, 2)

	// 向缓冲通道发送两个值（不会立即阻塞）
	ch <- 1
	ch <- 2

	// 从通道接收并打印值
	fmt.Println(<-ch) // 输出: 1
	fmt.Println(<-ch) // 输出: 2
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for range n {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// 通道遍历与关闭
func channelForClose() {
	fmt.Println("channelForClose(), Run...")
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range c {
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("--- go-goroutine main ---")
	pflag.Parse()
	switch name {
	case "hello":
		helloGoroutine()
	case "channelsum":
		channelSum()
	case "bufchannel":
		bufChannel()
	case "channelforclose":
		channelForClose()
	default:
		pflag.Usage()
	}
}
