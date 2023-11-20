package main

import (
	"fmt"
	"sync"
)

func main() {
	type Button struct {
		Clicked *sync.Cond
	}

	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fn func()) {
		var running sync.WaitGroup
		running.Add(1)
		go func() {
			running.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		running.Wait()
	}

	var clicked sync.WaitGroup
	clicked.Add(3)
	subscribe(button.Clicked, func() {
		fmt.Println("Fn1 end")
		clicked.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Fn2 end")
		clicked.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Fn3 end")
		clicked.Done()
	})

	button.Clicked.Broadcast()
	clicked.Wait()

	clicked.Add(3)
	subscribe(button.Clicked, func() {
		fmt.Println("Fn1 end")
		clicked.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Fn2 end")
		clicked.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Fn3 end")
		clicked.Done()
	})
	button.Clicked.Broadcast()
	clicked.Wait()
}
