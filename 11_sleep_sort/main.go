package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for _, s := range os.Args[1:] {
		n, _ := strconv.Atoi(s)
		wg.Add(1)
		go func(n int) {
			time.Sleep(time.Duration(n) * 100 * time.Millisecond)
			fmt.Printf("%d\n", n)
			wg.Done()
		}(n)
	}
	wg.Wait()
}
