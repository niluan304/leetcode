package graphql

type (
	ProblemQuestionReq struct {
		CategorySlug string                 `json:"categorySlug"`
		Skip         int                    `json:"skip"`
		Limit        int                    `json:"limit"`
		Filters      ProblemQuestionFilters `json:"filters"`
	}

	ProblemQuestionFilters struct {
		Difficulty     string   `json:"difficulty,omitempty"`     //
		Status         string   `json:"status,omitempty"`         //
		SearchKeywords string   `json:"searchKeywords,omitempty"` //
		Tags           []string `json:"tags,omitempty"`           //
		ListId         string   `json:"listId,omitempty"`         //
	}

	ProblemQuestionRes struct {
		ProblemsetQuestionList struct {
			Typename  string               `json:"__typename"`
			Questions []ProblemsetQuestion `json:"questions"`
			HasMore   bool                 `json:"hasMore"`
			Total     int                  `json:"total"`
		} `json:"problemsetQuestionList"`
	}

	ProblemsetQuestion struct {
		Typename           string               `json:"__typename"`
		AcRate             float64              `json:"acRate"`
		Difficulty         string               `json:"difficulty"`
		FreqBar            int                  `json:"freqBar"`
		PaidOnly           bool                 `json:"paidOnly"`
		Status             string               `json:"status"`
		FrontendQuestionId string               `json:"frontendQuestionId"`
		IsFavor            bool                 `json:"isFavor"`
		SolutionNum        int                  `json:"solutionNum"`
		Title              string               `json:"title"`
		TitleCn            string               `json:"titleCn"`
		TitleSlug          string               `json:"titleSlug"`
		TopicTags          []ProblemsetTopicTag `json:"topicTags"`
		Extra              ProblemsetExtra      `json:"extra"`
	}

	ProblemsetTopicTag struct {
		Id             string `json:"id"`
		Name           string `json:"name"`
		Slug           string `json:"slug"`
		NameTranslated string `json:"nameTranslated"`
		Typename       string `json:"__typename"`
	}

	ProblemsetExtra struct {
		CompanyTagNum    int  `json:"companyTagNum"`
		HasVideoSolution bool `json:"hasVideoSolution"`
		TopCompanyTags   []struct {
			ImgUrl   string `json:"imgUrl"`
			Slug     string `json:"slug"`
			Typename string `json:"__typename"`
		} `json:"topCompanyTags"`
		Typename string `json:"__typename"`
	}
)
