package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	for i := 0; i <= 5000; i++ {
		// go starts a goroutine
		go printHelloWorld(i, ch)
	}

	for {
		msg := <-ch
		fmt.Println(msg)
	}
	close(ch)

}

func printHelloWorld(i int, ch chan string) {
	ch <- fmt.Sprintf("Hello world! %d \n", i)
}
