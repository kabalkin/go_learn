<div id='up'/>

# Table of Contents
1. [Concurrency](#concurrency)
   1. [sync.Cond](#synccond)
   2. sync.WaitGroup
   3. sync.Mutex, sync.RWMutex
   4. [sync.Once](#synconce)
2. [To Watch](#towatch)
<div id='concurrency'/>


### Concurrency [^](#up)
<div id='synccond'/>

* sync.Cond (Signal, Broadcast) -  уведомить разные горутины, (важно cond.L.Lock/Unlock). Вообще принято не использовать, возможно будет 
удалено в версии 2go

```
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

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			cond.L.Lock()
			fmt.Printf("%d start wait\n", i)
			cond.Wait()
			fmt.Printf("%d read %s\n", i, data["key"])
			cond.L.Unlock()
		}(i)
	}

	time.Sleep(1 * time.Second)
	data["key"] = fmt.Sprintf("data")
	cond.Broadcast()

	wg.Wait()
}
```
---
* sync.Mutex, sync.RWMutex (аналог lock)
---
<div id = "synconce"/>

* sync.Once [^](#up)

---
>At the command prompt, type

<div id='towatch'/>



### ToWatch [^](#up)

[Concurrency is not parallelism](https://habr.com/ru/articles/761754/)

[Go blog](https://go.dev/blog/all)

[CPS Go](https://github.com/Q69K/using-cps-in-golang?ysclid=lovqsgfaae78260087)

[MutexOrChannel](https://github.com/golang/go/wiki/MutexOrChannel)

[design-patterns/go](https://refactoring.guru/ru/design-patterns/go)

