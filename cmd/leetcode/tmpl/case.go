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
	Params  []Param
	Return  string
}

type Param struct {
	Name string
	Type string
}

func NewParserCase(q *graphql.QuestionData) (p Parser) {
	var params []Param

	for _, param := range q.MetaData.Params {
		params = append(params, Param{
			Name: param.Name,
			Type: GoType(param.Type),
		})
	}
	data := caseData{
		PkgName: q.Pkg(),
		Params:  params,
		Return:  GoType(q.MetaData.Return.Type),
	}

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

func GoType(t string) string {
	m := map[string]string{
		"integer":   "int",
		"integer[]": "[]int",
		"string":    "string",
		"string[]":  "[]string",
		"boolean":   "bool",
		"boolean[]": "[]bool",
		"double":    "float64",
		"double[]":  "[]float64",

		// array
		"list<list<integer>>": "[][]int",
		"list<integer>":       "[]int",
		"list<string>":        "[]string",

		// tree or listNode
		"TreeNode": "*structs.TreeNode",
		"ListNode": "*structs.ListNode",
	}

	if v, ok := m[t]; ok {
		return v
	}
	return t
}
