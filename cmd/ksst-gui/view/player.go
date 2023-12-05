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

	pocket := widget.NewForm(
		widget.NewFormItem("#1", EntrySF(&saveData.PItem00)),
	)

	return container.NewPadded(
		container.NewVScroll(
			container.NewVBox(
				widget.NewLabel(fmt.Sprintf("Title: %s", saveData.PTitle0)),
				container.NewGridWithColumns(2, left, right),
				widget.NewSeparator(),
				widget.NewLabel("Pocket"),
				pocket,
			),
		),
	)
}

//				widget.NewAccordion(
//					widget.NewAccordionItem(
//						fmt.Sprintf("Echo Bond: %.0f", saveData.BondEcho),
//						container.NewVBox(
//							widget.NewLabel(fmt.Sprintf("Bond0: %.f", saveData.EchoBond0)),
//							widget.NewLabel(fmt.Sprintf("Bond1: %.f", saveData.EchoBond1)),
//							widget.NewLabel(fmt.Sprintf("Bond2: %.f", saveData.EchoBond2)),
//							widget.NewLabel(fmt.Sprintf("Bond3: %.f", saveData.EchoBond3)),
//							widget.NewLabel(fmt.Sprintf("Bond4: %.f", saveData.EchoBond4)),
//							widget.NewLabel(fmt.Sprintf("Bond5: %.f", saveData.EchoBond5)),
//							widget.NewLabel(fmt.Sprintf("Bond6: %.f", saveData.EchoBond6)),
//							widget.NewLabel(fmt.Sprintf("Bond7: %.f", saveData.EchoBond7)),
//							widget.NewLabel(fmt.Sprintf("Bond8: %.f", saveData.EchoBond8)),
//							widget.NewLabel(fmt.Sprintf("Bond9: %.f", saveData.EchoBond9)),
//						),
//					),
//					widget.NewAccordionItem(
//						fmt.Sprintf("Thea Bond: %.f", saveData.BondThea),
//						container.NewVBox(
//							widget.NewLabel(fmt.Sprintf("Bond0: %.f", saveData.TheaBond0)),
//							widget.NewLabel(fmt.Sprintf("Bond1: %.f", saveData.TheaBond1)),
//							widget.NewLabel(fmt.Sprintf("Bond2: %.f", saveData.TheaBond2)),
//							widget.NewLabel(fmt.Sprintf("Bond3: %.f", saveData.TheaBond3)),
//							widget.NewLabel(fmt.Sprintf("Bond4: %.f", saveData.TheaBond4)),
//							widget.NewLabel(fmt.Sprintf("Bond5: %.f", saveData.TheaBond5)),
//							widget.NewLabel(fmt.Sprintf("Bond6: %.f", saveData.TheaBond6)),
//							widget.NewLabel(fmt.Sprintf("Bond7: %.f", saveData.TheaBond7)),
//							widget.NewLabel(fmt.Sprintf("Bond8: %.f", saveData.TheaBond8)),
//							widget.NewLabel(fmt.Sprintf("Bond9: %.f", saveData.TheaBond9)),
//						),
//					),
//					widget.NewAccordionItem(
//						fmt.Sprintf("Dolus Bond: %.f", saveData.BondDolus),
//						container.NewVBox(
//							widget.NewLabel(fmt.Sprintf("Bond0: %.f", saveData.DolusBond0)),
//							widget.NewLabel(fmt.Sprintf("Bond1: %.f", saveData.DolusBond1)),
//							widget.NewLabel(fmt.Sprintf("Bond2: %.f", saveData.DolusBond2)),
//							widget.NewLabel(fmt.Sprintf("Bond3: %.f", saveData.DolusBond3)),
//							widget.NewLabel(fmt.Sprintf("Bond4: %.f", saveData.DolusBond4)),
//							widget.NewLabel(fmt.Sprintf("Bond5: %.f", saveData.DolusBond5)),
//							widget.NewLabel(fmt.Sprintf("Bond6: %.f", saveData.DolusBond6)),
//							widget.NewLabel(fmt.Sprintf("Bond7: %.f", saveData.DolusBond7)),
//							widget.NewLabel(fmt.Sprintf("Bond8: %.f", saveData.DolusBond8)),
//							widget.NewLabel(fmt.Sprintf("Bond9: %.f", saveData.DolusBond9)),
//						),
//					),
//					widget.NewAccordionItem(
//						fmt.Sprintf("Bside Bond: %.f", saveData.BondBside),
//						container.NewVBox(
//							widget.NewLabel(fmt.Sprintf("Bond0: %.f", saveData.BsideBond0)),
//							widget.NewLabel(fmt.Sprintf("Bond1: %.f", saveData.BsideBond1)),
//							widget.NewLabel(fmt.Sprintf("Bond2: %.f", saveData.BsideBond2)),
//							widget.NewLabel(fmt.Sprintf("Bond3: %.f", saveData.BsideBond3)),
//							widget.NewLabel(fmt.Sprintf("Bond4: %.f", saveData.BsideBond4)),
//							widget.NewLabel(fmt.Sprintf("Bond5: %.f", saveData.BsideBond5)),
//							widget.NewLabel(fmt.Sprintf("Bond6: %.f", saveData.BsideBond6)),
//							widget.NewLabel(fmt.Sprintf("Bond7: %.f", saveData.BsideBond7)),
//							widget.NewLabel(fmt.Sprintf("Bond8: %.f", saveData.BsideBond8)),
//							widget.NewLabel(fmt.Sprintf("Bond9: %.f", saveData.BsideBond9)),
//						),
//					),
//				),
