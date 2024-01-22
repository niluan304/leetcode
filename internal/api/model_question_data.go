package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"unicode"

	"github.com/pkg/errors"
)

type (
	QuestionDataReq struct {
		TitleSlug string
	}

	// QuestionDataRes 模板数据
	QuestionDataRes struct {
		Question QuestionDataQuestion `json:"question"`
	}

	QuestionDataQuestion struct {
		QuestionId            string                    `json:"questionId"`            //
		QuestionFrontendId    string                    `json:"questionFrontendId"`    //
		CategoryTitle         string                    `json:"categoryTitle"`         //
		BoundTopicId          int                       `json:"boundTopicId"`          //
		Title                 string                    `json:"title"`                 //
		TitleSlug             string                    `json:"titleSlug"`             //
		Content               string                    `json:"content"`               //
		TranslatedTitle       string                    `json:"translatedTitle"`       //
		TranslatedContent     string                    `json:"translatedContent"`     //
		IsPaidOnly            bool                      `json:"isPaidOnly"`            //
		Difficulty            string                    `json:"difficulty"`            //
		Likes                 int                       `json:"likes"`                 //
		Dislikes              int                       `json:"dislikes"`              //
		IsLiked               interface{}               `json:"isLiked"`               //
		SimilarQuestions      string                    `json:"similarQuestions"`      // json字符串, 需要再次反序列化, 见: SimilarQuestion
		Contributors          []interface{}             `json:"contributors"`          //
		LangToValidPlayground string                    `json:"langToValidPlayground"` // json字符串, 需要再次反序列化, 见: LangToValidPlayground
		TopicTags             []QuestionDataTopicTag    `json:"topicTags"`             //
		CompanyTagStats       interface{}               `json:"companyTagStats"`       //
		CodeSnippets          []QuestionDataCodeSnippet `json:"codeSnippets"`          //
		Stats                 string                    `json:"stats"`                 // json字符串, 需要再次反序列化, 见: Stats
		Hints                 []string                  `json:"hints"`                 //
		Solution              QuestionDataSolution      `json:"solution"`              //
		Status                interface{}               `json:"status"`                //
		SampleTestCase        string                    `json:"sampleTestCase"`        //
		MetaData              string                    `json:"metaData"`              //
		JudgerAvailable       bool                      `json:"judgerAvailable"`       //
		JudgeType             string                    `json:"judgeType"`             //
		MysqlSchemas          []interface{}             `json:"mysqlSchemas"`          //
		EnableRunCode         bool                      `json:"enableRunCode"`         //
		EnvInfo               string                    `json:"envInfo"`               // json字符串, 需要再次反序列化, 见: EnvInfo
		Book                  interface{}               `json:"book"`                  //
		IsSubscribed          bool                      `json:"isSubscribed"`          //
		IsDailyQuestion       bool                      `json:"isDailyQuestion"`       //
		DailyRecordStatus     interface{}               `json:"dailyRecordStatus"`     //
		EditorType            string                    `json:"editorType"`            //
		UgcQuestionId         interface{}               `json:"ugcQuestionId"`         //
		Style                 string                    `json:"style"`                 //
		ExampleTestcases      string                    `json:"exampleTestcases"`      //
		JsonExampleTestcases  string                    `json:"jsonExampleTestcases"`  //
		Typename              string                    `json:"__typename"`            //
	}

	QuestionDataTopicTag struct {
		Name           string `json:"name"`
		Slug           string `json:"slug"`
		TranslatedName string `json:"translatedName"`
		Typename       string `json:"__typename"`
	}

	QuestionDataCodeSnippet struct {
		Lang     string `json:"lang"`
		LangSlug string `json:"langSlug"`
		Code     string `json:"code"`
		Typename string `json:"__typename"`
	}

	QuestionDataSolution struct {
		Id           string `json:"id"`
		CanSeeDetail bool   `json:"canSeeDetail"`
		Typename     string `json:"__typename"`
	}
)

