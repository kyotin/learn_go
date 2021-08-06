package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan string, 10)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(workerId int, wg *sync.WaitGroup) {
			ch <- fmt.Sprintf("i'm worker %d", workerId)
			wg.Done()
		}(i, &wg)
	}

	wg.Wait()
	close(ch)

	for w := range ch {
		fmt.Printf("received: %s \n", w)
	}

	fmt.Println("over loop range")

}
