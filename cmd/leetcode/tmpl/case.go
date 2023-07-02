package tmpl

import (
	"io"
	"text/template"

	"github.com/pkg/errors"

	"leetcode/cmd/leetcode/graphql"
)

type _case struct {
	data caseData
	tmpl *template.Template
}

type caseData struct {
	PkgName string
	Params  []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}
	Return string
}

func NewParserCase(q *graphql.QuestionData) (p Parser) {
	data := caseData{
		PkgName: q.Pkg(),
		Params:  q.MetaData.Params,
		Return:  q.MetaData.Return.Type,
	}

	// TODO 再映射返回值类型

	return &_case{
		data: data,
		tmpl: tmplCase,
	}
}

func (p *_case) Parse(wr io.Writer) (err error) {
	err = p.tmpl.Execute(wr, p.data)
	if err != nil {
		return errors.Wrap(err, "fail execute template")
	}

	return nil
}

func (p *_case) Filename() string {
	return "case.go"
}
