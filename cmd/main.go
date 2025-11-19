package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/kglaus/geodienste-cli/pkg/stac"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("geodienst-cli")
	myWindow.SetContent(widget.NewLabel("geodienste-cli"))

	myWindow.Show()
	myApp.Run()
	tidyUp()
}

func tidyUp() {
	fmt.Println("Exited")
	stac.TestStac()
}
