package main

import (
	"github.com/dahausa/herestic/restic"
	"github.com/pterm/pterm"
)

func listSnapshots(w restic.ResticWrapper) {
	pterm.Info.Printfln("Snapshot list:")
	_, e := w.LoadListOfSnapshots()
	if e != nil {
		pterm.Error.Printfln("Error loading snapshots! %v", e)
	}

}
