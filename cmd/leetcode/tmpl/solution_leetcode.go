package tmpl

import (
	"fmt"
	"io"

	"github.com/pkg/errors"
)

type leetcode struct {
	PkgName string
}

func NewParserLeetcode(pkgName string) (p Parser) {
	return &leetcode{
		PkgName: pkgName,
	}
}

func (p leetcode) Parse(wr io.Writer) (err error) {
	_, err = wr.Write([]byte(fmt.Sprintf("package %s\n", p.PkgName)))
	if err != nil {
		return errors.Wrap(err, "fail write solution_leetcode")
	}

	return nil
}

func (p leetcode) Filename() string {
	return "solution_leetcode.go"
}
