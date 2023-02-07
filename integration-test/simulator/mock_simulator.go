package simulator

import "math/big"

type mockCatchSimulator struct {
	MaxTimeToCatch float64
	Spy            map[string]bool
}

func NewMockCatchSimulator(maxTimeToCatch float64) *mockCatchSimulator {
	return &mockCatchSimulator{
		MaxTimeToCatch: maxTimeToCatch,
		Spy:            make(map[string]bool),
	}
}

func (ms *mockCatchSimulator) CanCatch(distance float64, speed float64, catchSpeed float64) bool {
	ms.Spy["CanCatch"] = true

	timeToCatch := distance / (speed - catchSpeed)
	return timeToCatch > 0 && timeToCatch <= ms.MaxTimeToCatch
}

func (ms *mockCatchSimulator) GetLinearDistance(position [2]float64) float64 {
	ms.Spy["GetLinearDistance"] = true

	x := big.NewFloat(position[0])
	y := big.NewFloat(position[1])
	z := x.Add(x.Mul(x, x), y.Mul(y, y))
	res, _ := z.Sqrt(z).Float64()
	return res
}
