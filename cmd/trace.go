package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()

	// 2がないと、channelへの書き込みがブロックされる
	resultStream := make(chan error, 3)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		// 処理A
		time.Sleep(10)
		// resultStream <- fmt.Errorf("Error 1")
	}()
	fmt.Println(runtime.NumGoroutine())
	wg.Add(1)
	go func() {
		defer wg.Done()
		// 処理B
		// resultStream <- fmt.Errorf("Error 2")
	}()
	fmt.Println(runtime.NumGoroutine())
	wg.Add(1)
	go func() {
		defer wg.Done()
		// 処理C 成功した
		// resultStream <- nil
	}()
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
	close(resultStream)
	for result := range resultStream {
		fmt.Printf("Error %v\n", result)
	}
	fmt.Printf("DONE")
}

