package graphql

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/machinebox/graphql"
)

func (c *Client) ConsolePanelConfig(ctx context.Context, titleSlug string) (res *ConsolePanelConfigRes, err error) {
	req := graphql.NewRequest(`{
"query": "
query consolePanelConfig($titleSlug: String!) {
 question(titleSlug: $titleSlug) {
   questionId
   questionFrontendId
   questionTitle
   enableRunCode
   enableSubmit
   enableTestMode
   jsonExampleTestcases
   exampleTestcases
   metaData
   sampleTestCase
 }
}
"
}`)
	req.Var("titleSlug", titleSlug)
	req.Var("operationName", "consolePanelConfig")

	err = c.request(ctx, req, titleSlug, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) QuestionData(ctx context.Context, titleSlug string) (res *QuestionDataRes, err error) {
	req := graphql.NewRequest(`query questionData($titleSlug: String!) {
  question(titleSlug: $titleSlug) {
    questionId
    questionFrontendId
    categoryTitle
    boundTopicId
    title
    titleSlug
    content
    translatedTitle
    translatedContent
    isPaidOnly
    difficulty
    likes
    dislikes
    isLiked
    similarQuestions
    contributors {
      username
      profileUrl
      avatarUrl
      __typename
    }
    langToValidPlayground
    topicTags {
      name
      slug
      translatedName
      __typename
    }
    companyTagStats
    codeSnippets {
      lang
      langSlug
      code
      __typename
    }
    stats
    hints
    solution {
      id
      canSeeDetail
      __typename
    }
    status
    sampleTestCase
    metaData
    judgerAvailable
    judgeType
    mysqlSchemas
    enableRunCode
    envInfo
    book {
      id
      bookName
      pressName
      source
      shortDescription
      fullDescription
      bookImgUrl
      pressImgUrl
      productUrl
      __typename
    }
    isSubscribed
    isDailyQuestion
    dailyRecordStatus
    editorType
    ugcQuestionId
    style
    exampleTestcases
    jsonExampleTestcases
    __typename
  }
}
`)
	req.Var("titleSlug", titleSlug)
	req.Var("operationName", "questionData")

	err = c.request(ctx, req, titleSlug, &res)
	if err != nil {
		return nil, err
	}

	{
		err := c.save("questionData", titleSlug, res)
		if err != nil {
			log.Println("fail save file:", titleSlug, err)
		}
	}

	return res, nil
}

func (c *Client) save(path string, titleSlug string, res any) (err error) {
	data, err := json.Marshal(res)
	if err != nil {
		return err
	}
	_ = os.MkdirAll(path, 755)
	err = os.WriteFile(filepath.Join(path, titleSlug+".json"), data, 0644)
	if err != nil {
		return err
	}

	return nil
}
