package main

import (
	"bufio"
	"fmt"
	"go_project/pipeline"
	"os"
)

func mergeDemo() {
	p := pipeline.Merge(
		pipeline.InMemSort(
			pipeline.ArraySource(3, 2, 6, 7, 4)),
		pipeline.InMemSort(
			pipeline.ArraySource(7, 4, 0, 3, 2, 13, 8)))
	//for {
	//	if num, ok := <-p; ok {
	//		fmt.Println(num)
	//	} else {
	//		break
	//	}
	//num, ok := <-p
	//if ok {
	//	fmt.Println(num)
	//} else {
	//	break
	//}
	//}
	for v := range p {
		fmt.Println(v)
	}
}

func main() {
	//const filename = "large.in"
	//const n = 100000000
	const filename = "small.in"
	const n = 64
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	p := pipeline.RandomSource(n)
	writer := bufio.NewWriter(file)
	pipeline.WriteSink(writer, p)
	writer.Flush()
	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p = pipeline.ReadSource(bufio.NewReader(file), -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count ++
		if count >= 100 {
			break
		}
	}
}
