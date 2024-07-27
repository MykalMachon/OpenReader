package main

import (
	"fmt"

	"github.com/mykalmachon/openreader/internal/tools"
)

func main() {
	fmt.Println("hello world")
	sum := tools.Adder(1, 1)
	fmt.Println("1 + 1 is ", sum)
}
