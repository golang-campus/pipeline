package main

import (
	"gotest/pipeline"
	"fmt"
)

func main() {
	p := pipeline.ArraySource(3,2,6,7,4)
	for {
		if num,ok := <- p;ok{
			fmt.Println(num)
		}else {
			break
		}
	}
}
