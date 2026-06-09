package ui

import "github.com/rivo/tview"

type Layout struct {
	Layout *tview.Flex
}

func newLayout() *Layout{
	return &Layout{
		Layout: tview.NewFlex(),
	}
}

func BuildLayout(inv *tview.TextView, form *tview.Form) *Layout{
	layout :=newLayout()

	layout.Layout.
	AddItem(inv,0,1,false).
	AddItem(form,0,1,true)

	return layout

}