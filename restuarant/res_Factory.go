package restuarant

import "sync"

type ReFactory struct {
}

func (r ReFactory) Create(types string, i int, wg1 *sync.WaitGroup) person {
	if types == "cooker" {
		return Cooker{
			id:      100 + i,
			m_speed: i,
			wg:      wg1,
		}
	} else if types == "customer" {
		return Customer{
			id:      100 + i,
			e_speed: i,
			wg:      wg1,
			e_num:   3,
		}
	} else {
		return nil
	}
}
