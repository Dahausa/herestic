package restic

type Snapshot struct {
	Id   string
	Name string
}

type ResticWrapper interface {
	CheckResticVersion() (string, error)
	InitRepoIfNotExists() error
	LoadListOfSnapshots() ([]Snapshot, error)
}
