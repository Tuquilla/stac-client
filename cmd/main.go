package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/kglaus/geodienste-cli/internal/gui"
	"github.com/kglaus/geodienste-cli/internal/helpers"
	"github.com/kglaus/geodienste-cli/internal/stac/models"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("geodienst-cli")
	myWindow.SetContent(widget.NewLabel("geodienste-cli"))
	myWindow.Resize(fyne.NewSize(1050, 400))

	stateBindings := models.State{
		CompleteList: binding.NewUntypedList(),
		FilteredList: binding.NewUntypedList(),
		Search:       binding.NewString(),
	}

	var collections models.Collections

	selectOptions := helpers.Config("configs/config.json").BaseUrls

	// Mainframe
	contentBottomWrapper := gui.NewMainFrame(stateBindings)

	// Searchbar
	searchBar := gui.NewSearchBar(stateBindings, collections)

	// Select entry
	selectEntry := gui.NewSelectEntry(selectOptions)

	// Stac button
	buttonGenerate := gui.NewStacButton(stateBindings, collections, selectEntry)

	// Arrange widgets
	contentTop := container.New(layout.NewGridLayout(2), buttonGenerate, searchBar, selectEntry)
	content := container.NewBorder(contentTop, nil, nil, nil, contentBottomWrapper)

	myWindow.SetContent(content)

	myWindow.Show()
	myApp.Run()
	onClose()
}

func onClose() {
	fmt.Println("Exited")
}
