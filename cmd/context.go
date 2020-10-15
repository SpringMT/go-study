package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main()  {
	fmt.Println("start multiDo")
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	num := 10
	wg.Add(num)
	multiTasks1(ctx, &wg, num)
	time.Sleep(time.Duration(5) * time.Second)
	cancel()
	wg.Wait()
	fmt.Println("DONE!")
}

func multiTasks1(ctx context.Context, wg *sync.WaitGroup, num int) {
	for i := 0; i < num; i++ {
		go taskController1(ctx, wg, i)
	}
	fmt.Println("generate groutine")
}

func taskController1(ctx context.Context, wg *sync.WaitGroup, i int) {
	defer func() {
		wg.Done()
	}()
	go task1()
	for {
		select {
		// 中断
		case <-ctx.Done():
			// 途中できりたい
			fmt.Printf("time: %s, goroutine: %d canceled\n", time.Now(), i)
			return
		// timeout
		case <- time.After(time.Duration(i) * time.Second):
			fmt.Printf("time: %s, goroutine: %d\n", time.Now(), i)
			return
		}
	}
}

// goroutine リークすると思われる
func task1() {
	fmt.Print(1)
}
