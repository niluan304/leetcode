package tmpl

import (
	"strings"

	"github.com/niluan304/leetcode/internal/graphql"
)

type EndlessTest struct {
	Name           string
	RunType        string
	Type           string
	Samples        string
	NeedDefinition bool

	Data *graphql.QuestionData
}

type Leetcode struct {
	NeedDefinition bool // 是否需要导入 ListNode/TreeNode 等结构体

	Data *graphql.QuestionData // 可用于修改包名
}

type Solution struct {
	Code           string // 模板代码
	NeedMod        bool   // 是否需要取模
	NeedDefinition bool   // 是否需要取模

	Data *graphql.QuestionData // 可用于修改包名
}

func (t *Template) Samples() Samples {
	q := t.data
	content := q.Content
	if len(content) < 100 {
		// TODO only Chinese
		//content = q.TranslatedContent
	}
	return NewSamples(content, q.MetaData.Systemdesign)
}

const langSlug = "golang"

func (t *Template) Leetcode() *Leetcode {
	q := t.data
	codeSnippet := q.CodeSnippet(langSlug)
	return &Leetcode{
		NeedDefinition: strings.Contains(codeSnippet, "Definition for"),
		Data:           q,
	}
}

func (t *Template) EndlessTest() (data EndlessTest) {
	q := t.data

	codeSnippet := q.CodeSnippet(langSlug)
	idx := strings.Index(codeSnippet, "func ")

	funcType := codeSnippet[idx:]
	i := strings.Index(funcType, "(")
	j := strings.Index(funcType, "{")
	funcType = funcType[:5] + funcType[i:j]

	data = EndlessTest{
		Name:           q.MetaData.Name,
		RunType:        "Func",
		Type:           funcType,
		Samples:        "`\n" + t.Samples().String() + "\n`",
		NeedDefinition: q.NeedDefinition(),
		Data:           q,
	}

	if q.MetaData.Systemdesign {
		data.Name = "Constructor"
		data.RunType = "Class"
		data.Type = "interface{}"
	}

	return data
}

func (t *Template) Solution() *Solution {
	q := t.data

	code := q.CodeSnippet(langSlug)
	i := strings.LastIndex(code, "}")
	if i >= 0 {
		code = code[:i] + "\tvar ans = \n\n\treturn ans\n" + code[i:]
	}

	return &Solution{
		Code:           code,
		NeedMod:        q.NeedMode(),
		NeedDefinition: q.NeedDefinition(),
		Data:           t.data,
	}
}
