package settings

type RepositoryReader interface {
	GetRepositoryPathFromConfig() string
	GetRepositoryPathFromEnv() string
}

type ResticBinReader interface {
	GetResticBinPathFromConfig() string
}

type SettingsReader interface {
	RepositoryReader
	ResticBinReader
}