type (
	SimilarQuestion struct {
		Title           string `json:"title"`
		TitleSlug       string `json:"titleSlug"`
		Difficulty      string `json:"difficulty"`
		TranslatedTitle string `json:"translatedTitle"`
		IsPaidOnly      any    `json:"isPaidOnly"`
	}

	Stats struct {
		TotalAccepted      string `json:"totalAccepted"`
		TotalSubmission    string `json:"totalSubmission"`
		TotalAcceptedRaw   int    `json:"totalAcceptedRaw"`
		TotalSubmissionRaw int    `json:"totalSubmissionRaw"`
		AcRate             string `json:"acRate"`
	}

	MetaData struct {
		Name   string     `json:"name"` // 函数名
		Params []struct { // 参数列表
			Name string `json:"name"` // 参数名
			Type string `json:"type"` // 参数类型
		} `json:"params"`
		Return struct { // 返回值
			Type string `json:"type"` // 返回值类型
		} `json:"return"`

		Classname   string `json:"classname"`
		Constructor struct {
			Params []struct {
				Name string `json:"name"` // 参数名
				Type string `json:"type"` // 参数类型
			} `json:"params"`
		} `json:"constructor"`
		Methods []struct {
			Params []struct {
				Type string `json:"type"`
				Name string `json:"name"`
			} `json:"params"`
			Return struct {
				Type string `json:"type"`
			} `json:"return"`
			Name string `json:"name"`
		} `json:"methods"`

		Systemdesign bool `json:"systemdesign"`
	}
	LangToValidPlayground map[string]bool

	EnvInfo map[string][]string

	ExampleTestcase any

	QuestionData struct {
		QuestionId            string                    `json:"questionId"`            //
		QuestionFrontendId    string                    `json:"questionFrontendId"`    //
		CategoryTitle         string                    `json:"categoryTitle"`         //
		BoundTopicId          int                       `json:"boundTopicId"`          //
		Title                 string                    `json:"title"`                 //
		TitleSlug             string                    `json:"titleSlug"`             //
		Content               string                    `json:"content"`               //
		TranslatedTitle       string                    `json:"translatedTitle"`       //
		TranslatedContent     string                    `json:"translatedContent"`     //
		IsPaidOnly            bool                      `json:"isPaidOnly"`            //
		Difficulty            string                    `json:"difficulty"`            //
		Likes                 int                       `json:"likes"`                 //
		Dislikes              int                       `json:"dislikes"`              //
		IsLiked               interface{}               `json:"isLiked"`               //
		SimilarQuestions      []SimilarQuestion         `json:"similarQuestions"`      //
		Contributors          []interface{}             `json:"contributors"`          //
		LangToValidPlayground LangToValidPlayground     `json:"langToValidPlayground"` //
		TopicTags             []QuestionDataTopicTag    `json:"topicTags"`             //
		CompanyTagStats       interface{}               `json:"companyTagStats"`       //
		CodeSnippets          []QuestionDataCodeSnippet `json:"codeSnippets"`          //
		Stats                 Stats                     `json:"stats"`                 //
		Hints                 []string                  `json:"hints"`                 //
		Solution              QuestionDataSolution      `json:"solution"`              //
		Status                interface{}               `json:"status"`                //
		SampleTestCase        string                    `json:"sampleTestCase"`        //
		MetaData              MetaData                  `json:"metaData"`              //
		JudgerAvailable       bool                      `json:"judgerAvailable"`       //
		JudgeType             string                    `json:"judgeType"`             //
		MysqlSchemas          []interface{}             `json:"mysqlSchemas"`          //
		EnableRunCode         bool                      `json:"enableRunCode"`         //
		EnvInfo               EnvInfo                   `json:"envInfo"`               //
		Book                  interface{}               `json:"book"`                  //
		IsSubscribed          bool                      `json:"isSubscribed"`          //
		IsDailyQuestion       bool                      `json:"isDailyQuestion"`       //
		DailyRecordStatus     interface{}               `json:"dailyRecordStatus"`     //
		EditorType            string                    `json:"editorType"`            //
		UgcQuestionId         interface{}               `json:"ugcQuestionId"`         //
		Style                 string                    `json:"style"`                 //
		ExampleTestcases      []ExampleTestcase         `json:"exampleTestcases"`      //
		JsonExampleTestcases  string                    `json:"jsonExampleTestcases"`  //
		Typename              string                    `json:"__typename"`            //

	}
)

