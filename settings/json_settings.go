package settings

import (
	"encoding/json"
	"os"
)

type JsonConfigurationFileStructure struct {
	repository string
	resticPath string
}

type JsonSettings struct {
	ConfigFilePath string
	configFile     JsonConfigurationFileStructure
}

func (r *JsonSettings) readConfigFile() {
	if r.configFile.repository != "" {
		return
	}

	c := &JsonConfigurationFileStructure{}
	b, err := os.ReadFile(configFile)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, c)
	if err != nil {
		panic(err)
	}
}

func (r *JsonSettings) GetRepositoryPathFromConfig() string {
	r.readConfigFile()
	return r.configFile.repository
}

func (r *JsonSettings) GetRepositoryPathFromEnv() string {
	return os.Getenv("RESTIC_REPOSITORY")
}

func (r *JsonSettings) GetResticBinPathFromConfig() string {
	r.readConfigFile()
	return r.configFile.resticPath
}

const configFile string = "herestic.json"

func (r *JsonSettings) WriteSettings() error {
	f, err := os.Create(configFile)
	if err != nil {
		return err
	}
	defer f.Close()

	j, err := json.Marshal(r.configFile)
	if err != nil {
		return err
	}
	_, err = f.WriteString(string(j))
	if err != nil {
		return err
	}

	return nil
}
