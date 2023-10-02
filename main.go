package main

import (
	"os"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

const RESTIC_REPOSITORY_ENV_VARIABLE = "RESTIC_REPOSITORY"

type Configuration struct {
	Repository string
	BinPath    string
}

func main() {
	c := &Configuration{}
	renderHeading()
	getConfiguration(c)
	mainMenu(c)
}

func renderHeading() {
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("H", pterm.FgCyan.ToStyle()),
		putils.LettersFromStringWithStyle("erestic", pterm.FgLightMagenta.ToStyle())).
		Render()
	pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.BgRed)).Printfln("Summon the backup god!")
	pterm.Info.Printfln("Program started...")
}

func getConfiguration(c *Configuration) {

	everythingWasAlreadyConfigured := true

	if checkResticInPath() != nil {
		everythingWasAlreadyConfigured = false
		pterm.Warning.Println("Restic binary was not found in path")
		pterm.Warning.Println("Please specify path to restic binary")
		binPath, err := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show()
		if err != nil {
			panic("Something was wrong with the binary path")
		}
		c.BinPath = binPath
	}
	restic_repository := os.Getenv(RESTIC_REPOSITORY_ENV_VARIABLE)
	if restic_repository == "" {
		everythingWasAlreadyConfigured = false
		pterm.Warning.Println("RESTIC_REPOSITORY_ENVIRONMENT Variable not set!")
		pterm.Warning.Println("You have to configure the repository first...")
	}
	pterm.Info.Println("Please enter the Restic Repository URL")
	restic_repository, err := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show()
	if err != nil {
		pterm.Error.Println("Something was wrong with the repository!")
		getConfiguration(c)
	}
	pterm.Info.Println("Restic Repository set to: " + restic_repository)
	c.Repository = restic_repository

	if !everythingWasAlreadyConfigured {
		pterm.Info.Println("Will write a configuration file 'herestic.conf' with this new informations")
		writeConf(c)
	}
}

func mainMenu(c *Configuration) {
	result, _ := pterm.DefaultInteractiveSelect.
		WithOptions([]string{"Perform Backup", "List Snapshots", "Configure Settings", "Quit"}).Show()

	switch result {
	case "Perform Backup":
		performBackup(c.Repository)
		mainMenu(c)
	case "List Snapshots":
		listSnapshots(c.Repository)
		mainMenu(c)
	case "Configure Settings":
		configureSettings()
		mainMenu(c)
	case "Quit":
		os.Exit(0)
	}
}
