package main

import (
	"io"
	"net/url"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/domsim1/ksst/cmd/ksst-gui/data"
	"github.com/domsim1/ksst/cmd/ksst-gui/view"
	"github.com/domsim1/ksst/pkg/encoder"
	"github.com/domsim1/ksst/pkg/model"
	"github.com/domsim1/ksst/pkg/util"
)

type View string

var (
	w            fyne.Window
	a            fyne.App
	saveMenuItem *fyne.MenuItem
	mainMenu     *fyne.MainMenu
	filePrefix   string
	fileSuffix   string
	saveData     *model.SaveData

	lastUri    fyne.ListableURI
	noSaveView *fyne.Container
)

func main() {
	a = app.New()
	a.SetIcon(data.KidIcon)
	w = a.NewWindow("ksst - v0.0.1")
	mainMenu = makeMenu()
	w.SetMainMenu(mainMenu)
	w.SetMaster()

	noSaveView = makeNoSaveContainer()
	w.SetContent(noSaveView)

	w.Resize(fyne.NewSize(640, 480))
	w.ShowAndRun()
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

func makeTabs() *container.AppTabs {
	tabs := container.NewAppTabs(
		container.NewTabItem("New Kid", view.MakePlayerContainer(saveData)),
		container.NewTabItem("Thea", view.MakeTheaContainer(saveData)),
		container.NewTabItem("Echo", view.MakeEchoContainer(saveData)),
		container.NewTabItem("Dolus", view.MakeDolusContainer(saveData)),
	)
	tabs.SetTabLocation(container.TabLocationTop)
	return tabs
}

func makeNoSaveContainer() *fyne.Container {
	return container.NewPadded(
		container.NewVBox(
			widget.NewLabel("How to use:"),
			widget.NewLabel("1. Backup save data"),
			widget.NewLabel("2. Make sure game is closed"),
			widget.NewLabel("3. Click 'Load Save' and find the save file"),
			widget.NewLabel("4. Edit stuff"),
			widget.NewLabel("5. Click file->save and overwrite the save file"),
			widget.NewLabel("6. Remove savefile.ini"),
			widget.NewLabel("7. Run the game and go HAM"),
			container.NewHBox(
				widget.NewButton("Load Save", openFileSelect),
				layout.NewSpacer(),
			),
			widget.NewLabel("Note: I am not responsible if you lose any save data.\nAlway backup saves and use at your own risk!"),
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
	w.SetContent(noSaveView)
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
	tabs := makeTabs()
	w.SetContent(tabs)
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
	w.SetContent(noSaveView)
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
