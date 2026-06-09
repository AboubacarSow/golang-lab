package ui

import (
	"fmt"
	"inventory/internal/inventory"
	"strconv"

	"github.com/rivo/tview"
)

type App struct {
	App *tview.Application
}

func NewApp() *App {
	return &App{
		App: tview.NewApplication(),
	}
}

func (application *App) Run(items *inventory.Inventory) error {

	itemNameInput := tview.NewInputField().SetLabel("Item Name")
	itemStockInput := tview.NewInputField().SetLabel("Item Stock")
	itemIndexForDelete := tview.NewInputField().SetLabel("Index item")

	inventory := BuildInventoryView(items)
	inventory.RefereshView()	
	form := BuildForm(itemNameInput,
		itemStockInput,
		itemIndexForDelete,
		inventory.addHandler,
		inventory.deleteHandler,
		application.App.Stop)

	layout := BuildLayout(inventory.View, form.Form)
	return application.App.SetRoot(layout.Layout, true).Run()
}

func (inv *inventoryView) addHandler(name string, stockStr string) {
	if name != "" || stockStr != "" {
		quantity, err := strconv.Atoi(stockStr)
		if err != nil {
			fmt.Fprintln(inv.View, "Invalid stock!")
			return
		}

		inv.Items.AddItem(name, quantity)
		inv.RefereshView()
	}
}

func (inv *inventoryView) deleteHandler(indexStr string) {
	if indexStr != "" {
		val, err := strconv.Atoi(indexStr)
		if err != nil {
			fmt.Fprintln(inv.View, "Invalid index item")
			return
		}
		inv.Items.DeleteItem(val-1)
		inv.RefereshView()
	}
}
