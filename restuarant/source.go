package restuarant

import "fmt"

var ch_mat chan int              //材料总数
var ch_cunum chan struct{}       //客人总数
var ch_conum chan struct{}       //厨师总数
var ch_shou = make(chan int, 20) //寿司管道
var m1 int
var n1 int

func Set(m, n, c int) {
	ch_mat = make(chan int, m)
	ch_cunum = make(chan struct{}, n)
	ch_conum = make(chan struct{}, c)
	m1 = m
	n1 = n
	for i := 0; i < m1; i++ {
		ch_mat <- 1
	}
}
func Judge(n, c int, ch chan struct{}) {
	if len(ch_cunum) == n {
		fmt.Println("顾客全部离席")
		ch <- struct{}{}
	}
	if len(ch_conum) == c {
		fmt.Println("厨师全部下班")
		ch <- struct{}{}
	}
}
func Close_All() {
	close(ch_shou)
	close(ch_conum)
	close(ch_cunum)
}
