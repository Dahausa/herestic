package main

import (
	"github.com/pterm/pterm"
)

func performBackup(repository string) {
	pterm.Info.Printfln("Backing up yo! " + repository)
}
