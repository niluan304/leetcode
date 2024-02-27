package tmpl

import (
	"fmt"
	"strings"

	"github.com/niluan304/leetcode/internal/api"
)

type EndlessTest struct {
	Name           string
	RunType        string
	Type           string
	Samples        string
	NeedDefinition bool

	Data *api.QuestionData
}

type Leetcode struct {
	NeedDefinition bool // 是否需要导入 ListNode/TreeNode 等结构体

	Data *api.QuestionData // 可用于修改包名
}

type Solution struct {
	Code           string // 模板代码
	NeedMod        bool   // 是否需要取模
	NeedDefinition bool   // 是否需要取模

	Data *api.QuestionData // 可用于修改包名
}

func (t *Template) Samples() Samples {
	q := t.data
	content := q.Content
	if len(content) < 100 {
		// TODO only Chinese
		// content = q.TranslatedContent
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
		Type:           strings.Replace(funcType, "func (", "func(", 1),
		Samples:        t.Samples().String(),
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
	if !q.MetaData.Systemdesign {
		ansType := funcReturnType(code)

		i := strings.LastIndex(code, "}")
		varAns := "var ans = 0 // math.MaxInt32 // math.MinInt32"
		ans := "ans"
		switch ansType {
		case "int", "float64":
			// initValue = "var ans = 0 \n\t //var ans = math.MaxInt32\n\t//var ans = math.MinInt32"
		case "int64":
			// initValue = "var ans = 0 \n\t //var ans = math.MaxInt32\n\t//var ans = math.MinInt32"
			ans = "int64(ans)"
		case "[]int", "[]byte", "[][]int":
			varAns = fmt.Sprintf("var ans = make(%s, n)", ansType)
		default:
			varAns = "var ans " + ansType
		}

		code = fmt.Sprintf(`%s

	// var n = len()
	%s

	return %s
%s
`, strings.TrimRight(code[:i], "\n"), varAns, ans, code[i:])
	} else {
		const this = "(this "
		i := strings.Index(code, this)
		j := strings.Index(code[i:], ")")
		m := strings.ReplaceAll(strings.ToLower(code[i+len(this):i+j]), "*", "")
		code = strings.ReplaceAll(code, this, "("+m[:1]+" ")
	}

	return &Solution{
		Code:           code,
		NeedMod:        q.NeedMode(),
		NeedDefinition: q.NeedDefinition(),
		Data:           t.data,
	}
}

func funcReturnType(code string) string {
	i := strings.LastIndexByte(code, ')')
	j := strings.LastIndex(code, "{")
	return strings.TrimSpace(code[i+1 : j])
}
