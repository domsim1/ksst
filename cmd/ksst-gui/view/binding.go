package view

import (
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/domsim1/ksst/pkg/model"
)

func BindSaveFloat(v *model.SaveFloat) binding.String {
	return binding.FloatToStringWithFormat(binding.BindFloat((*float64)(v)), "%.0f")
}

func EntrySF(v *model.SaveFloat) *widget.Entry {
	return widget.NewEntryWithData(BindSaveFloat(v))
}
