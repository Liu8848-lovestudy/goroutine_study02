package restuarant

import (
	"fmt"
	"sync"
	"time"
)

type Cooker struct {
	id      int
	m_speed int
	wg      *sync.WaitGroup
}

var (
	flag1 = true
	lock1 sync.Mutex
)

func (c Cooker) Action() {
	defer func() {
		err := recover()
		if err != nil {
			err = nil
		}
	}()
	for _, flag := <-ch_mat; flag; {
		if len(ch_num) != n1 {
			ch_shou <- 1
			time.Sleep(time.Duration(c.m_speed) * time.Millisecond * 800)
			fmt.Printf("厨师%d号，做了一个寿司\n", c.id)
		} else {
			break
		}
	}

	lock1.Lock()
	if flag1 {
		flag1 = false
		close(ch_shou)
	}
	lock1.Unlock()
	c.wg.Done()
}
