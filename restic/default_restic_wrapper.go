package restic

import (
	"fmt"
	"os/exec"
	//"strings"
)

type DefaultResticWrapper struct {
	ResticBinaryPath string
	ResticRepo       string
}

func (w *DefaultResticWrapper) invokeRestic(args string) ([]byte, error) {

	commandString := fmt.Sprintf("%s %s -r %s --json", w.ResticBinaryPath, args, w.ResticRepo)

	stdout, err := exec.Command(commandString).Output()
	if err != nil {
		return []byte{}, err
	}
	return stdout, nil
}

func (w *DefaultResticWrapper) CheckResticVersion() (string, error) {
	o, err := exec.Command(w.ResticBinaryPath, "version").CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(o), nil
}

func (w *DefaultResticWrapper) InitRepoIfNotExists() error {
	_, err := w.invokeRestic("init")
	return err
}

func (w *DefaultResticWrapper) LoadListOfSnapshots() ([]Snapshot, error) {
	_, err := w.invokeRestic("snapshots")
	s := []Snapshot{}
	//TODO: Convert bytes to snapshot
	return s, err
}
