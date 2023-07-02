package tmpl

import (
	"io"
	"text/template"

	"github.com/pkg/errors"

	"leetcode/cmd/leetcode/graphql"
)

type zh struct {
	q    *graphql.QuestionData
	tmpl *template.Template
}

func NewZN(q *graphql.QuestionData) (l Parser) {
	return &zh{
		q:    q,
		tmpl: tmplZH,
	}
}

func (l *zh) Parse(wr io.Writer) (err error) {
	err = l.tmpl.Execute(wr, l.q)
	if err != nil {
		return errors.Wrap(err, "fail execute template")
	}

	return nil
}

func (l *zh) Filename() string {
	return "README.md"
}
