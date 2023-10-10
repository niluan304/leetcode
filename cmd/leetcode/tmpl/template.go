package tmpl

import (
	"text/template"
)

var (
	tmplZH = template.Must(template.New("zh").Parse(TemplateZH))
	tmplEN = template.Must(template.New("en").Parse(TemplateEN))
)

var (
	tmplCase    = template.Must(template.New("case").Parse(TemplateCase))
	tmplTest    = template.Must(template.New("test").Parse(TemplateTest))
	endlessTest = template.Must(template.New("test").Parse(EndlessTest))
)

// TemplateEN 题目描述 README 英文模板
const TemplateEN = `
| English | [简体中文](README.md) |

# [{{.QuestionFrontendId}}.{{.Title}}](https://leetcode.com/problems/{{.TitleSlug}}/)
Difficulty:{{.Difficulty}}, Likes: {{.Likes}}

## Description

{{.Content}}

## Related Topics
{{range .TopicTags }}
- [{{.Name}}](https://leetcode.com/tag/{{.Slug}}/)
{{- end }}

## Similar Questions
{{range .SimilarQuestions}}
- [{{.Title}}](../{{.TitleSlug}}/README_EN.md) {{.Difficulty}} {{if.IsPaidOnly}}🔒{{end}}{{/* 是否付费 */}}
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
- [{{.TranslatedName}}](https://leetcode.cn/tag/{{.Slug}}/)
{{- end}}

## 相似题目
{{range .SimilarQuestions}}
- [{{.TranslatedTitle}}](../{{.TitleSlug}}/README.md) {{.Difficulty}} {{if.IsPaidOnly}}🔒{{end}}{{/* 是否付费 */}}
{{- end}}
`

const TemplateCase = `package {{.PkgName}}

import (
	"reflect"
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
		// TODO add question case
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

func Equal(x, y Output) bool {

	return reflect.DeepEqual(x, y)
}

func Unit(t *testing.T) {
	tests.Unit(t, tests.Test[Input, Output]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  Equal,
	})
}

func Bench(b *testing.B) {
	tests.Bench(b, tests.Test[Input, Output]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  Equal,
	})
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

const EndlessTest = `package {{.PkgName}}

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func Test_{{.PkgName}}(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		{{.Name}},
	}

	for _, f := range fs{
		name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		i := strings.LastIndex(name, ".")

		t.Run(name[i+1:], func(t *testing.T) {
			err := testutil.{{.RunFuncName}}(t, f, "sample.txt", targetCaseNum) 
			if err != nil {
				t.Error(err)
			}
		})
	}
}
`
