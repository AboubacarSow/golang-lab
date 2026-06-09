package inventory

import (
	"errors"
)

func (inventory *Inventory) DeleteItem(index int) error {
	if index >= 0 && len(*inventory) > index {
		*inventory = append((*inventory)[:index], (*inventory)[index+1:]...)
		return inventory.saveInventory()

	}
	return errors.New("Invalid index")

}

func (inventory *Inventory) AddItem(name string, stock int) error {
	item := item{Name: name, Stock: stock}
	*inventory = append(*inventory, item)
	return inventory.saveInventory()
}
