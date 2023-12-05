package view

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/widget"
	"github.com/domsim1/ksst/pkg/model"
)

func MakePlayerContainer(saveData *model.SaveData) *fyne.Container {
	name := widget.NewEntryWithData(binding.BindString(&saveData.Name))
	name.Validator = validation.NewRegexp("^[a-zA-Z]+$", "name must me a-z only")
	money := widget.NewEntryWithData(BindSaveFloat(&saveData.Money))
	nameMoney := widget.NewForm(
		widget.NewFormItem("Name", name),
	)
	nameMoney.Validate()

	left := widget.NewForm(
		widget.NewFormItem("Name", name),
		widget.NewFormItem("Level", EntrySF(&saveData.PLevel0)),
		widget.NewFormItem("HP", EntrySF(&saveData.PHp0)),
		widget.NewFormItem("EP", EntrySF(&saveData.PMana0)),
		widget.NewFormItem("Attack", EntrySF(&saveData.PAtk0)),
		widget.NewFormItem("Speed", EntrySF(&saveData.PSpd0)),
		widget.NewFormItem("Charima", EntrySF(&saveData.PIq0)),
	)
	right := widget.NewForm(
		widget.NewFormItem("Money", money),
		widget.NewFormItem("Exp", EntrySF(&saveData.PExp0)),
		widget.NewFormItem("Max HP", EntrySF(&saveData.PMaxhp0)),
		widget.NewFormItem("Max EP", EntrySF(&saveData.PMaxmana0)),
		widget.NewFormItem("Defence", EntrySF(&saveData.PDef0)),
		widget.NewFormItem("Luck", EntrySF(&saveData.PLck0)),
	)

	pocket := container.NewGridWithColumns(2, makeItemSelect(
		&saveData.PItem00,
		&saveData.PItem01,
		&saveData.PItem02,
		&saveData.PItem03,
		&saveData.PItem04,
		&saveData.PItem05,
		&saveData.PItem06,
		&saveData.PItem07,
	)...)

	bonds := makeBondsAccordion(saveData)

	return container.NewPadded(
		container.NewVScroll(
			container.NewVBox(
				widget.NewLabel(fmt.Sprintf("Title: %s", saveData.PTitle0)),
				container.NewGridWithColumns(2, left, right),
				widget.NewSeparator(),
				widget.NewLabel("Pocket"),
				pocket,
				widget.NewSeparator(),
				widget.NewLabel("Bonds"),
				bonds,
			),
		),
	)
}

func makeItemSelect(data ...*model.SaveFloat) (selects []fyne.CanvasObject) {
	for i := range data {
		d := data[i]
		s := widget.NewSelect(model.ItemNameList, func(item string) {
			*d = model.ItemNameMap[item]
		})
		s.SetSelected(model.ItemIDMap[*data[i]])
		selects = append(selects, s)
	}
	return selects
}

func makeBondsAccordion(saveData *model.SaveData) *widget.Accordion {

	return widget.NewAccordion(
		widget.NewAccordionItem(
			fmt.Sprintf("Echo Bond: %.0f", saveData.BondEcho),
			container.NewGridWithColumns(2,
				widget.NewForm(
					widget.NewFormItem("Bond0", EntrySF(&saveData.EchoBond0)),
					widget.NewFormItem("Bond1", EntrySF(&saveData.EchoBond1)),
					widget.NewFormItem("Bond2", EntrySF(&saveData.EchoBond2)),
					widget.NewFormItem("Bond3", EntrySF(&saveData.EchoBond3)),
					widget.NewFormItem("Bond4", EntrySF(&saveData.EchoBond4)),
				),
				widget.NewForm(
					widget.NewFormItem("Bond5", EntrySF(&saveData.EchoBond5)),
					widget.NewFormItem("Bond6", EntrySF(&saveData.EchoBond6)),
					widget.NewFormItem("Bond7", EntrySF(&saveData.EchoBond7)),
					widget.NewFormItem("Bond8", EntrySF(&saveData.EchoBond8)),
					widget.NewFormItem("Bond9", EntrySF(&saveData.EchoBond9)),
				),
			),
		),
		widget.NewAccordionItem(
			fmt.Sprintf("Thea Bond: %.f", saveData.BondThea),
			container.NewGridWithColumns(2,
				widget.NewForm(
					widget.NewFormItem("Bond0", EntrySF(&saveData.TheaBond0)),
					widget.NewFormItem("Bond1", EntrySF(&saveData.TheaBond1)),
					widget.NewFormItem("Bond2", EntrySF(&saveData.TheaBond2)),
					widget.NewFormItem("Bond3", EntrySF(&saveData.TheaBond3)),
					widget.NewFormItem("Bond4", EntrySF(&saveData.TheaBond4)),
				),
				widget.NewForm(
					widget.NewFormItem("Bond5", EntrySF(&saveData.TheaBond5)),
					widget.NewFormItem("Bond6", EntrySF(&saveData.TheaBond6)),
					widget.NewFormItem("Bond7", EntrySF(&saveData.TheaBond7)),
					widget.NewFormItem("Bond8", EntrySF(&saveData.TheaBond8)),
					widget.NewFormItem("Bond9", EntrySF(&saveData.TheaBond9)),
				),
			),
		),
		widget.NewAccordionItem(
			fmt.Sprintf("Dolus Bond: %.f", saveData.BondDolus),
			container.NewGridWithColumns(2,
				widget.NewForm(
					widget.NewFormItem("Bond0", EntrySF(&saveData.DolusBond0)),
					widget.NewFormItem("Bond1", EntrySF(&saveData.DolusBond1)),
					widget.NewFormItem("Bond2", EntrySF(&saveData.DolusBond2)),
					widget.NewFormItem("Bond3", EntrySF(&saveData.DolusBond3)),
					widget.NewFormItem("Bond4", EntrySF(&saveData.DolusBond4)),
				),
				widget.NewForm(

					widget.NewFormItem("Bond5", EntrySF(&saveData.DolusBond5)),
					widget.NewFormItem("Bond6", EntrySF(&saveData.DolusBond6)),
					widget.NewFormItem("Bond7", EntrySF(&saveData.DolusBond7)),
					widget.NewFormItem("Bond8", EntrySF(&saveData.DolusBond8)),
					widget.NewFormItem("Bond9", EntrySF(&saveData.DolusBond9)),
				),
			),
		),
		widget.NewAccordionItem(
			fmt.Sprintf("Bside Bond: %.f", saveData.BondBside),
			container.NewGridWithColumns(2,
				widget.NewForm(
					widget.NewFormItem("Bond0", EntrySF(&saveData.BsideBond0)),
					widget.NewFormItem("Bond1", EntrySF(&saveData.BsideBond1)),
					widget.NewFormItem("Bond2", EntrySF(&saveData.BsideBond2)),
					widget.NewFormItem("Bond3", EntrySF(&saveData.BsideBond3)),
					widget.NewFormItem("Bond4", EntrySF(&saveData.BsideBond4)),
				),
				widget.NewForm(
					widget.NewFormItem("Bond5", EntrySF(&saveData.BsideBond5)),
					widget.NewFormItem("Bond6", EntrySF(&saveData.BsideBond6)),
					widget.NewFormItem("Bond7", EntrySF(&saveData.BsideBond7)),
					widget.NewFormItem("Bond8", EntrySF(&saveData.BsideBond8)),
					widget.NewFormItem("Bond9", EntrySF(&saveData.BsideBond9)),
				),
			),
		),
	)
}
