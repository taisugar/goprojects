package storage

import (
	"encoding/csv"
	"os"
	"strconv"
	"task-manager/internal/domain"
	"time"
)

type CSV struct{ FileName string }

func NewCSV(file string) *CSV { return &CSV{FileName: file} }

func (s *CSV) Save(tasks domain.Tasks) error {
	f, _ := os.Create(s.FileName)
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	w.Write([]string{"Title", "Description", "CreatedAt", "IsCompleted", "CompletedAt"})

	for _, t := range tasks {
		ca := ""
		if t.CompletedAt != nil {
			ca = t.CompletedAt.Format(time.RFC3339)
		}
		w.Write([]string{
			t.Title,
			t.Description,
			t.CreatedAt.Format(time.RFC3339),
			strconv.FormatBool(t.IsCompleted),
			ca,
		})
	}
	return nil
}

func (s *CSV) Load(data *domain.Tasks) error {
	f, err := os.Open(s.FileName)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer f.Close()

	rows, _ := csv.NewReader(f).ReadAll()

	for i, r := range rows {
		if i == 0 {
			continue
		}
		created, _ := time.Parse(time.RFC3339, r[2])
		done, _ := strconv.ParseBool(r[3])

		var ca *time.Time
		if r[4] != "" {
			t, _ := time.Parse(time.RFC3339, r[4])
			ca = &t
		}

		*data = append(*data, domain.Task{
			Title: r[0], Description: r[1],
			CreatedAt:   created,
			IsCompleted: done,
			CompletedAt: ca,
		})
	}
	return nil
}
