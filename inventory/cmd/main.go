package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/alexeyco/simpletable"
	"github.com/rivo/tview"
)

type item struct {
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

// constant params
var (
	inventory     = []item{}
	inventoryPath = "data/inventory.json"
)

//load inventory

func loadInventory() error {

	data, err := os.ReadFile(inventoryPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(data, &inventory)

}

// save
func saveInventory() error {
	data, err := json.Marshal(&inventory)
	if err != nil {
		return err
	}
	return os.WriteFile(inventoryPath, data, 0644)
}

func deleteItem(index int) error {
	if index >= 0 || len(inventory) > index {
		inventory = append(inventory[:index], inventory[index+1:]...)
		saveInventory()
		return nil
	}
	return errors.New("Invalid index")

}
func addItem(i item) {
	inventory = append(inventory, i)
	saveInventory()
}

func main() {
	app := tview.NewApplication()
	loadInventory()

	inventoryList := tview.NewTextView().
		SetDynamicColors(true).
		SetWordWrap(true)

	inventoryList.SetBorder(true).
		SetTitle("Inventory Items")

	refresh := func() {
		inventoryList.Clear()
		if len(inventory) == 0 {
			fmt.Fprintln(inventoryList, "No items on inventory")
			return
		}
		table := simpletable.New()
		table.Header = &simpletable.Header{
			Cells: []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: "#"},
				{Align: simpletable.AlignCenter, Text: "Name"},
				{Align: simpletable.AlignCenter, Text: "Stock"},
			},
		}
		for i, item := range inventory {

			row := []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%d", i+1)},
				{Align: simpletable.AlignLeft, Text: item.Name},
				{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%d", item.Stock)},
			}

			table.Body.Cells = append(table.Body.Cells, row)
		}

		table.SetStyle(simpletable.StyleUnicode)
		inventoryList.SetText(table.String())
	}

	itemNameInput := tview.NewInputField().SetLabel("Item Name")
	itemStockInput := tview.NewInputField().SetLabel("Item Stock")
	itemIndexForDelete := tview.NewInputField().SetLabel("Index item")

	form := tview.NewForm().
		AddFormItem(itemNameInput).
		AddFormItem(itemStockInput).
		AddFormItem(itemIndexForDelete).
		AddButton("Add Item", func() {
			name := itemNameInput.GetText()
			stock := itemStockInput.GetText()
			if name != "" || stock != "" {
				quantity, err := strconv.Atoi(stock)
				if err != nil {
					fmt.Fprintln(inventoryList, "Invalid stock!")
					return
				}
				newItem := item{
					Name:  name,
					Stock: quantity,
				}

				addItem(newItem)
				refresh()
				itemNameInput.SetText("")
				itemStockInput.SetText("")
			}
		}).
		AddButton("Delete an item", func() {
			index := itemIndexForDelete.GetText()
			if index != "" {
				val, err := strconv.Atoi(index)
				if err != nil {
					fmt.Fprintln(inventoryList, "Invalid index item")
					return
				}
				deleteItem(val)
				refresh()
				itemIndexForDelete.SetText("")

			}
		}).
		AddButton("Exit", func() {
			app.Stop()
		})

	form.SetBorder(true).SetTitle("Manage Inventory").SetTitleAlign(tview.AlignLeft)

	flex := tview.NewFlex().AddItem(inventoryList, 0, 1, false).
		AddItem(form, 0, 1, true)

	refresh()

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
