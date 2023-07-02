package tmpl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSolution_Save(t *testing.T) {
	var args = []struct {
		input  string
		output string
	}{
		{
			input:  "./data/two-sum.json",
			output: "./data/1_two-sum",
		},
		{
			input:  "./data/flip_chess.json",
			output: "./data/LCP41_flip-chess",
		},
	}

	for _, arg := range args {
		input := arg.input

		question, err := LoadQuestion(input)
		if err != nil {
			t.Errorf("read file error: %v", err)
			return
		}

		output := filepath.Join("./data", question.Dir())
		if output != arg.output {
			t.Logf("output error: %s != %s", output, arg.output)
		}
		_ = os.MkdirAll(output, os.ModePerm)

		// 新建各语言solution文件
		s, err := NewSolution(question)
		if err != nil {
			t.Error(err)
			return
		}

		err = s.Save(output)
		if err != nil {
			t.Error(err)
			return
		}
	}
}
