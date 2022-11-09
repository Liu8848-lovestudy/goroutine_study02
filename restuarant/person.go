package restuarant

type person interface {
	Action(ch chan struct{})
}
