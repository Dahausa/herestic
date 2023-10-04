package settings

type SettingsWriter interface {
	WriteSettings() error
}
