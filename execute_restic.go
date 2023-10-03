package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/pterm/pterm"
	"github.com/tidwall/gjson"
)

type Snapshot struct {
	Id   string
	Date string
}

func (s *Snapshot) fromJson(b []byte) {
	stringJson := string(b)
	value := gjson.Get(stringJson, "id")
	s.Id = value.String()
}

func invokeRestic(repo string, args string) ([]byte, error) {
	os.Setenv("RESTIC_REPOSITORY", repo)

	commandString := fmt.Sprintf("restic %s --json", args)

	stdout, err := exec.Command(commandString).Output()
	if err != nil {
		return []byte{}, err
	}
	return stdout, nil
}

func loadListOfSnapshots(repo string) ([]byte, error) {
	return invokeRestic(repo, "snapshots")
}

func checkIfResticWorks(path string) error {
	o, err := exec.Command(path, "version").CombinedOutput()
	if err != nil {
		pterm.Error.Println(err)
		return err
	}
	pterm.Info.Printfln("Restic comes to life: '%v'", strings.TrimSpace(string(o)))
	return nil
}
