package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/kglaus/geodienste-cli/pkg/gui"
	"github.com/kglaus/geodienste-cli/pkg/stac"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("geodienst-cli")
	myWindow.SetContent(widget.NewLabel("geodienste-cli2"))
	myWindow.Resize(fyne.NewSize(1050, 400))

	var canvasObjects []fyne.CanvasObject

	contentBottom := container.New(layout.NewGridWrapLayout(fyne.Size{Width: 200, Height: 125}), canvasObjects...)
	contentBottomWrapper := container.NewVScroll(contentBottom)

	buttonGenerate := widget.NewButton("click me", func() {
		collections := stac.GetCollections()
		var collectionObjects []fyne.CanvasObject

		for _, element := range collections.Collections {
			collectionObjects = append(collectionObjects, gui.CollectionButton(element, contentBottom))
		}

		contentBottom.Objects = collectionObjects
		contentBottom.Refresh()
	})

	contentTop := container.New(layout.NewGridLayout(2), buttonGenerate)

	content := container.NewBorder(contentTop, nil, nil, nil, contentBottomWrapper)

	myWindow.SetContent(content)

	myWindow.Show()
	myApp.Run()
	tidyUp()
}

func tidyUp() {
	fmt.Println("Exited")
}
