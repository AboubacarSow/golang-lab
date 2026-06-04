package todocli

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type item struct {
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at,omitempty"`
	Task        string    `json:"task"`
}

type Todos []item

func (t *Todos) Add(task string) {
	todo := item{
		Done:        false,
		CreatedAt:   time.Now(),
		Task:        task,
		CompletedAt: time.Time{},
	}
	*t = append(*t, todo)
}
func (t *Todos) Complete(index int) error {
	list := *t
	if index-1 < 0 || index-1 >= len(list) {
		return fmt.Errorf("invalid index")
	}
	(list)[index-1].Done = true
	(list)[index-1].CompletedAt = time.Now()
	return nil
}
func (t *Todos) Delete(index int) error {
	list := *t
	if index-1 < 0 || index-1 >= len(list) {
		return errors.New("invalid index")
	}
	*t = append((list)[:index-1], (list)[index:]...)
	return nil
}

func (t *Todos) List() []item {
	return *t
}
func (t *Todos) LoadFromFile(filename string) error {
	path := filepath.Join(os.TempDir(), filename)
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	if len(data) == 0 {
		return nil
	}
	err = json.Unmarshal(data, t)
	if err != nil {
		return err
	}
	return nil
}

func (t *Todos) SaveToFile(filename string) error {
	path := filepath.Join(os.TempDir(), filename)
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
