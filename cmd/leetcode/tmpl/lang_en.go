package tmpl

import (
	"io"
	"text/template"

	"github.com/pkg/errors"

	"leetcode/cmd/leetcode/graphql"
)

type en struct {
	data *graphql.QuestionData
	tmpl *template.Template
}

func NewParserEN(q *graphql.QuestionData) (p Parser) {
	return &en{
		data: q,
		tmpl: tmplEN,
	}
}

func (p *en) Parse(wr io.Writer) (err error) {
	err = p.tmpl.Execute(wr, p.data)
	if err != nil {
		return errors.Wrap(err, "fail execute template")
	}

	return nil
}

func (p *en) Filename() string {
	return "README_EN.md"
}
