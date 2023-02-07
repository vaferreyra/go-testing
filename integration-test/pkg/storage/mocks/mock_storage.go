package mocks

import "github.com/stretchr/testify/mock"

type mockStorage struct {
	mock.Mock
}

func NewMockStorage() *mockStorage {
	return &mockStorage{}
}

func (ms *mockStorage) GetValue(key string) interface{} {
	args := ms.Called(key)
	return args.Get(0)
}
