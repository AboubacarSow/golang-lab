package inventory

import (
	"encoding/json"
	"os"
	"path/filepath"
)

var (
	inventoryPath = "data/inventory.json"
)

func (items *Inventory) LoadInventory()  error {
	dir, _ := os.Getwd()
	path := filepath.Join(dir, inventoryPath)
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(data, items)
	

}

func (inv *Inventory) saveInventory() error {
	data, err := json.Marshal(*inv)
	if err != nil {
		return err
	}
	return os.WriteFile(inventoryPath, data, 0644)
}
