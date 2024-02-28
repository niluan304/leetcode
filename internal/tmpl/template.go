package tmpl

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/pkg/errors"
	"github.com/skratchdot/open-golang/open"
)

type Template struct {
	config Config

	tmpl *template.Template
	data Data
}

func NewTemplate(config Config, data Data) *Template {
	return &Template{
		config: config,
		tmpl:   template.Must(template.ParseFiles(config.Template)),
		data:   data,
	}
}

func (t *Template) Execute(w io.Writer) (err error) {
	err = t.tmpl.Execute(w, t.data)
	if err != nil {
		return errors.Wrap(err, "fail execute template")
	}
	return nil
}

func (t *Template) Save(root string) (err error) {
	buf := new(bytes.Buffer)
	err = t.Execute(buf)
	if err != nil {
		return err
	}

	path := filepath.Join(root, t.config.Filename)
	defer t.open(path)

	data := bytes.ReplaceAll(buf.Bytes(), []byte("\r\n"), []byte("\n")) // git配置了：只使用 Unix 下的换行符
	err = os.WriteFile(path, data, 0o666)
	if err != nil {
		return errors.Wrap(err, "fail write file: "+path)
	}

	return nil
}

func (t *Template) open(path string) {
	if !t.config.Open {
		return
	}

	run, app := open.RunWith, t.config.App
	if app == "" {
		run = func(path string, appName string) error {
			return open.Run(path)
		}
	}

	absPath, _ := filepath.Abs(path)
	err := run(absPath, app)
	if err != nil {
		fmt.Printf("error: %v\nopen.RunWith(input:%s, appName:%s)\n", err, absPath, app)
	}
	time.Sleep(time.Second)
}
