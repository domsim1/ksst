package main

import (
	"fmt"
	"io"
	"net/url"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/domsim1/ksst/cmd/ksst-gui/data"
	"github.com/domsim1/ksst/pkg/encoder"
	"github.com/domsim1/ksst/pkg/model"
	"github.com/domsim1/ksst/pkg/util"
)

type View string

const (
	NoSaveView   = View("nosave")
	OverviewView = View("overview")
)

var (
	w            fyne.Window
	a            fyne.App
	saveMenuItem *fyne.MenuItem
	mainMenu     *fyne.MainMenu
	filePrefix   string
	fileSuffix   string
	saveData     *model.SaveData

	lastUri fyne.ListableURI
	views   map[View]*fyne.Container
)

func main() {
	a = app.New()
	a.SetIcon(data.KidIcon)
	w = a.NewWindow("ksst")
	mainMenu = makeMenu()
	w.SetMainMenu(mainMenu)
	w.SetMaster()

	initViews()
	w.SetContent(views[NoSaveView])

	w.Resize(fyne.NewSize(640, 480))
	w.ShowAndRun()
}

func initViews() {
	views = map[View]*fyne.Container{
		NoSaveView: makeNoSaveContainer(),
	}
}

func loadSaveViews() {
	views[OverviewView] = makeOverviewContainer()
}

func makeMenu() *fyne.MainMenu {
	loadItem := fyne.NewMenuItem("Load", openFileSelect)
	saveMenuItem = fyne.NewMenuItem("Save", openSaveSelect)
	saveMenuItem.Disabled = true
	file := fyne.NewMenu("File", loadItem, saveMenuItem)

	githubItem := fyne.NewMenuItem("Github", openGithub)
	source := fyne.NewMenu("Source", githubItem)

	main := fyne.NewMainMenu(file, source)
	return main
}

func makeNoSaveContainer() *fyne.Container {
	return container.NewPadded(
		container.NewVBox(
			container.NewHBox(
				widget.NewButton("Load Save", openFileSelect),
				layout.NewSpacer(),
			),
		),
	)
}

func makeOverviewContainer() *fyne.Container {
	nameEntry := widget.NewEntryWithData(binding.BindString(&saveData.Name))
	nameEntry.Validator = validation.NewRegexp("^[a-zA-Z]+$", "name must me a-z only")
	moneyEntry := widget.NewEntryWithData(binding.FloatToString(binding.BindFloat((*float64)(&saveData.Money))))
	form := widget.NewForm(
		widget.NewFormItem("Name", nameEntry),
		widget.NewFormItem("Money", moneyEntry),
	)
	form.Validate()
	return container.NewPadded(
		container.NewVScroll(
			container.NewVBox(
				form,
				widget.NewAccordion(
					widget.NewAccordionItem(
						fmt.Sprintf("Echo Bond: %.f", saveData.BondEcho),
						container.NewVBox(
							widget.NewLabel(fmt.Sprintf("Bond0: %.f", saveData.EchoBond0)),
							widget.NewLabel(fmt.Sprintf("Bond1: %.f", saveData.EchoBond1)),
							widget.NewLabel(fmt.Sprintf("Bond2: %.f", saveData.EchoBond2)),
							widget.NewLabel(fmt.Sprintf("Bond3: %.f", saveData.EchoBond3)),
							widget.NewLabel(fmt.Sprintf("Bond4: %.f", saveData.EchoBond4)),
							widget.NewLabel(fmt.Sprintf("Bond5: %.f", saveData.EchoBond5)),
							widget.NewLabel(fmt.Sprintf("Bond6: %.f", saveData.EchoBond6)),
							widget.NewLabel(fmt.Sprintf("Bond7: %.f", saveData.EchoBond7)),
							widget.NewLabel(fmt.Sprintf("Bond8: %.f", saveData.EchoBond8)),
							widget.NewLabel(fmt.Sprintf("Bond9: %.f", saveData.EchoBond9)),
						),
					),
					widget.NewAccordionItem(
						fmt.Sprintf("Thea Bond: %.f", saveData.BondThea),
						container.NewVBox(
							widget.NewLabel(fmt.Sprintf("Bond0: %.f", saveData.TheaBond0)),
							widget.NewLabel(fmt.Sprintf("Bond1: %.f", saveData.TheaBond1)),
							widget.NewLabel(fmt.Sprintf("Bond2: %.f", saveData.TheaBond2)),
							widget.NewLabel(fmt.Sprintf("Bond3: %.f", saveData.TheaBond3)),
							widget.NewLabel(fmt.Sprintf("Bond4: %.f", saveData.TheaBond4)),
							widget.NewLabel(fmt.Sprintf("Bond5: %.f", saveData.TheaBond5)),
							widget.NewLabel(fmt.Sprintf("Bond6: %.f", saveData.TheaBond6)),
							widget.NewLabel(fmt.Sprintf("Bond7: %.f", saveData.TheaBond7)),
							widget.NewLabel(fmt.Sprintf("Bond8: %.f", saveData.TheaBond8)),
							widget.NewLabel(fmt.Sprintf("Bond9: %.f", saveData.TheaBond9)),
						),
					),
					widget.NewAccordionItem(
						fmt.Sprintf("Dolus Bond: %.f", saveData.BondDolus),
						container.NewVBox(
							widget.NewLabel(fmt.Sprintf("Bond0: %.f", saveData.DolusBond0)),
							widget.NewLabel(fmt.Sprintf("Bond1: %.f", saveData.DolusBond1)),
							widget.NewLabel(fmt.Sprintf("Bond2: %.f", saveData.DolusBond2)),
							widget.NewLabel(fmt.Sprintf("Bond3: %.f", saveData.DolusBond3)),
							widget.NewLabel(fmt.Sprintf("Bond4: %.f", saveData.DolusBond4)),
							widget.NewLabel(fmt.Sprintf("Bond5: %.f", saveData.DolusBond5)),
							widget.NewLabel(fmt.Sprintf("Bond6: %.f", saveData.DolusBond6)),
							widget.NewLabel(fmt.Sprintf("Bond7: %.f", saveData.DolusBond7)),
							widget.NewLabel(fmt.Sprintf("Bond8: %.f", saveData.DolusBond8)),
							widget.NewLabel(fmt.Sprintf("Bond9: %.f", saveData.DolusBond9)),
						),
					),
					widget.NewAccordionItem(
						fmt.Sprintf("Bside Bond: %.f", saveData.BondBside),
						container.NewVBox(
							widget.NewLabel(fmt.Sprintf("Bond0: %.f", saveData.BsideBond0)),
							widget.NewLabel(fmt.Sprintf("Bond1: %.f", saveData.BsideBond1)),
							widget.NewLabel(fmt.Sprintf("Bond2: %.f", saveData.BsideBond2)),
							widget.NewLabel(fmt.Sprintf("Bond3: %.f", saveData.BsideBond3)),
							widget.NewLabel(fmt.Sprintf("Bond4: %.f", saveData.BsideBond4)),
							widget.NewLabel(fmt.Sprintf("Bond5: %.f", saveData.BsideBond5)),
							widget.NewLabel(fmt.Sprintf("Bond6: %.f", saveData.BsideBond6)),
							widget.NewLabel(fmt.Sprintf("Bond7: %.f", saveData.BsideBond7)),
							widget.NewLabel(fmt.Sprintf("Bond8: %.f", saveData.BsideBond8)),
							widget.NewLabel(fmt.Sprintf("Bond9: %.f", saveData.BsideBond9)),
						),
					),
				),
			),
		),
	)
}

