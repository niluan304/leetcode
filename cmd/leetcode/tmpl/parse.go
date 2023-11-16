package tmpl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/pkg/errors"
	"github.com/skratchdot/open-golang/open"

	"leetcode/cmd/leetcode/graphql"
)

type Parser interface {
	Parse(w io.Writer) (err error)
	Filename() string
}

type tmpl struct {
	filename string
	data     any
	tmpl     *template.Template
}

func NewParser(filename string, t *template.Template, data any) Parser {
	return &tmpl{
		filename: filename,
		data:     data,
		tmpl:     t,
	}
}

func (t *tmpl) Parse(w io.Writer) (err error) {
	err = t.tmpl.Execute(w, t.data)
	if err != nil {
		return errors.Wrap(err, "fail execute template")
	}

	return nil
}

func (t *tmpl) Filename() string {
	return t.filename
}

type Parse struct {
	Open bool
	App  string
	Parser
}

func (p *Parse) Save(root string) (err error) {
	var buf = new(bytes.Buffer)
	err = p.Parse(buf)
	if err != nil {
		return errors.Wrap(err, "fail parse")
	}

	data := bytes.ReplaceAll(buf.Bytes(), []byte("\r\n"), []byte("\n"))

	name := p.Filename()
	path := filepath.Join(root, name)
	file, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "fail create file")
	}
	defer file.Close()

	p.open(path)

	_, err = file.Write(data)
	if err != nil {
		return errors.Wrap(err, "fail write file")
	}

	return nil
}

func (p *Parse) open(path string) {
	if !p.Open {
		return
	}

	run := open.RunWith
	if p.App == "" {
		run = func(path string, appName string) error {
			return open.Run(path)
		}
	}

	absPath, _ := filepath.Abs(path)
	err := run(absPath, p.App)
	if err != nil {
		fmt.Println("error: open.RunWith(input, appName)", absPath, p.App)
	}
	time.Sleep(time.Second)
}

func LoadQuestion(responseFile string) (q *graphql.QuestionData, err error) {
	file, err := os.Open(responseFile)
	if err != nil {
		return q, errors.Wrap(err, "fail read file")
	}
	defer file.Close()

	var v graphql.QuestionDataRes
	err = json.NewDecoder(file).Decode(&v)
	if err != nil {
		return q, errors.Wrap(err, "fail decode json to QuestionDataRes")
	}

	return v.Question.ReUnmarshal()
}
