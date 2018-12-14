package main

import (
	"fmt"
	//	"time"
)

/* goroutine
func f(str string) {
	for i := 0; i < 10; i++ {
		fmt.Println(str, i)
	}
}

func main() {
	go f("test")

	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}
*/

/* //timeout
func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout after 1 second 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout after 1 second 2")
	}
}
*/

// channel
func main() {
	msg := make(chan string, 10)

	go func() {
		msg <- "duang!"
	}()

	message := <-msg
	fmt.Println(message)
}
