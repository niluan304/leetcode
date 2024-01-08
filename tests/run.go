package tests

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func funcName(f any) string {
	name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	i := strings.LastIndex(name, ".")
	return name[i+1:]
}

func RunFunc(t *testing.T, f any, data string, targetCaseNum int) (err error) {
	lines := trimSpaceAndEmptyLine(data)
	n := len(lines)
	if n == 0 {
		return fmt.Errorf("输入数据为空，请检查文件路径和文件内容是否正确")
	}

	fType := reflect.TypeOf(f)
	if fType.Kind() != reflect.Func {
		return fmt.Errorf("f 必须是函数")
	}

	// 每 fNumIn+fNumOut 行一组数据
	fNumIn := fType.NumIn()
	fNumOut := fType.NumOut()
	tcSize := fNumIn + fNumOut
	if n%tcSize != 0 {
		return fmt.Errorf("有效行数 %d，应该是 %d 的倍数", n, tcSize)
	}

	examples := make([][]string, 0, n/tcSize)
	for i := 0; i < n; i += tcSize {
		examples = append(examples, lines[i:i+tcSize])
	}

	t.Run(funcName(f), func(t *testing.T) {
		err = testutil.RunLeetCodeFuncWithExamples(t, f, examples, targetCaseNum)
	})
	return err
}

func RunClass(t *testing.T, f any, data string, targetCaseNum int) (err error) {
	lines := trimSpaceAndEmptyLine(data)
	n := len(lines)
	if n == 0 {
		return fmt.Errorf("输入数据为空，请检查文件路径和文件内容是否正确")
	}

	// 每三行一组数据
	if n%3 != 0 {
		return fmt.Errorf("有效行数 %d，应该是 3 的倍数", n)
	}

	examples := make([][3]string, 0, n/3)
	for i := 0; i < n; i += 3 {
		examples = append(examples, [3]string{lines[i], lines[i+1], lines[i+2]})
	}

	t.Run(funcName(f), func(t *testing.T) {
		err = testutil.RunLeetCodeClassWithExamples(t, f, examples, targetCaseNum)
	})
	return err
}

// 移除多余空行和左右空格
func trimSpaceAndEmptyLine(s string) (res []string) {
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			res = append(res, line)
		}
	}
	return
}
