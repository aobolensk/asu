package asu

type UpdaterTask struct {
	APIVersion string `yaml:"api_version"`
	Name       string `yaml:"name"`
	Directory  string
	OnStart    struct {
		Script string   `yaml:"script"`
		Env    []string `yaml:"env"`
	} `yaml:"on_start"`
	Update struct {
		Interval int `yaml:"interval"`
		Before   struct {
			Script string   `yaml:"script"`
			Env    []string `yaml:"env"`
		} `yaml:"before"`
		On struct {
			Script string   `yaml:"script"`
			Env    []string `yaml:"env"`
		} `yaml:"on"`
		After struct {
			Script string   `yaml:"script"`
			Env    []string `yaml:"env"`
		} `yaml:"after"`
	} `yaml:"update"`
	OnStop struct {
		Script string   `yaml:"script"`
		Env    []string `yaml:"env"`
	} `yaml:"on_stop"`
}

type Config struct {
	Task UpdaterTask `yaml:"task"`
}
