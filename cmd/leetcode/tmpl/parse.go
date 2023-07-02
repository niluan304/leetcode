package tmpl

import (
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

type Lang struct {
	Parser
}

func (l *Lang) Save(root string) (err error) {
	name := l.Filename()
	file, err := os.Create(filepath.Join(root, name))
	if err != nil {
		return errors.Wrap(err, "fail create file")
	}
	defer file.Close()

	err = l.Parse(file)
	if err != nil {
		return errors.Wrap(err, "fail parse"+name)
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
