package ui

import (
	"fmt"
	"inventory/internal/inventory"

	"github.com/alexeyco/simpletable"
	"github.com/rivo/tview"
)

type inventoryView struct {
	View  *tview.TextView
	Items *inventory.Inventory
}

func BuildInventoryView(items *inventory.Inventory) *inventoryView {
	inv := inventoryView{
		View:  tview.NewTextView(),
		Items: items,
	}
	inv.View.SetDynamicColors(true).
		SetWordWrap(true)

	inv.View.SetBorder(true).
		SetTitle("Inventory Items")

	return &inv
}

func (inv *inventoryView) RefereshView() {
	inv.View.Clear()
	if len(*inv.Items) == 0 {
		fmt.Fprintln(inv.View, "No items on inventory")
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
	for i, item := range *inv.Items {

		row := []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%d", i+1)},
			{Align: simpletable.AlignLeft, Text: item.Name},
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%d", item.Stock)},
		}

		table.Body.Cells = append(table.Body.Cells, row)
	}

	table.SetStyle(simpletable.StyleUnicode)
	inv.View.SetText(table.String())

}
