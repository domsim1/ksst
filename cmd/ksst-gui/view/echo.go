package view

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/domsim1/ksst/pkg/model"
)

func MakeEchoContainer(saveData *model.SaveData) *fyne.Container {
	left := widget.NewForm(
		widget.NewFormItem("Level", EntrySF(&saveData.PLevel8)),
		widget.NewFormItem("HP", EntrySF(&saveData.PHp8)),
		widget.NewFormItem("EP", EntrySF(&saveData.PMana8)),
		widget.NewFormItem("Attack", EntrySF(&saveData.PAtk8)),
		widget.NewFormItem("Speed", EntrySF(&saveData.PSpd8)),
		widget.NewFormItem("Charima", EntrySF(&saveData.PIq8)),
	)
	right := widget.NewForm(
		widget.NewFormItem("Exp", EntrySF(&saveData.PExp8)),
		widget.NewFormItem("Max HP", EntrySF(&saveData.PMaxhp8)),
		widget.NewFormItem("Max EP", EntrySF(&saveData.PMaxmana8)),
		widget.NewFormItem("Defence", EntrySF(&saveData.PDef8)),
		widget.NewFormItem("Luck", EntrySF(&saveData.PLck8)),
	)

	pocket := container.NewGridWithColumns(2, makeItemSelect(
		&saveData.PItem80,
		&saveData.PItem81,
		&saveData.PItem82,
		&saveData.PItem83,
		&saveData.PItem84,
		&saveData.PItem85,
		&saveData.PItem86,
		&saveData.PItem87,
	)...)

	return container.NewPadded(
		container.NewVScroll(
			container.NewVBox(
				widget.NewLabel(fmt.Sprintf("Title: %s", saveData.PTitle8)),
				container.NewGridWithColumns(2, left, right),
				widget.NewSeparator(),
				widget.NewLabel("Pocket"),
				pocket,
			),
		),
	)
}
