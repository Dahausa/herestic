package main

import (
	"fmt"
	"os"
	"os/exec"

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

func checkResticInPath() error {
	_, err := exec.Command("restic version").Output()
	if err != nil {
		return err
	}
	return nil
}
