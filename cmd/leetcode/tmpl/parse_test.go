package tmpl

import (
	"os"
	"path/filepath"
	"testing"
)

func Test_Parser(t *testing.T) {
	var args = []struct {
		input  string
		output string
	}{
		{
			input:  "D:/__Project/leetcode/questionData/two-sum.json",
			output: "./data/1_two-sum",
		},
		{
			input:  "D:/__Project/leetcode/questionData/3Etpl5.json",
			output: "./data/offer_LCOF2",
		},
		{
			input:  "D:/__Project/leetcode/questionData/container-with-most-water.json",
			output: "./data/container-with-most-water",
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

		// 解析模板
		t.Run(input, func(t *testing.T) {
			list := []Parser{
				NewParserEN(question),
				NewParserZH(question),
				NewParserCase(question),
				NewParserSolution(question),
				NewParserUnitCase(question),
				NewParserLeetcode(question.Pkg()),
			}

			for _, elem := range list {
				t.Run(elem.Filename(), func(t *testing.T) {
					p := Parse{
						Parser: elem,
					}

					err = p.Save(output)
					if err != nil {
						t.Errorf("save error: %v", err)
					}
				})
			}
		})
	}
}
