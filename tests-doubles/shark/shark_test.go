package shark

import (
	"errors"
	"testdoubles/prey"
	"testdoubles/simulator/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHunt(t *testing.T) {
	t.Run("Hunt happy path with simulator mock", func(t *testing.T) {
		// Arrange.
		simulatorMock := mocks.NewSimulatorMock()
		simulatorMock.Reset()
		simulatorMock.Can_Catch = true

		whiteShark := CreateWhiteShark(simulatorMock)
		prey := prey.CreateTuna()

		// Act.
		err := whiteShark.Hunt(prey)

		// Assert.
		assert.NoError(t, err)
		assert.True(t, simulatorMock.Spy["CanCatch"])
		assert.True(t, simulatorMock.Spy["GetLinearDistance"])
	})

	t.Run("Hunt prey error can't hunt", func(t *testing.T) {
		// Arrange.
		simulatorMock := mocks.NewSimulatorMock()
		simulatorMock.Reset()
		simulatorMock.Can_Catch = false

		whiteShark := CreateWhiteShark(simulatorMock)
		prey := prey.CreateTuna()

		// Act.
		err := whiteShark.Hunt(prey)

		// Assert
		assert.Error(t, err)
		assert.True(t, errors.Is(err, ErrCantHuntThePrey))
		assert.True(t, simulatorMock.Spy["CanCatch"])
		assert.True(t, simulatorMock.Spy["GetLinearDistance"])
	})

}
