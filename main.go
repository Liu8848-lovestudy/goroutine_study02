package main

import (
	"fmt"
	"goroutine_study02/restuarant"
	"sync"
)

func main() {
	m, n, c := 100, 5, 10
	rf := restuarant.ReFactory{}
	restuarant.Set(m, n)
	wg := sync.WaitGroup{}
	for i := 1; i < c+1; i++ {
		wg.Add(1)
		go rf.Create("cooker", i, &wg).Action()
	}
	for j := 1; j < n+1; j++ {
		wg.Add(1)
		go rf.Create("customer", j, &wg).Action()
	}
	wg.Wait()
	fmt.Println("饭店打样！")
}
