package hunt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSharkHuntsSuccessfully(t *testing.T) {
	//Arrange
	shark := NewShark(true, false, 6)
	prey := NewPrey("presa", 5)

	//Act
	err := shark.Hunt(&prey)

	//Assert
	assert.NoError(t, err)
	assert.False(t, shark.hungry)
	assert.True(t, shark.tired)
}

func TestSharkCannotHuntBecauseIsTired(t *testing.T) {
	//Arrange
	shark := NewShark(true, true, 6)
	prey := NewPrey("presa", 5)

	//Act
	err := shark.Hunt(&prey)

	//Assert
	assert.NotNil(t, err)
	assert.EqualError(t, err, "cannot hunt, i am really tired")
}

func TestSharkCannotHuntBecaisIsNotHungry(t *testing.T) {
	//Arrange
	shark := NewShark(false, false, 6)
	prey := NewPrey("presa", 5)

	//Act
	err := shark.Hunt(&prey)

	//Assert
	assert.NotNil(t, err)
	assert.EqualError(t, err, "cannot hunt, i am not hungry")
}

func TestSharkCannotReachThePrey(t *testing.T) {
	//Arrange
	shark := NewShark(true, false, 5)
	prey := NewPrey("presa", 5)

	//Act
	err := shark.Hunt(&prey)

	//Assert
	assert.NotNil(t, err)
	assert.EqualError(t, err, "could not catch it")
	assert.True(t, shark.tired)
}

func TestSharkHuntNilPrey(t *testing.T) {
	shark := NewShark(true, false, 5)

	//Act
	err := shark.Hunt(nil)

	//Assert
	assert.NotNil(t, err)
	assert.EqualError(t, err, "there is no prey to catch")
}