func (q QuestionDataQuestion) ReUnmarshal() (*QuestionData, error) {
	var (
		similarQuestions []SimilarQuestion
		meta             MetaData
		validPlayground  LangToValidPlayground
		envInfo          EnvInfo
		stats            Stats
		exampleTestcases []ExampleTestcase

		err error
	)

	// 有 \n 的, 可能是 class
	exampleTestCaseData := fmt.Sprintf("[%s]", strings.ReplaceAll(q.ExampleTestcases, "\n", ","))

	var list = []struct {
		Data  string
		Point any
	}{
		{Data: q.SimilarQuestions, Point: &similarQuestions},
		{Data: q.MetaData, Point: &meta},
		{Data: q.LangToValidPlayground, Point: &validPlayground},
		{Data: q.EnvInfo, Point: &envInfo},
		{Data: exampleTestCaseData, Point: &exampleTestcases},
	}

	for _, item := range list {
		var (
			data    = item.Data
			pointer = item.Point
		)
		if data == "" {
			fmt.Println("data is empty", reflect.ValueOf(pointer).String())
			continue
		}

		err = json.Unmarshal([]byte(data), pointer)
		if err != nil {
			return nil, errors.Wrap(err, "fail unmarshal")
		}
	}

	return &QuestionData{
		QuestionId:            q.QuestionId,
		QuestionFrontendId:    q.QuestionFrontendId,
		CategoryTitle:         q.CategoryTitle,
		BoundTopicId:          q.BoundTopicId,
		Title:                 q.Title,
		TitleSlug:             q.TitleSlug,
		Content:               q.Content,
		TranslatedTitle:       q.TranslatedTitle,
		TranslatedContent:     q.TranslatedContent,
		IsPaidOnly:            q.IsPaidOnly,
		Difficulty:            q.Difficulty,
		Likes:                 q.Likes,
		Dislikes:              q.Dislikes,
		IsLiked:               q.IsLiked,
		SimilarQuestions:      similarQuestions,
		Contributors:          q.Contributors,
		LangToValidPlayground: validPlayground,
		TopicTags:             q.TopicTags,
		CompanyTagStats:       q.CompanyTagStats,
		CodeSnippets:          q.CodeSnippets,
		Stats:                 stats,
		Hints:                 q.Hints,
		Solution:              q.Solution,
		Status:                q.Status,
		SampleTestCase:        q.SampleTestCase,
		MetaData:              meta,
		JudgerAvailable:       q.JudgerAvailable,
		JudgeType:             q.JudgeType,
		MysqlSchemas:          q.MysqlSchemas,
		EnableRunCode:         q.EnableRunCode,
		EnvInfo:               envInfo,
		Book:                  q.Book,
		IsSubscribed:          q.IsSubscribed,
		IsDailyQuestion:       q.IsDailyQuestion,
		DailyRecordStatus:     q.DailyRecordStatus,
		EditorType:            q.EditorType,
		UgcQuestionId:         q.UgcQuestionId,
		Style:                 q.Style,
		ExampleTestcases:      exampleTestcases,
		JsonExampleTestcases:  q.JsonExampleTestcases,
		Typename:              q.Typename,
	}, nil
}

func (q QuestionData) Dir() string {
	frontendId := strings.ReplaceAll(q.QuestionFrontendId, " ", "")
	titleSlug := q.TitleSlug
	if !isStringNumeric(frontendId) {
		titleSlug = camel2Case(q.MetaData.Name, "-")
		frontendId = q.CategoryTitle
	}

	return frontendId + "_" + titleSlug
}

func (q QuestionData) Pkg() string {
	frontendId := strings.ReplaceAll(q.QuestionFrontendId, " ", "")
	titleSlug := q.TitleSlug
	if !isStringNumeric(frontendId) {
		titleSlug = camel2Case(q.MetaData.Name, "_")
	}

	titleSlug = strings.ReplaceAll(titleSlug, "-", "_")
	return titleSlug
}

func (q QuestionData) NeedMode() bool {
	content := q.TranslatedContent
	return strings.Contains(content, "取余") ||
		strings.Contains(content, "取模") ||
		strings.Contains(content, "答案可能很大")
}

func (q QuestionData) NeedDefinition() bool {
	for _, item := range q.CodeSnippets {
		if strings.Contains(item.Code, " * Definition for ") {
			return true
		}
	}
	return false
}

func (q QuestionData) CodeSnippet(langSlug string) string {
	for _, item := range q.CodeSnippets {
		if item.LangSlug == langSlug {
			return item.Code
		}
	}
	return ""
}

func camel2Case(str string, with string) string {
	buffer := bytes.Buffer{}
	for i, r := range str {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.WriteString(with)
			}
			buffer.WriteRune(unicode.ToLower(r))
		} else {
			buffer.WriteRune(r)
		}
	}
	return buffer.String()
}

func isStringNumeric(str string) bool {
	for _, char := range str {
		if '0' > char || char > '9' {
			return false
		}
	}

	return true
}
