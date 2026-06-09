package main

import (
	"fmt"
	"inventory/internal/inventory"
	"inventory/internal/ui"
)

func main() {

	application := ui.NewApp()

	inventoryItems := &inventory.Inventory{}

	if err := inventoryItems.LoadInventory(); err != nil {
		fmt.Printf("Error occured while loading:%v\n", err)
		return
	}

	if err := application.Run(inventoryItems); err != nil {
		panic(err)
	}
}
