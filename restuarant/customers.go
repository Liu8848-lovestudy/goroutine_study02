package restuarant

import (
	"fmt"
	"sync"
	"time"
)

type Customer struct {
	id      int //顾客ID
	e_speed int //顾客吃饭速度
	e_num   int //顾客的饭量
	wg      *sync.WaitGroup
}

var (
	flag2 = true
	lock  sync.Mutex
)

func (c Customer) Action() {
	count := 0
	for data, flag := <-ch_shou; flag; {
		if data == 1 {
			count++
			time.Sleep(time.Duration(c.e_speed) * time.Millisecond * 500)
			fmt.Printf("顾客%d吃到了%d个寿司\n", c.id, count)
		}
		if count == c.e_num {
			fmt.Printf("顾客%d吃饱了,一共吃了%d个寿司\n", c.id, count)
			break
		}
	}
	ch_num <- 1
	lock.Lock()
	if len(ch_num) == n1 && flag2 == true {
		flag2 = false
		close(ch_mat)
	}
	lock.Unlock()
	c.wg.Done()
}
