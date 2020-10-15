package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Duration(i) * time.Second)
			fmt.Printf("time: %s, goroutine: %d\n", time.Now(), i)
		} (i)
	}
	wg.Wait()
	fmt.Println("DONE!")
}
