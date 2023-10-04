package main

import (
	"github.com/dahausa/herestic/restic"
	"github.com/pterm/pterm"
)

func performBackup(w restic.ResticWrapper) {
	pterm.Info.Printfln("Backing up yo! ")
}
