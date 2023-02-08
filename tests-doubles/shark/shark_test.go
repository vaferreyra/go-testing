package shark

import (
	"errors"
	"testdoubles/prey"
	"testdoubles/prey/prey_mock"
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

		// Act
		err := whiteShark.Hunt(prey)

		// Assert
		assert.Error(t, err)
		assert.True(t, errors.Is(err, ErrCantHuntThePrey))
		assert.True(t, simulatorMock.Spy["CanCatch"])
		assert.True(t, simulatorMock.Spy["GetLinearDistance"])
	})

	t.Run("Hunt prey being fastly and short distance", func(t *testing.T) {
		// Arrange.
		simulatorMock := mocks.NewSimulatorMock()
		simulatorMock.Reset()
		simulatorMock.Can_Catch = true
		simulatorMock.LinearDistance = 10

		whiteShark := CreateWhiteShark(simulatorMock)
		prey := prey_mock.NewPreyStub()

		// Act.
		err := whiteShark.Hunt(prey)

		// Assert.
		assert.NoError(t, err)
		assert.True(t, simulatorMock.Spy["GetLinearDistance"])
	})

	t.Run("Shark cannot hunt prey for being slowlier", func(t *testing.T) {
		simulatorMock := mocks.NewSimulatorMock()
		simulatorMock.Reset()
		simulatorMock.LinearDistance = 10

		whiteShark := CreateWhiteShark(simulatorMock)
		prey := prey_mock.NewPreyStub()

		prey.NewSpeed(300)

		// Act.
		err := whiteShark.Hunt(prey)

		// Assert.
		assert.Error(t, err)
		assert.True(t, errors.Is(err, ErrCantHuntThePrey))
		assert.True(t, simulatorMock.Spy["GetLinearDistance"])
	})
}
