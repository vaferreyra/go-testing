package hunt

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestSharkHuntsSuccessfully(t *testing.T) {
}

func TestSharkCannotHuntBecauseIsTired(t *testing.T) {
	// Arrange.
	shark := Shark{
		true,
		true,
		100,
	}
	prey := Prey{
		"Golden fish",
		120,
	}

	// Act.
	err := shark.Hunt(&prey)

	// Assert.
	assert.Equal(t, err, ErrTiredShark)
}

func TestSharkCannotHuntBecaisIsNotHungry(t *testing.T) {
	// Arrange.
	shark := Shark{
		false,
		false,
		100,
	}
	prey := Prey{
		"Golden fish",
		120,
	}

	// Act.
	err := shark.Hunt(&prey)

	// Assert.
	assert.Equal(t, err, ErrSharkNotHungry)
}

func TestSharkCannotReachThePrey(t *testing.T) {
	// Arrange.
	shark := Shark{
		true,
		false,
		100,
	}
	prey := Prey{
		"Golden fish",
		120,
	}

	// Act.
	err := shark.Hunt(&prey)

	// Assert.
	assert.Equal(t, err, ErrCantCatch)
}

func TestSharkHuntNilPrey(t *testing.T) {
}
