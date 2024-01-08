package tmpl

import (
	"text/template"

	"leetcode/cmd/leetcode/graphql"
)

func NewParserCase(q *graphql.QuestionData, tmpl *template.Template) (p Parser) {
	type Param struct {
		Name string
		Type string
	}

	type Data struct {
		PkgName string
		Params  []Param
		Return  string
	}

	var params []Param
	for _, param := range q.MetaData.Params {
		params = append(params, Param{
			Name: param.Name,
			Type: GoType(param.Type),
		})
	}

	data := Data{
		PkgName: "main",
		Params:  params,
		Return:  GoType(q.MetaData.Return.Type),
	}

	return NewParser("case.go", tmpl, data)
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

		//// tree or listNode
		//"TreeNode": "*structs.TreeNode",
		//"ListNode": "*structs.ListNode",
	}

	if v, ok := m[t]; ok {
		return v
	}
	return t
}

func NewParserZH(q *graphql.QuestionData, tmpl *template.Template) (p Parser) {
	return NewParser("README.md", tmpl, q)
}

func NewParserEN(q *graphql.QuestionData, tmpl *template.Template) (p Parser) {
	return NewParser("README_EN.md", tmpl, q)
}
