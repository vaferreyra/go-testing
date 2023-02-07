package shark

import (
	"errors"
	"integrationtests/pkg/storage/mocks"
	"integrationtests/prey"
	"integrationtests/simulator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHunt_Ok(t *testing.T) {
	// Arrange.
	storage := mocks.NewMockStorage()
	storage.On("GetValue", "white_shark_speed").Return(20.0).
		On("GetValue", "white_shark_x").Return(250.0).
		On("GetValue", "white_shark_y").Return(250.0).
		On("GetValue", "tuna_speed").Return(10.0)

	simulator := simulator.NewMockCatchSimulator(100)
	shark := CreateWhiteShark(simulator, storage)
	prey := prey.CreateTuna(storage)

	// Act.
	err := shark.Hunt(prey)

	// Assert.
	assert.NoError(t, err)
	storage.AssertExpectations(t)
	assert.True(t, simulator.Spy["GetLinearDistance"])
}

func TestHunt_Error(t *testing.T) {
	t.Run("Shark cannot hut prey for breing slowlier", func(t *testing.T) {
		// Arrange.
		storage := mocks.NewMockStorage()
		storage.On("GetValue", "white_shark_speed").Return(10.0).
			On("GetValue", "white_shark_x").Return(250.0).
			On("GetValue", "white_shark_y").Return(250.0).
			On("GetValue", "tuna_speed").Return(20.0)

		simulator := simulator.NewMockCatchSimulator(100)
		shark := CreateWhiteShark(simulator, storage)
		prey := prey.CreateTuna(storage)

		// Act.
		err := shark.Hunt(prey)

		// Assert.
		assert.Error(t, err)
		assert.True(t, errors.Is(err, ErrCantCatch))
		assert.True(t, simulator.Spy["GetLinearDistance"])
	})

	t.Run("Shark cannot hunt prey for being far away", func(t *testing.T) {

		// Arrange.
		storage := mocks.NewMockStorage()
		storage.On("GetValue", "white_shark_speed").Return(20.0).
			On("GetValue", "white_shark_x").Return(1000.0).
			On("GetValue", "white_shark_y").Return(1000.0).
			On("GetValue", "tuna_speed").Return(10.0)

		simulator := simulator.NewMockCatchSimulator(100)
		shark := CreateWhiteShark(simulator, storage)
		prey := prey.CreateTuna(storage)

		// Act.
		err := shark.Hunt(prey)

		// Assert.
		assert.Error(t, err)
		assert.True(t, errors.Is(err, ErrCantCatch))
		assert.True(t, simulator.Spy["GetLinearDistance"])
	})
}
