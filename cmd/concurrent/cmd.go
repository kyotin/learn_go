package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan string, 2)

	var cwg sync.WaitGroup
	for i := 0; i < 2; i++ {
		cwg.Add(1)
		go func(workerId int, wg *sync.WaitGroup) {
			for v := range ch {
				fmt.Printf("consumer %d: glad to meet: %s \n", workerId, v)
			}
			wg.Done()
		}(i, &cwg)
	}

	var pwg sync.WaitGroup
	for i := 0; i < 10; i++ {
		pwg.Add(1)
		go func(workerId int, wg *sync.WaitGroup) {
			ch <- fmt.Sprintf("my name is producer %d \n", workerId)
			wg.Done()
		}(i, &pwg)
	}

	pwg.Wait()
	close(ch)
	cwg.Wait()
}
