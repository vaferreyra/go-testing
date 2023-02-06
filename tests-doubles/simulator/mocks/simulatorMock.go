package mocks

func NewSimulatorMock() *simulatorMock {
	return &simulatorMock{}
}

type simulatorMock struct {
	Can_Catch      bool
	LinearDistance float64
	Spy            map[string]bool
}

func (mock *simulatorMock) CanCatch(distance float64, speed float64, catchSpeed float64) bool {
	mock.Spy["CanCatch"] = true
	return mock.Can_Catch
}

func (mock *simulatorMock) GetLinearDistance(position [2]float64) float64 {
	mock.Spy["GetLinearDistance"] = true
	return mock.LinearDistance
}

func (mock *simulatorMock) Reset() {
	mock.Can_Catch = false
	mock.LinearDistance = 0
	mock.Spy = make(map[string]bool)
}
