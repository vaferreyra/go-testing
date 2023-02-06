package prey

import (
	mocks "testdoubles/prey/prey_mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSpeed(t *testing.T) {
	t.Run("Get Speed with Dummy Prey", func(t *testing.T) {
		// El test siempre funciona, pero con valores por default
		tuna := mocks.NewPreyDummy()

		speed := tuna.GetSpeed()

		assert.Equal(t, speed, float64(0))
	})

	t.Run("Get Speed with Stub Prey", func(t *testing.T) {
		// El test siempre funciona con valores hardcodeados.
		tuna := mocks.NewPreyStub()

		speed := tuna.GetSpeed() // Siempre retorna 100.0

		assert.Equal(t, speed, 100.0)
	})
}
