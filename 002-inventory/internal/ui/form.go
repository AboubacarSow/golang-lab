package ui

import "github.com/rivo/tview"

type Form struct {
	Form *tview.Form
}

func NewForm() *Form{
	form :=Form{
		Form: tview.NewForm(),
	}

	form.Form.
	SetBorder(true).
	SetTitle("Manage Inventory").
	SetTitleAlign(tview.AlignLeft)
	return &form
}

func BuildForm(
	inputName *tview.InputField,
	inputStock *tview.InputField,
	inputIndex *tview.InputField,
	addHandler func(name string,stockStr string),
	deleteHandler func(indexStr string),
	onExit func()) *Form{
	
	form:=NewForm()
	//Add Labels
	form.Form.
	AddFormItem(inputName).
	AddFormItem(inputStock).
	AddFormItem(inputIndex)

	//Add Button
	form.Form.
	AddButton("Add Item", func(){
		addHandler(inputName.GetText(),inputStock.GetText())
		inputName.SetText("")
		inputStock.SetText("")
	}).
	AddButton("Delete Item", func(){
		deleteHandler(inputIndex.GetText())
		inputIndex.SetText("")
	}).
	AddButton("Exit", func(){
		onExit()
	})
	return form
}