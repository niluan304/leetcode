package tmpl

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"leetcode/cmd/leetcode/graphql"
)

type Solution struct {
	CodeSnippets []graphql.QuestionDataCodeSnippet //

	PkgName string //
}

func NewSolution(q *graphql.QuestionData) (*Solution, error) {
	pkg := q.Pkg()

	return &Solution{
		CodeSnippets: q.CodeSnippets,
		PkgName:      pkg,
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
	pkg := q.Pkg()

	return &Solution{
		CodeSnippets: q.CodeSnippets,
		PkgName:      pkg,
	}
}

func (s *Solution) Parse(wr io.Writer) (err error) {
	for _, item := range s.CodeSnippets {
		ext := extname(item.LangSlug)
		if ext != "go" {
			continue
		}

		_, err = wr.Write([]byte(fmt.Sprintf("package %s\n\n%s", s.PkgName, item.Code)))
		if err != nil {
			return errors.Wrap(err, "fail write solution")
		}
	}
	return nil
}

func (s *Solution) Filename() string {
	return "solution.go"
}
