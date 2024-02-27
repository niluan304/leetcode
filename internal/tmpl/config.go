package tmpl

type Config struct {
	Name     string `yaml:"name"`
	Filename string `yaml:"filename"`
	Open     bool   `yaml:"open"`
	App      string `yaml:"app,omitempty"`
	Template string `yaml:"tmpl,omitempty"`
}