func openGithub() {
	u, err := url.Parse("https://github.com/domsim1/ksst")
	util.Check(err)
	err = a.OpenURL(u)
	util.Check(err)
}

func openFileSelect() {
	fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		if reader == nil {
			return
		}
		loadSave(reader)
	}, w)
	fd.SetFilter(storage.NewExtensionFileFilter([]string{".sav"}))
	if lastUri != nil {
		fd.SetLocation(lastUri)
	}
	fd.Resize(w.Canvas().Size())
	fd.Show()
}

func openSaveSelect() {
	fd := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		if writer == nil {
			return
		}
		saveFile(writer)
	}, w)
	fd.SetFileName("knucklesandwich.sav")
	if lastUri != nil {
		fd.SetLocation(lastUri)
	}
	fd.Resize(w.Canvas().Size())
	fd.Show()
}

func loadSave(f fyne.URIReadCloser) {
	if f == nil {
		return
	}
	defer f.Close()
	disableSaveMenuItem()
	w.SetContent(views[NoSaveView])
	saveData = nil
	path, err := storage.ParseURI(strings.Replace(f.URI().String(), f.URI().Name(), "", 1))
	if err != nil {
		dialog.ShowError(err, w)
		return
	}
	fp, err := storage.ListerForURI(path)
	if err != nil {
		dialog.ShowError(err, w)
		return
	}
	lastUri = fp
	data, err := io.ReadAll(f)
	if err != nil {
		dialog.ShowError(err, w)
		return
	}
	dataStr, err := encoder.DecodeData(data)
	if err != nil {
		dialog.ShowError(err, w)
		return
	}
	saveData, filePrefix, err = model.ConvertStringData(dataStr)
	if err != nil {
		dialog.ShowError(err, w)
		return
	}
	enableSaveMenuItem()
	loadSaveViews()
	w.SetContent(views[OverviewView])
}

func saveFile(f fyne.URIWriteCloser) {
	if f == nil {
		return
	}
	defer f.Close()
	defer func() {
		saveData = nil
		filePrefix = ""
	}()
	disableSaveMenuItem()
	w.SetContent(views[NoSaveView])
	data, err := model.ConvertModelToStringData(saveData, filePrefix)
	if err != nil {
		dialog.ShowError(err, w)
		return
	}
	encodedData := encoder.EncodeData([]byte(data))
	_, err = f.Write([]byte(encodedData))
	if err != nil {
		dialog.ShowError(err, w)
		return
	}
	err = f.Close()
	if err != nil {
		dialog.ShowError(err, w)
		return
	}

}

func disableSaveMenuItem() {
	saveMenuItem.Disabled = true
	mainMenu.Refresh()
}

func enableSaveMenuItem() {
	saveMenuItem.Disabled = false
	mainMenu.Refresh()
}
