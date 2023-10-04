package main

import (
	"os"

	"github.com/dahausa/herestic/restic"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	renderHeading()
	restic := createResticWrapper()
	mainMenu(restic)
}

func renderHeading() {
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("HE", pterm.FgCyan.ToStyle()),
		putils.LettersFromStringWithStyle("restic", pterm.FgLightMagenta.ToStyle())).
		Render()
	pterm.Println("Herestic - The magic wand to summon the backup god -")
	pterm.Info.Printfln("Starting by casting some magic spells...")
}

func mainMenu(w restic.ResticWrapper) {
	result, _ := pterm.DefaultInteractiveSelect.WithDefaultText("Select the next spell to cast:").
		WithOptions([]string{"Perform a Backup", "Let me see all Snapshots", "Configure Settings", "Quit"}).Show()

	switch result {
	case "Perform a Backup":
		performBackup(w)
		mainMenu(w)
	case "Let me see all Snapshots":
		listSnapshots(w)
		mainMenu(w)
	case "Quit":
		os.Exit(0)
	}
}
