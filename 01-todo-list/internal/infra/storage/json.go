package storage

import (
	"encoding/json"
	"os"
	"task-manager/internal/domain"
)

type JSON struct{ FileName string }

func NewJSON(file string) *JSON { return &JSON{FileName: file} }

func (s *JSON) Save(data domain.Tasks) error {
	b, _ := json.MarshalIndent(data, "", "	")
	return os.WriteFile(s.FileName, b, 0644)
}

func (s *JSON) Load(data *domain.Tasks) error {
	b, err := os.ReadFile(s.FileName)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(b, data)
}
