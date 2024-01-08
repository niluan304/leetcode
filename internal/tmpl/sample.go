package tmpl

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Sample struct {
	Input  []string
	Output []string
}

type Samples []Sample

func (s Samples) String() (samples string) {
	for _, sample := range s {
		samples += strings.Join(sample.Input, "\n") + "\n" + strings.Join(sample.Output, "\n") + "\n\n"
	}
	return
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
