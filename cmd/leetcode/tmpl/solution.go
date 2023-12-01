package tmpl

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"

	"leetcode/cmd/leetcode/graphql"
)

type Solution struct {
	CodeSnippets []graphql.QuestionDataCodeSnippet //

	NeedMod bool   //
	PkgName string //
}

func NewSolution(q *graphql.QuestionData) (*Solution, error) {
	//pkg := q.Pkg()

	return &Solution{
		CodeSnippets: q.CodeSnippets,
		PkgName:      "main",
		NeedMod: strings.Contains(q.TranslatedContent, "取余") ||
			strings.Contains(q.TranslatedContent, "取模") ||
			strings.Contains(q.TranslatedContent, "答案可能很大"),
	}, nil
}

func (s *Solution) Save(root string) (err error) {
	_ = os.MkdirAll(root, os.ModePerm)

	for _, item := range s.CodeSnippets {
		ext := extname(item.LangSlug)
		prefix := ""
		if ext == "go" {
			prefix = fmt.Sprintf("package %s\n\n", s.PkgName)
		} else {
			// TODO 暂时硬编码, 后续改为读取配置文件
			continue
		}

		filename := "solution" + "." + ext
		err := os.WriteFile(filepath.Join(root, filename), []byte(prefix+item.Code), os.ModePerm)
		if err != nil {
			return errors.Wrap(err, "fail write: "+filename)
		}
	}

	return nil
}

// extname  get extension name
func extname(langSlug string) string {
	switch langSlug {
	case "csharp":
		return "cs"
	case "golang":
		return "go"
	case "javascript":
		return "js"
	case "kotlin":
		return "kt"
	case "python", "python3":
		return "py"
	case "racket":
		return "rkt"
	case "ruby":
		return "rb"
	case "rust":
		return "rs"
	case "typescript":
		return "ts"

	default:
		// ext name is same as langSlug
		// eg: c, cpp, dart, java, javascript, kotlin, scala, swift
		return langSlug
	}
}

func NewParserSolution(q *graphql.QuestionData) Parser {
	return &Solution{
		CodeSnippets: q.CodeSnippets,
		PkgName:      "main",
		NeedMod: strings.Contains(q.TranslatedContent, "取余") ||
			strings.Contains(q.TranslatedContent, "取模") ||
			strings.Contains(q.TranslatedContent, "答案可能很大"),
	}
}

func (s *Solution) Parse(w io.Writer) (err error) {
	for _, item := range s.CodeSnippets {
		ext := extname(item.LangSlug)
		if ext != "go" {
			continue
		}

		data := fmt.Sprintf("package main\n\n")
		if strings.Contains(item.Code, "Definition for") {
			data += `import(
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

`
		}

		if s.NeedMod {
			data += `const mod = 1e9 + 7` + "\n\n"
		}

		data += item.Code

		_, err = w.Write([]byte(data))
		if err != nil {
			return errors.Wrap(err, "fail write solution")
		}
	}
	return nil
}

func (s *Solution) Filename() string {
	return "solution.go"
}

type leetcode struct {
	PkgName string
}

func NewParserLeetcode(pkgName string) (p Parser) {
	return &leetcode{
		PkgName: "main",
	}
}

func (p leetcode) Parse(w io.Writer) (err error) {
	_, err = w.Write([]byte(fmt.Sprintf("package main\n")))
	if err != nil {
		return errors.Wrap(err, "fail write solution_leetcode")
	}

	return nil
}

func (p leetcode) Filename() string {
	return "solution_leetcode.go"
}

func NewParserUnitCase(q *graphql.QuestionData) (p Parser) {
	type Data struct {
		PkgName string
		Name    string
	}

	data := &Data{
		PkgName: "main",
		Name:    q.MetaData.Name,
	}

	return NewParser("solution_test.go", tmplTest, data)
}

func NewParserEndlessTest(q *graphql.QuestionData) (p Parser) {
	type Data struct {
		PkgName     string
		Name        string
		RunFuncName string
		Type        string
	}

	code := "interface{}"
	for _, item := range q.CodeSnippets {
		ext := extname(item.LangSlug)
		if ext != "go" {
			continue
		}
		code = item.Code

		i := strings.Index(code, "(")
		j := strings.Index(code, "{")
		code = code[:5] + code[i:j]
	}

	data := &Data{
		PkgName:     "main",
		Name:        q.MetaData.Name,
		RunFuncName: "RunLeetCodeFuncWithFile",
		Type:        code,
	}
	if q.MetaData.Systemdesign {
		data = &Data{
			PkgName:     data.PkgName,
			Name:        "Constructor",
			RunFuncName: "RunLeetCodeClassWithFile",
			Type:        "interface{}",
		}
	}

	return NewParser("solution_test.go", endlessTest, data)
}
