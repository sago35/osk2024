package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Printf("hello world!\n")
		time.Sleep(1000 * time.Millisecond)
	}
}
