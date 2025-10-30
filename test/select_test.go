package test

import (
	"fmt"
	"testing"
)

func fibonacci_w(c chan int, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func TestSelect(t *testing.T) {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for range 10 {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci_w(c, quit)
}
