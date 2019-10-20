package main

import (
	"fmt"

	"github.com/sueken5/golang-ddd/pkg/dddshop"
)

func main() {
	fmt.Println("hello world")

	if err := dddshop.Execute(); err != nil {
		panic(err)
	}
}
