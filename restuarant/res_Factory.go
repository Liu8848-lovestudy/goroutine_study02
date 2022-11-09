package restuarant

type ReFactory struct {
}

func (r ReFactory) Create(types string, i int) person {
	if types == "cooker" {
		return Cooker{
			id:      100 + i,
			m_speed: i,
			m_num:   10 * i,
		}
	} else if types == "customer" {
		return Customer{
			id:      1000 + i,
			e_speed: i,
			e_num:   3 * i,
		}
	} else {
		return nil
	}
}
