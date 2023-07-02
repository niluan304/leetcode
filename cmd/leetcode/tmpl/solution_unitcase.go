package tmpl

import (
	"io"
	"text/template"

	"github.com/pkg/errors"

	"leetcode/cmd/leetcode/graphql"
)

type unitCase struct {
	data *unitCaseData
	tmpl *template.Template
}

type unitCaseData struct {
	PkgName string
	Name    string
}

func NewParserUnitCase(q *graphql.QuestionData) (p Parser) {
	data := &unitCaseData{
		PkgName: q.Pkg(),
		Name:    q.MetaData.Name,
	}

	return &unitCase{
		data: data,
		tmpl: tmplTest,
	}
}

func (p unitCase) Parse(wr io.Writer) (err error) {
	err = p.tmpl.Execute(wr, p.data)
	if err != nil {
		return errors.Wrap(err, "fail execute template")
	}

	return nil
}

func (p unitCase) Filename() string {
	return "solution_test.go"
}
