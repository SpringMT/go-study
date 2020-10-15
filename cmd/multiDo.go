package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	fmt.Println("start multiDo")
	wg := sync.WaitGroup{}
	num := 10
	done := make(chan interface{})
	wg.Add(num)
	multiTasks(&wg, done, num)
	time.Sleep(time.Duration(5) * time.Second)
	close(done)
	wg.Wait()
	fmt.Println("DONE!")
}

func multiTasks(wg *sync.WaitGroup, done <-chan interface{}, num int) {
	for i := 0; i < num; i++ {
		go taskController(wg, done, i)
	}
	fmt.Println("generate groutine")
}

func taskController(wg *sync.WaitGroup, done <-chan interface{}, i int) {
	defer func() {
		wg.Done()
	}()
	go task()
	for {
		select {
		// 中断
		case <-done:
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
func task() {
	fmt.Print(1)
}