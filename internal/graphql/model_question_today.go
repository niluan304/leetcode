package graphql

// QuestionOfTodayRes 每日一题
type QuestionOfTodayRes struct {
	TodayRecord []struct {
		Date       string `json:"date"`
		UserStatus string `json:"userStatus"`
		Question   struct {
			QuestionId         string      `json:"questionId"`
			FrontendQuestionId string      `json:"frontendQuestionId"`
			Difficulty         string      `json:"difficulty"`
			Title              string      `json:"title"`
			TitleCn            string      `json:"titleCn"`
			TitleSlug          string      `json:"titleSlug"`
			PaidOnly           bool        `json:"paidOnly"`
			FreqBar            interface{} `json:"freqBar"`
			IsFavor            bool        `json:"isFavor"`
			AcRate             float64     `json:"acRate"`
			Status             interface{} `json:"status"`
			SolutionNum        int         `json:"solutionNum"`
			HasVideoSolution   bool        `json:"hasVideoSolution"`
			TopicTags          []struct {
				Name           string `json:"name"`
				NameTranslated string `json:"nameTranslated"`
				Id             string `json:"id"`
			} `json:"topicTags"`
			Extra struct {
				TopCompanyTags []interface{} `json:"topCompanyTags"`
			} `json:"extra"`
		} `json:"question"`
		LastSubmission interface{} `json:"lastSubmission"`
	} `json:"todayRecord"`
}
