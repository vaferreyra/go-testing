package mocks

type FakeStorage interface {
	GetValue(key string) interface{}
}

type fakeStorage struct {
	ResultOnGet interface{}
	ErrOnGet    error
}

func NewFakeStorage() FakeStorage {
	return &fakeStorage{}
}

func (fs *fakeStorage) GetValue(key string) interface{} {
	if fs.ErrOnGet != nil {
		return fs.ErrOnGet
	}

	return fs.ResultOnGet
}
