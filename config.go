package main

import (
	"encoding/json"
	"os"

	"github.com/pterm/pterm"
)

const configFile string = "herestic.json"

func writeConf(c *Configuration) error {
	f, err := os.Create(configFile)
	if err != nil {
		return err
	}
	defer f.Close()

	j, err := json.Marshal(c)
	if err != nil {
		return err
	}
	_, err = f.WriteString(string(j))
	if err != nil {
		return err
	}

	pterm.Info.Println("New spells are written to " + configFile)

	return nil
}

func loadConf() (*Configuration, error) {
	c := &Configuration{}
	b, err := os.ReadFile(configFile)
	if err != nil {
		return c, err
	}
	err = json.Unmarshal(b, c)
	if err != nil {
		return c, err
	}
	return c, nil
}
