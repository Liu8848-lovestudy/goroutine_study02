package restuarant

var ch_mat chan int              //材料总数
var ch_num chan int              //客人总数
var ch_shou = make(chan int, 20) //寿司管道
var m1 int
var n1 int

func Set(m int, n int) {
	ch_mat = make(chan int, m)
	ch_num = make(chan int, n)
	m1 = m
	n1 = n
	for i := 0; i < m1; i++ {
		ch_mat <- 1
	}

}
