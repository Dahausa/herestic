package main

import (
	"testing"

	"github.com/pterm/pterm"
)

func TestPterm(t *testing.T) {

	result, _ := pterm.DefaultInteractiveSelect.
		WithOptions([]string{"a", "b", "c", "d"}).
		Show()
	pterm.Info.Printfln("You answered: %s", result)
}
