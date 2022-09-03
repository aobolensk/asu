package asu

type UpdaterTask struct {
	APIVersion string `yaml:"api_version"`
	Name       string `yaml:"name"`
	OnStart    struct {
		Script string   `yaml:"script"`
		Env    []string `yaml:"env"`
	} `yaml:"on_start"`
	OnStop struct {
		Script string   `yaml:"script"`
		Env    []string `yaml:"env"`
	} `yaml:"on_stop"`
}

type Config struct {
	Task UpdaterTask `yaml:"task"`
}
