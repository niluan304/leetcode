package tmpl

import (
	"text/template"
)

var (
	tmplZH = template.Must(template.New("zh").Parse(TemplateZH))
	tmplEN = template.Must(template.New("en").Parse(TemplateEN))
)

var (
	tmplCase = template.Must(template.New("en").Parse(TemplateCase))
	tmplTest = template.Must(template.New("en").Parse(TemplateTest))
)

// TemplateEN 题目描述 README 英文模板
const TemplateEN = `
| English | [简体中文](README.md) |

# [{{.QuestionFrontendId}}.{{.Title}}](https://leetcode.com/problems/{{.TitleSlug}}/)
Difficulty:{{.Difficulty}}, Likes: {{.Likes}}

## Description

# {{.Content}}

## Related Topics
{{range .TopicTags }}
- [{{.Name}}](https://leetcode-cn.com/tag/{{.Slug}}/)
{{- end }}

## Similar Questions
{{range .SimilarQuestions}}
- [{{.Title}}](../{{.TitleSlug}}/README.md) {{.Difficulty}} {{if.IsPaidOnly}}🔒{{end}}{{/* 是否付费 */}}
{{- end }}
`

// TemplateZH 题目描述 README 中文模板
const TemplateZH = `
| [English](README_EN.md) | 简体中文 |

# [{{.QuestionFrontendId}}. {{.TranslatedTitle}}](https://leetcode.cn/problems/{{.TitleSlug}}/)
Difficulty:{{.Difficulty}}, Likes: {{.Likes}}

## 题目描述

{{.TranslatedContent}}

## 相关话题
{{range .TopicTags}}
- [{{.TranslatedName}}](https://leetcode-cn.com/tag/{{.Slug}}/)
{{- end}}

## 相似题目
{{range .SimilarQuestions}}
- [{{.TranslatedTitle}}](../{{.TitleSlug}}/README.md) {{.Difficulty}} {{if.IsPaidOnly}}🔒{{end}}{{/* 是否付费 */}}
{{- end}}
`

const TemplateCase = `package {{.PkgName}}

import (
	"testing"

	"leetcode/tests"
)

type Input struct {
{{- range .Params}}
	{{.Name}} {{.Type}}
{{- end}}
}

type Output {{.Return}}

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		// add self case
	}
}

type Func func({{- range .Params}}{{.Name}} {{.Type}}, {{- end}}) {{.Return}}

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			{{- range .Params}}
			in.{{.Name}},
			{{- end}}
		))
	}
}

func AddCases(c func() []tests.Case[Input, Output]) {
	_cases := cases()
	cases = func() []tests.Case[Input, Output] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...Func) {
	funcs = append(funcs, tests.NewFuncWithAdaptor(adaptor, f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, cases, funcs...)
}

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, cases, funcs...)
}
`

const TemplateTest = `package {{.PkgName}}

import (
	"testing"

	"leetcode/tests"
)

func init() {

	AddCases(func() []tests.Case[Input, Output] {
		return []tests.Case[Input, Output]{
			// Add self case
		}
	})

	AddFunc(
		{{.Name}}, // use func in case.go
	)
}

func Test_{{.Name}}(t *testing.T) {
	Unit(t)
}

func Benchmark_{{.Name}}(b *testing.B) {
	Bench(b)
}

`
