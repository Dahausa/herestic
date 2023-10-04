package main

import (
	"github.com/dahausa/herestic/restic"
	"github.com/dahausa/herestic/settings"
	"github.com/pterm/pterm"
)

/*
Create a restic wrapper which will be used to interact with restic
*/
func createResticWrapper() restic.ResticWrapper {
	var reader settings.SettingsReader = &settings.JsonSettings{}

	resticPath := reader.GetResticBinPathFromConfig()
	var wrapper restic.DefaultResticWrapper = restic.DefaultResticWrapper{ResticBinaryPath: resticPath}

	checkAndSetResticBinary(&wrapper)

	repository_env := reader.GetRepositoryPathFromEnv()
	if repository_env != "" {
		wrapper.ResticRepo = repository_env
		if wrapper.InitRepoIfNotExists() == nil {
			pterm.Info.Printfln("The repository is worth using it")
			return &wrapper
		}
	}

	pterm.Info.Printfln("I feel no presence of 'RESTIC_REPOSITORY' in the environment")
	pterm.Info.Printfln("I might then try the one from the file of config")

	wrapper.ResticRepo = reader.GetRepositoryPathFromConfig()
	checkAndSetRepo(&wrapper)

	return &wrapper
}

func checkAndSetResticBinary(wrapper *restic.DefaultResticWrapper) {
	resticVersion, err := wrapper.CheckResticVersion()
	if err == nil {
		pterm.Info.Printfln("Restic version '%s'", resticVersion)
		return
	}

	pterm.Error.Printfln("Behold! The pass to restic '%' is wrong", wrapper.ResticBinaryPath)
	pterm.Warning.Println("Do you want to set the path? Otherwise you will be terminated")
	yes, _ := pterm.DefaultInteractiveConfirm.Show()
	if !yes {
		pterm.Warning.Printfln("Sorry, but it is time to leave")
		panic("Panic on the titanic!")
	}
	pterm.Warning.Println("Please specify correct path to restic binary:")
	binPath, err := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show()
	if err != nil {
		panic("Something was wrong with the binary path")
	}
	wrapper.ResticBinaryPath = binPath
	checkAndSetResticBinary(wrapper)
}

func checkAndSetRepo(wrapper *restic.DefaultResticWrapper) {
	err := wrapper.InitRepoIfNotExists()
	if err == nil {
		pterm.Info.Printfln("You shall use the repo %s", wrapper.ResticRepo)
	}

	pterm.Warning.Printfln("The path of the repo seems not correct")
	pterm.Info.Println("So you shall provide a new path or die in vain")
	yes, _ := pterm.DefaultInteractiveConfirm.Show()
	if !yes {
		pterm.Warning.Printfln("Sorry, but it is time to leave")
		panic("Panic on the titanic!")
	}
	pterm.Warning.Println("Please specify the path to the repository")
	repo, err := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show()
	if err != nil {
		panic("Something was wrong with the binary path")
	}
	wrapper.ResticRepo = repo
	checkAndSetRepo(wrapper)
}
