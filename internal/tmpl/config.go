package tmpl

type Config struct {
	Filename string `yaml:"filename"`       // 生成的文件名
	Template string `yaml:"tmpl,omitempty"` // 用于渲染的 go 模板文件路径

	Open bool   `yaml:"open"`          // 文件生成后，是否自动打开该文件
	App  string `yaml:"app,omitempty"` // Open 字段为 true 时才可用，使用指定软件打开文件，通常是 GoLand 或 VSCode 的安装位置，如果为空则使用系统默认软件打开
}
