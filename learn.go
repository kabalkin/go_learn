package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	wg := sync.WaitGroup{}
	data := make(map[string]string)
	b := true

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			cond.L.Lock()
			fmt.Printf("%d start wait\n", i)
			if b {
				cond.Wait()
			}
			fmt.Printf("%d read %s\n", i, data["key"])
			cond.L.Unlock()
		}(i)
	}
	time.Sleep(1 * time.Second)

	for i := 0; i < 2; i++ {
		cond.L.Lock()
		data["key"] = fmt.Sprintf("data-%d", i)
		cond.Broadcast()
		cond.L.Unlock()
		time.Sleep(1 * time.Second)
	}
	time.Sleep(1 * time.Second)
	b = false

	wg.Wait()
}
