package tmpl

import (
	"text/template"
)

var (
	tmplZH = template.Must(template.New("zh").Parse(TemplateZH))
	tmplEN = template.Must(template.New("en").Parse(TemplateEN))
)

// TemplateEN 题目描述 README 英文模板
const TemplateEN = `
| English | [简体中文](README.md) |

# [{{.QuestionFrontendId}}.{{.Title}}](https://leetcode.com/problems/{{.TitleSlug}}/)
Difficulty:{{.Difficulty}}, Likes: {{.Likes}}

## Description

# {{.Content}}

## Related Topics
{{range .TopicTags }}
- [{{.Name}}](https://leetcode-cn.com/tag/{{.Slug}}/)
{{- end }}

## Similar Questions
{{range .SimilarQuestions}}
- [{{.Title}}](../{{.TitleSlug}}/README.md) {{.Difficulty}} {{if.IsPaidOnly}}🔒{{end}}{{/* 是否付费 */}}
{{- end }}
`

// TemplateZH 题目描述 README 中文模板
const TemplateZH = `
| [English](README_EN.md) | 简体中文 |

# [{{.QuestionFrontendId}}. {{.TranslatedTitle}}](https://leetcode.cn/problems/{{.TitleSlug}}/)
Difficulty:{{.Difficulty}}, Likes: {{.Likes}}

## 题目描述

{{.TranslatedContent}}

## 相关话题
{{range .TopicTags}}
- [{{.TranslatedName}}](https://leetcode-cn.com/tag/{{.Slug}}/)
{{- end}}

## 相似题目
{{range .SimilarQuestions}}
- [{{.TranslatedTitle}}](../{{.TitleSlug}}/README.md) {{.Difficulty}} {{if.IsPaidOnly}}🔒{{end}}{{/* 是否付费 */}}
{{- end}}
`
