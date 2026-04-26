package storage

import (
	"encoding/json"
	"os"
)

type JSON[T any] struct {
	FileName string
}

func NewJSON[T any](file string) *JSON[T] {
	return &JSON[T]{FileName: file}
}

func (s *JSON[T]) Save(data T) error {
	b, _ := json.MarshalIndent(data, "", "	")
	return os.WriteFile(s.FileName, b, 0644)
}

func (s *JSON[T]) Load(data *T) error {
	b, err := os.ReadFile(s.FileName)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(b, data)
}
