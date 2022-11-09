package restuarant

import (
	"fmt"
	"time"
)

type Cooker struct {
	id      int
	m_speed int //厨师做饭速度
	m_num   int //厨师做饭产量
}

//var (
//	flag1 = true
//	lock1 sync.Mutex
//)

func (c Cooker) Action(ch chan struct{}) {
	//defer func() {
	//	err := recover()
	//	if err != nil {
	//		err = nil
	//	}
	//}()
	//for _, flag := <-ch_mat; flag; {
	//	if len(ch_num) != n1 {
	//		ch_shou <- 1
	//		time.Sleep(time.Duration(c.m_speed) * time.Millisecond * 800)
	//		fmt.Printf("厨师%d号，做了一个寿司\n", c.id)
	//	} else {
	//		break
	//	}
	//}
	//for {
	//	select {
	//	case <-ch_mat:
	//		ch_shou <- 1
	//		time.Sleep(time.Duration(c.m_speed) * time.Millisecond * 800)
	//		fmt.Printf("厨师%d号，做了一个寿司\n", c.id)
	//	default:
	//
	//	}
	//}
	count := 0
	for c.m_num > 0 {
		select {
		case m := <-ch_mat:
			if m == 1 {
				time.Sleep(time.Duration(c.m_speed) * time.Millisecond * 800)
				ch_shou <- 1
				count++
				fmt.Printf("厨师%d号，一共做了%d个寿司\n", c.id, count)
				c.m_num--
			}
		default:
			fmt.Println("材料用尽！")
			ch <- struct{}{}
			return
		}
	}
	fmt.Printf("厨师%d号达到最大产量现在下班，一共做了%d个寿司\n", c.id, count)
	ch_conum <- struct{}{}
	//ch <- 1
	//c.wg.Done()
	//lock1.Lock()
	//if flag1 {
	//	flag1 = false
	//	close(ch_shou)
	//}
	//lock1.Unlock()
	//c.wg.Done()
}
