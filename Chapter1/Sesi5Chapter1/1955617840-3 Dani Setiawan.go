package main

import (
	"fmt"
	"sync"
)

func main() {
	interface1 := []interface{}{
		"coba1",
		"coba2",
		"coba3",
	}
	interface2 := []interface{}{
		"bisa1",
		"bisa2",
		"bisa3",
	}
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 1; i <= 4; i++ {
		wg.Add(2)
		go func(index int) {
			m.Lock()
			fmt.Println(interface1, index)
			m.Unlock()
			wg.Done()
		}(i)

		go func(index int) {
			m.Lock()
			fmt.Println(interface2, index)
			m.Unlock()
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("=====================")

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go func(index int) {
			fmt.Println(interface1, index)
			wg.Done()
		}(i)
	}

	for j := 1; j <= 4; j++ {
		wg.Add(1)
		go func(index int) {
			fmt.Println(interface2, index)
			wg.Done()
		}(j)
	}
	wg.Wait()
}
