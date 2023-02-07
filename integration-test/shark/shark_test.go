package shark

import (
	"integrationtests/pkg/storage/mocks"
	"testing"
)

func TestHunt_Ok(t *testing.T) {
	// Arrange.
	storage := mocks.NewMockStorage()
	storage.On("GetValue", "white_shark_speed").Return(200)
	storage.On("GetValue", "white_shark_x").Return(0)
	storage.On("GetValue", "white_shark_y").Return(0)
}
