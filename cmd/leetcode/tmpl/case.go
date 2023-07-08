package tmpl

import (
	"leetcode/cmd/leetcode/graphql"
)

func NewParserCase(q *graphql.QuestionData) (p Parser) {
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
		PkgName: q.Pkg(),
		Params:  params,
		Return:  GoType(q.MetaData.Return.Type),
	}

	return NewParser("case.go", tmplCase, data)
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

func NewParserZH(q *graphql.QuestionData) (p Parser) {
	return NewParser("README.md", tmplZH, q)
}

func NewParserEN(q *graphql.QuestionData) (p Parser) {
	return NewParser("README_EN.md", tmplEN, q)
}
