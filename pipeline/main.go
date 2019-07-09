package main

import (
	"fmt"
	"go_project/pipeline"
)

func main() {
	p := pipeline.ArraySource(3, 2, 6, 7, 4)
	for {
		num := <-p
		fmt.Println(num)
	}
}
