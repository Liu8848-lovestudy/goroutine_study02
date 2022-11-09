package main

import (
	"fmt"
	"goroutine_study02/restuarant"
)

func main() {
	m, n, c := 50, 10, 3 //材料，顾客，厨师数量
	rf := restuarant.ReFactory{}
	restuarant.Set(m, n, c)
	ch := make(chan struct{}, 1)
	for i := 1; i < c+1; i++ {

		go rf.Create("cooker", i).Action(ch)
	}
	for j := 1; j < n+1; j++ {
		go rf.Create("customer", j).Action(ch)
	}
	for {
		restuarant.Judge(n, c, ch)
		select {
		case <-ch:
			fmt.Println("饭店打样！")
			restuarant.Close_All()
			return
		default:
			break
		}

	}
}
