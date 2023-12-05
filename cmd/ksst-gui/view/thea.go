package view

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/domsim1/ksst/pkg/model"
)

func MakeTheaContainer(saveData *model.SaveData) *fyne.Container {
	left := widget.NewForm(
		widget.NewFormItem("Level", EntrySF(&saveData.PLevel9)),
		widget.NewFormItem("HP", EntrySF(&saveData.PHp9)),
		widget.NewFormItem("EP", EntrySF(&saveData.PMana9)),
		widget.NewFormItem("Attack", EntrySF(&saveData.PAtk9)),
		widget.NewFormItem("Speed", EntrySF(&saveData.PSpd9)),
		widget.NewFormItem("Charima", EntrySF(&saveData.PIq9)),
	)
	right := widget.NewForm(
		widget.NewFormItem("Exp", EntrySF(&saveData.PExp9)),
		widget.NewFormItem("Max HP", EntrySF(&saveData.PMaxhp9)),
		widget.NewFormItem("Max EP", EntrySF(&saveData.PMaxmana9)),
		widget.NewFormItem("Defence", EntrySF(&saveData.PDef9)),
		widget.NewFormItem("Luck", EntrySF(&saveData.PLck9)),
	)

	pocket := container.NewGridWithColumns(2, makeItemSelect(
		&saveData.PItem90,
		&saveData.PItem91,
		&saveData.PItem92,
		&saveData.PItem93,
		&saveData.PItem94,
		&saveData.PItem95,
		&saveData.PItem96,
		&saveData.PItem97,
	)...)

	return container.NewPadded(
		container.NewVScroll(
			container.NewVBox(
				widget.NewLabel(fmt.Sprintf("Title: %s", saveData.PTitle9)),
				container.NewGridWithColumns(2, left, right),
				widget.NewSeparator(),
				widget.NewLabel("Pocket"),
				pocket,
			),
		),
	)
}
