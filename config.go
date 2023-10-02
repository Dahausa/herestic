package main

import (
	"encoding/json"
	"os"
)

const configFile string = "herestic.json"

func writeConf(c *Configuration) error {
	f, err := os.Create(configFile)

	if err != nil {
		return err
	}

	j, err := json.Marshal(c)
	if err != nil {
		return err
	}
	_, err = f.WriteString(string(j))
	if err != nil {
		return err
	}

	return nil
}
