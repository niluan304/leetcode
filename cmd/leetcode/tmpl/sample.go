package tmpl

import (
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"leetcode/cmd/leetcode/graphql"
)

type Sample struct {
	Input  []string
	Output []string
}

type Samples []Sample

func (s Samples) Parse(w io.Writer) (err error) {
	for _, sample := range s {
		for _, input := range sample.Input {
			_, err = w.Write([]byte(input + "\n"))
			if err != nil {
				return
			}
		}
		for _, output := range sample.Output {
			_, err = w.Write([]byte(output + "\n"))
			if err != nil {
				return
			}
		}
		_, err = w.Write([]byte("\n"))
	}
	return err
}

func (s Samples) Filename() string {
	return "sample.txt"
}

func NewParserSamples(q *graphql.QuestionData) (p Parser) {
	content := q.Content
	if len(content) < 100 {
		// TODO only Chinese
		//content = q.TranslatedContent
	}
	return NewSamples(content, q.MetaData.Systemdesign)
}

func NewSamples(content string, systemDesign bool) (samples Samples) {
	content = strings.ReplaceAll(content, "put:</strong>", "put</strong>") // 设计题末尾没有 ':'
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		return nil
	}

	doc.Find("pre").Each(func(i int, pre *goquery.Selection) {
		var input []string
		var output []string

		pre.Find("strong:contains('Input')").Each(func(i int, selection *goquery.Selection) {
			text := selection.Nodes[0].NextSibling.Data
			if systemDesign {
				input = []string{strings.TrimSpace(text)}
				return
			}

			splits := strings.Split(text, "=")

			input = make([]string, 0, len(splits)-1)
			for _, s := range splits[1 : len(splits)-1] {
				end := strings.LastIndexByte(s, ',')
				input = append(input, strings.TrimSpace(s[:end]))
			}
			input = append(input, strings.TrimSpace(splits[len(splits)-1]))
		})

		pre.Find("strong:contains('Output')").Each(func(i int, selection *goquery.Selection) {
			text := selection.Nodes[0].NextSibling.Data
			output = append(output, strings.TrimSpace(text))
		})

		samples = append(samples, Sample{
			Input:  input,
			Output: output,
		})
	})

	return samples
}
