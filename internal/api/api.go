package api

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

func (c *Client) ConsolePanelConfig(ctx context.Context, titleSlug string) (res *ConsolePanelConfigRes, err error) {
	req := graphql.NewRequest(`
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
`)
	req.Var("titleSlug", titleSlug)
	req.Var("operationName", "consolePanelConfig")

	path := fmt.Sprintf("/problems/%s/", titleSlug)
	err = c.request(ctx, req, path, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// ProblemsetQuestionList 根据搜索条件查看题库列表
func (c *Client) ProblemsetQuestionList(ctx context.Context, in ProblemQuestionReq) (res *ProblemQuestionRes, err error) {
	req := graphql.NewRequest(`query problemsetQuestionList($categorySlug: String, $limit: Int, $skip: Int, $filters: QuestionListFilterInput) {
  problemsetQuestionList(
    categorySlug: $categorySlug
    limit: $limit
    skip: $skip
    filters: $filters
  ) {
    hasMore
    total
    questions {
      acRate
      difficulty
      freqBar
      frontendQuestionId
      isFavor
      paidOnly
      solutionNum
      status
      title
      titleCn
      titleSlug
      topicTags {
        name
        nameTranslated
        id
        slug
      }
      extra {
        hasVideoSolution
        topCompanyTags {
          imgUrl
          slug
          numSubscribed
        }
      }
    }
  }
}
`)

	var values = map[string]any{
		"categorySlug":  in.CategorySlug,
		"skip":          in.Skip,
		"limit":         in.Limit,
		"filters":       in.Filters,
		"operationName": "problemsetQuestionList",
	}
	for key, value := range values {
		req.Var(key, value)
	}

	referer := fmt.Sprintf("/problemset/all/?page=1&search=%s", in.Filters.SearchKeywords)
	referer = ""
	err = c.request(ctx, req, referer, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// QaQuestionDetail 查看讨论内容
func (c *Client) QaQuestionDetail(ctx context.Context, in QaQuestionDetailReq) (res *QaQuestionDetailRes, err error) {
	req := graphql.NewRequest(`
query qaQuestionDetail($uuid: ID!) {
  qaQuestion(uuid: $uuid) {
    ...qaQuestion
    myAnswerId
    __typename
  }
}

fragment qaQuestion on QAQuestionNode {
  ipRegion
  uuid
  slug
  title
  thumbnail
  summary
  content
  slateValue
  sunk
  pinned
  pinnedGlobally
  byLeetcode
  isRecommended
  isRecommendedGlobally
  subscribed
  hitCount
  numAnswers
  numPeopleInvolved
  numSubscribed
  createdAt
  updatedAt
  status
  identifier
  resourceType
  articleType
  alwaysShow
  alwaysExpand
  score
  favoriteCount
  isMyFavorite
  isAnonymous
  canEdit
  reactionType
  atQuestionTitleSlug
  blockComments
  reactionsV2 {
    count
    reactionType
    __typename
  }
  tags {
    name
    nameTranslated
    slug
    imgUrl
    tagType
    __typename
  }
  subject {
    slug
    title
    __typename
  }
  contentAuthor {
    ...contentAuthor
    __typename
  }
  realAuthor {
    ...realAuthor
    __typename
  }
  __typename
}

fragment contentAuthor on ArticleAuthor {
  username
  userSlug
  realName
  avatar
  __typename
}

fragment realAuthor on UserNode {
  username
  profile {
    userSlug
    realName
    userAvatar
    __typename
  }
  __typename
}
    `)

	var values = map[string]any{
		"uuid": in.Uuid,
	}
	for key, value := range values {
		req.Var(key, value)
	}

	referer := fmt.Sprintf("/circle/discuss/" + in.Uuid)
	err = c.request(ctx, req, referer, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// QuestionData 返回题目内容和模板代码
func (c *Client) QuestionData(ctx context.Context, in QuestionDataReq) (res *QuestionDataRes, err error) {
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
	req.Var("titleSlug", in.TitleSlug)
	req.Var("operationName", "questionData")

	path := fmt.Sprintf("/problems/%s/", in.TitleSlug)
	err = c.request(ctx, req, path, &res)
	if err != nil {
		return nil, err
	}

	//{
	//	err := c.save("questionData", titleSlug, res)
	//	if err != nil {
	//		log.Println("fail save file:", titleSlug, err)
	//	}
	//}

	return res, nil
}

// QuestionOfToday 查看今日的每日一题信息
func (c *Client) QuestionOfToday(ctx context.Context, in QuestionOfTodayReq) (res *QuestionOfTodayRes, err error) {
	req := graphql.NewRequest(`
    query questionOfToday {
  todayRecord {
    date
    userStatus
    question {
      questionId
      frontendQuestionId: questionFrontendId
      difficulty
      title
      titleCn: translatedTitle
      titleSlug
      paidOnly: isPaidOnly
      freqBar
      isFavor
      acRate
      status
      solutionNum
      hasVideoSolution
      topicTags {
        name
        nameTranslated: translatedName
        id
      }
      extra {
        topCompanyTags {
          imgUrl
          slug
          numSubscribed
        }
      }
    }
    lastSubmission {
      id
    }
  }
}
    `)
	err = c.request(ctx, req, "https://leetcode.cn/problems/stock-price-fluctuation/description", &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// SolutionArticle 查看题解文章
func (c *Client) SolutionArticle(ctx context.Context, in SolutionArticleReq) (res *SolutionArticleRes, err error) {
	req := graphql.NewRequest(`
    query discussTopic($slug: String) {
  solutionArticle(slug: $slug, orderBy: DEFAULT) {
    ...solutionArticle
    content
    next {
      slug
      title
    }
    prev {
      slug
      title
    }
  }
}
    
    fragment solutionArticle on SolutionArticleNode {
  ipRegion
  rewardEnabled
  canEditReward
  uuid
  title
  slug
  sunk
  chargeType
  status
  identifier
  canEdit
  canSee
  reactionType
  reactionsV2 {
    count
    reactionType
  }
  tags {
    name
    nameTranslated
    slug
    tagType
  }
  createdAt
  thumbnail
  author {
    username
    isDiscussAdmin
    isDiscussStaff
    profile {
      userAvatar
      userSlug
      realName
      reputation
    }
  }
  summary
  topic {
    id
    subscribed
    commentCount
    viewCount
    post {
      id
      status
      voteStatus
      isOwnPost
    }
  }
  byLeetcode
  isMyFavorite
  isMostPopular
  favoriteCount
  isEditorsPick
  hitCount
  videosInfo {
    videoId
    coverUrl
    duration
  }
}
    `)

	var values = map[string]any{
		"slug":          in.Slug,
		"operationName": "discussTopic",
	}
	for key, value := range values {
		req.Var(key, value)
	}

	referer := fmt.Sprintf("/problemset/all")
	referer = ""
	err = c.request(ctx, req, referer, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
