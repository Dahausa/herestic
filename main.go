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
	renderHeading()
	c := getConfiguration()
	mainMenu(c)
}

func renderHeading() {
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("HE", pterm.FgCyan.ToStyle()),
		putils.LettersFromStringWithStyle("restic", pterm.FgLightMagenta.ToStyle())).
		Render()
	pterm.Println("Herestic - The magic wand to summon the backup god -")
	pterm.Info.Printfln("Starting by casting some magic spells...")
}

func setResticBinary(c *Configuration) {
	if checkIfResticWorks(c.BinPath) != nil {
		pterm.Warning.Println("Restic binary was not found in " + c.BinPath)
		pterm.Warning.Println("Do you want to set the path? Otherwise program will be terminated")
		yes, _ := pterm.DefaultInteractiveConfirm.Show()
		if !yes {
			pterm.Warning.Printfln("Sorry, but it is time to leave")
			os.Exit(1)
		}
		pterm.Warning.Println("Please specify correct path to restic binary:")
		binPath, err := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show()
		if err != nil {
			panic("Something was wrong with the binary path")
		}
		c.BinPath = binPath
		setResticBinary(c)
	}
}

func setRepository(c *Configuration) {
	restic_repository := os.Getenv(RESTIC_REPOSITORY_ENV_VARIABLE)
	if restic_repository == "" {
		pterm.Warning.Println("But what's going on?!")
		pterm.Warning.Println("RESTIC_REPOSITORY_ENVIRONMENT is not set!")
		pterm.Warning.Println("So we shall not pass and have to configure the repository first...")
	}
	pterm.Info.Println("Enter the Restic Repository URL")
	restic_repository, err := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show()
	if err != nil {
		pterm.Error.Println("Something was wrong with the repository!")
	}
	pterm.Info.Println("Restic Repository set to: " + restic_repository)
	c.Repository = restic_repository
}

func getConfiguration() *Configuration {

	everythingWasAlreadyConfigured := true

	c, err := loadConf()

	if err != nil {
		pterm.Warning.Printfln("Missed my spellbook. A new one needs to be configured!")
		everythingWasAlreadyConfigured = false
	}

	setResticBinary(c)
	setRepository(c)

	if !everythingWasAlreadyConfigured {
		writeConf(c)
	}
	return c
}

func mainMenu(c *Configuration) {
	result, _ := pterm.DefaultInteractiveSelect.WithDefaultText("Select the next spell to cast:").
		WithOptions([]string{"Perform a Backup", "Let me see all Snapshots", "Configure Settings", "Quit"}).Show()

	switch result {
	case "Perform a Backup":
		performBackup(c.Repository)
		mainMenu(c)
	case "Let me see all Snapshots":
		listSnapshots(c.Repository)
		mainMenu(c)
	case "Configure Settings":
		configureSettings()
		mainMenu(c)
	case "Quit":
		os.Exit(0)
	}
}
