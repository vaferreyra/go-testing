package prey_mock

type preyStub struct {
	maxSpeed float64
}

func NewPreyStub() *preyStub {
	return &preyStub{
		maxSpeed: 100,
	}
}

func (ps *preyStub) GetSpeed() (speed float64) {
	return ps.maxSpeed
}

func (ps *preyStub) NewSpeed(speed float64) {
	ps.maxSpeed = speed
}
