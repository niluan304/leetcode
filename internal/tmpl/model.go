package tmpl

import (
	"fmt"
	"strings"

	"github.com/niluan304/leetcode/internal/api"
)

// Data 用于渲染模板文件的数据
type Data struct {
	// for code snippet
	Solution

	// for unit test
	UnitTest

	Data *api.QuestionData
}

type UnitTest struct {
	RunType  string // Class or Func
	FuncName string
	FuncType string
	Samples  string
}

type Solution struct {
	Code           string // 模板代码
	NeedMod        bool   // 是否需要取模
	NeedDefinition bool   // 是否需要取模
}

func NewData(data *api.QuestionData) Data {
	return Data{
		Solution: NewSolution(data),
		UnitTest: NewUnitTest(data),
		Data:     data,
	}
}

func NewSamples(q *api.QuestionData) Samples {
	content := q.Content
	if len(content) < 100 {
		// TODO only Chinese
		// content = q.TranslatedContent
	}
	return newSamples(content, q.MetaData.Systemdesign)
}

const langSlug = "golang"

func NewUnitTest(q *api.QuestionData) (data UnitTest) {
	codeSnippet := q.CodeSnippet(langSlug)
	idx := strings.Index(codeSnippet, "func ")

	funcType := codeSnippet[idx:]
	i := strings.Index(funcType, "(")
	j := strings.Index(funcType, "{")
	funcType = funcType[:5] + funcType[i:j]

	data = UnitTest{
		FuncName: q.MetaData.Name,
		RunType:  "Func",
		FuncType: strings.Replace(funcType, "func (", "func(", 1),
		Samples:  NewSamples(q).String(),
	}

	if q.MetaData.Systemdesign {
		data.RunType = "Class"
		data.FuncName = "Constructor"
		data.FuncType = "interface{}"
	}

	return data
}

func NewSolution(q *api.QuestionData) Solution {
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

	return Solution{
		Code:           code,
		NeedMod:        q.NeedMode(),
		NeedDefinition: q.NeedDefinition(),
	}
}

func funcReturnType(code string) string {
	i := strings.LastIndexByte(code, ')')
	j := strings.LastIndex(code, "{")
	return strings.TrimSpace(code[i+1 : j])
}
