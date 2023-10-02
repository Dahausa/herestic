package main

import (
	"github.com/pterm/pterm"
)

func listSnapshots(repository string) {
	pterm.Info.Printfln("Snapshot list:")
	b, e := loadListOfSnapshots(repository)
	if e != nil {
		pterm.Error.Printfln("Error loading snapshots! %v", e)
	}
	pterm.Info.Printfln(string(b))
}
