package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			fmt.Printf("5\n")
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			fmt.Printf("   12\n")
			time.Sleep(1200 * time.Millisecond)
		}
	}()

	select {}
}
