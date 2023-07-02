package tmpl

import (
	"io"
	"text/template"

	"github.com/pkg/errors"

	"leetcode/cmd/leetcode/graphql"
)

type en struct {
	q    *graphql.QuestionData
	tmpl *template.Template
}

func NewEN(q *graphql.QuestionData) (l Parser) {
	return &en{
		q:    q,
		tmpl: tmplEN,
	}
}

func (l *en) Parse(wr io.Writer) (err error) {
	err = l.tmpl.Execute(wr, l.q)
	if err != nil {
		return errors.Wrap(err, "fail execute template")
	}

	return nil
}

func (l *en) Filename() string {
	return "README_EN.md"
}
