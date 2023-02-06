package mocks

type preyDummy struct {
	maxSpeed float64
}

func NewPreyDummy() *preyDummy {
	return &preyDummy{}
}

func (pd *preyDummy) GetSpeed() (speed float64) {
	return
}
