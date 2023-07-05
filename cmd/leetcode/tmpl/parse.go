package tmpl

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"leetcode/cmd/leetcode/graphql"
)

type Parser interface {
	Parse(wr io.Writer) (err error)
	Filename() string
}

type Parse struct {
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
	file, err := os.Create(filepath.Join(root, name))
	if err != nil {
		return errors.Wrap(err, "fail create file")
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return errors.Wrap(err, "fail write file")
	}

	return nil
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
