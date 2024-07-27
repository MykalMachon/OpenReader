package main

import (
	"fmt"

	"github.com/mykalmachon/openreader/internal/tools"
	"github.com/mykalmachon/openreader/pkg/database"
)

func main() {
	fmt.Println("hello world")
	sum := tools.Adder(1, 1)
	fmt.Println("1 + 1 is ", sum)
	db, err := database.Init("openreader.sqlite")

	if err != nil {
		fmt.Println("something went wrong when initializing the db", err)
	}

	defer db.Close()
}
