# Table of Contents
1. [Concurrency](#concurrency)
   1. [sync.Cond](#synccond)
   2. 

<div id='concurrency'/>

##### Concurrency
<div id='synccond'/>

* sync.Cond -  уведомить разные горутины, (важно cond.L.Lock/Unlock). Вообще принято не использовать, возможно будет 
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

>At the command prompt, type


### ToWatch

My favorite search engine is [Duck Duck Go](https://duckduckgo.com).