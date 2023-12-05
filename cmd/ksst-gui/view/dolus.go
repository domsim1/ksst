package view

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/domsim1/ksst/pkg/model"
)

func MakeDolusContainer(saveData *model.SaveData) *fyne.Container {
	left := widget.NewForm(
		widget.NewFormItem("Level", EntrySF(&saveData.PLevel11)),
		widget.NewFormItem("HP", EntrySF(&saveData.PHp11)),
		widget.NewFormItem("EP", EntrySF(&saveData.PMana11)),
		widget.NewFormItem("Attack", EntrySF(&saveData.PAtk11)),
		widget.NewFormItem("Speed", EntrySF(&saveData.PSpd11)),
		widget.NewFormItem("Charima", EntrySF(&saveData.PIq11)),
	)
	right := widget.NewForm(
		widget.NewFormItem("Exp", EntrySF(&saveData.PExp11)),
		widget.NewFormItem("Max HP", EntrySF(&saveData.PMaxhp11)),
		widget.NewFormItem("Max EP", EntrySF(&saveData.PMaxmana11)),
		widget.NewFormItem("Defence", EntrySF(&saveData.PDef11)),
		widget.NewFormItem("Luck", EntrySF(&saveData.PLck11)),
	)

	pocket := container.NewGridWithColumns(2, makeItemSelect(
		&saveData.PItem110,
		&saveData.PItem111,
		&saveData.PItem112,
		&saveData.PItem113,
		&saveData.PItem114,
		&saveData.PItem115,
		&saveData.PItem116,
		&saveData.PItem117,
	)...)

	return container.NewPadded(
		container.NewVScroll(
			container.NewVBox(
				widget.NewLabel(fmt.Sprintf("Title: %s", saveData.PTitle11)),
				container.NewGridWithColumns(2, left, right),
				widget.NewSeparator(),
				widget.NewLabel("Pocket"),
				pocket,
			),
		),
	)
}
