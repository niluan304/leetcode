package tmpl

import (
	"io"
	"text/template"

	"github.com/pkg/errors"

	"leetcode/cmd/leetcode/graphql"
)

type zh struct {
	data *graphql.QuestionData
	tmpl *template.Template
}

func NewParserZN(q *graphql.QuestionData) (p Parser) {
	return &zh{
		data: q,
		tmpl: tmplZH,
	}
}

func (p *zh) Parse(wr io.Writer) (err error) {
	err = p.tmpl.Execute(wr, p.data)
	if err != nil {
		return errors.Wrap(err, "fail execute template")
	}

	return nil
}

func (p *zh) Filename() string {
	return "README.md"
}
