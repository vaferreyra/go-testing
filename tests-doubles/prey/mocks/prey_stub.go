package mocks

type preyStub struct {
	maxSpeed float64
}

func NewPreyStub() *preyStub {
	return &preyStub{}
}

func (ps *preyStub) GetSpeed() (speed float64) {
	return 100
}
