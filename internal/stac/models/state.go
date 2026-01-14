package models

import "fyne.io/fyne/v2/data/binding"

type State struct {
	CompleteList binding.UntypedList
	FilteredList binding.UntypedList
	Search       binding.String
}
