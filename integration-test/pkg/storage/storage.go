package storage

import (
	"encoding/json"
	"errors"
	"os"
)

var (
	ErrInternal = errors.New("Something internal has wrong")
	ErrNotFound = errors.New("Cannot found")
)

type Storage interface {
	GetValue(key string) interface{}
}

type storage struct {
	file string
}

func (s *storage) GetValue(key string) interface{} {
	file, err := os.ReadFile(s.file)
	if err != nil {
		return ErrInternal
	}

	data := make(map[string]interface{})
	err = json.Unmarshal(file, &data)
	if err != nil {
		return ErrInternal
	}

	if v, ok := data[key]; ok {
		return v
	}

	return ErrNotFound
}

func NewStorage() Storage {
	file := "config.json"
	return &storage{file: file}
}
