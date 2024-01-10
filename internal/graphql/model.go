package graphql

import "time"

type (
	ConsolePanelConfigRes struct {
		Question ConsolePanelConfigQuestion `json:"question"`
	}
	ConsolePanelConfigQuestion struct {
		QuestionId           string `json:"questionId"`
		QuestionFrontendId   string `json:"questionFrontendId"`
		QuestionTitle        string `json:"questionTitle"`
		EnableRunCode        bool   `json:"enableRunCode"`
		EnableSubmit         bool   `json:"enableSubmit"`
		EnableTestMode       bool   `json:"enableTestMode"`
		JsonExampleTestcases string `json:"jsonExampleTestcases"`
		ExampleTestcases     string `json:"exampleTestcases"`
		MetaData             string `json:"metaData"`
		SampleTestCase       string `json:"sampleTestCase"`
	}
)

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

type (
	QaQuestionDetailReq struct {
		Uuid string
	}

	QaQuestionDetailRes struct {
		QaQuestion QaQuestionDetail `json:"qaQuestion"`
	}

	QaQuestionDetail struct {
		IpRegion              string      `json:"ipRegion"`
		Uuid                  string      `json:"uuid"`
		Slug                  string      `json:"slug"`
		Title                 string      `json:"title"`
		Thumbnail             string      `json:"thumbnail"`
		Summary               string      `json:"summary"`
		Content               string      `json:"content"`
		SlateValue            string      `json:"slateValue"`
		Sunk                  bool        `json:"sunk"`
		Pinned                bool        `json:"pinned"`
		PinnedGlobally        bool        `json:"pinnedGlobally"`
		ByLeetcode            bool        `json:"byLeetcode"`
		IsRecommended         bool        `json:"isRecommended"`
		IsRecommendedGlobally bool        `json:"isRecommendedGlobally"`
		Subscribed            bool        `json:"subscribed"`
		HitCount              int         `json:"hitCount"`
		NumAnswers            int         `json:"numAnswers"`
		NumPeopleInvolved     int         `json:"numPeopleInvolved"`
		NumSubscribed         int         `json:"numSubscribed"`
		CreatedAt             time.Time   `json:"createdAt"`
		UpdatedAt             time.Time   `json:"updatedAt"`
		Status                string      `json:"status"`
		Identifier            string      `json:"identifier"`
		ResourceType          string      `json:"resourceType"`
		ArticleType           string      `json:"articleType"`
		AlwaysShow            bool        `json:"alwaysShow"`
		AlwaysExpand          bool        `json:"alwaysExpand"`
		Score                 interface{} `json:"score"`
		FavoriteCount         int         `json:"favoriteCount"`
		IsMyFavorite          bool        `json:"isMyFavorite"`
		IsAnonymous           bool        `json:"isAnonymous"`
		CanEdit               bool        `json:"canEdit"`
		ReactionType          interface{} `json:"reactionType"`
		AtQuestionTitleSlug   string      `json:"atQuestionTitleSlug"`
		BlockComments         bool        `json:"blockComments"`
		ReactionsV2           []struct {
			Count        int    `json:"count"`
			ReactionType string `json:"reactionType"`
			Typename     string `json:"__typename"`
		} `json:"reactionsV2"`
		Tags []struct {
			Name           string      `json:"name"`
			NameTranslated string      `json:"nameTranslated"`
			Slug           string      `json:"slug"`
			ImgUrl         interface{} `json:"imgUrl"`
			TagType        string      `json:"tagType"`
			Typename       string      `json:"__typename"`
		} `json:"tags"`
		Subject struct {
			Slug     string `json:"slug"`
			Title    string `json:"title"`
			Typename string `json:"__typename"`
		} `json:"subject"`
		ContentAuthor struct {
			Username string `json:"username"`
			UserSlug string `json:"userSlug"`
			RealName string `json:"realName"`
			Avatar   string `json:"avatar"`
			Typename string `json:"__typename"`
		} `json:"contentAuthor"`
		RealAuthor interface{} `json:"realAuthor"`
		Typename   string      `json:"__typename"`
		MyAnswerId interface{} `json:"myAnswerId"`
	}
)

type (
	QuestionOfTodayReq struct {
	}

	// QuestionOfTodayRes 每日一题
	QuestionOfTodayRes struct {
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
)

type (
	SolutionArticleReq struct {
		Slug string
	}

	SolutionArticleRes struct {
		SolutionArticle struct {
			IpRegion      string      `json:"ipRegion"`
			RewardEnabled bool        `json:"rewardEnabled"`
			CanEditReward bool        `json:"canEditReward"`
			Uuid          string      `json:"uuid"`
			Title         string      `json:"title"`
			Slug          string      `json:"slug"`
			Sunk          bool        `json:"sunk"`
			ChargeType    string      `json:"chargeType"`
			Status        string      `json:"status"`
			Identifier    string      `json:"identifier"`
			CanEdit       bool        `json:"canEdit"`
			CanSee        bool        `json:"canSee"`
			ReactionType  interface{} `json:"reactionType"`
			ReactionsV2   []struct {
				Count        int    `json:"count"`
				ReactionType string `json:"reactionType"`
			} `json:"reactionsV2"`
			Tags []struct {
				Name           string `json:"name"`
				NameTranslated string `json:"nameTranslated"`
				Slug           string `json:"slug"`
				TagType        string `json:"tagType"`
			} `json:"tags"`
			CreatedAt time.Time `json:"createdAt"`
			Thumbnail string    `json:"thumbnail"`
			Author    struct {
				Username       string `json:"username"`
				IsDiscussAdmin bool   `json:"isDiscussAdmin"`
				IsDiscussStaff bool   `json:"isDiscussStaff"`
				Profile        struct {
					UserAvatar string `json:"userAvatar"`
					UserSlug   string `json:"userSlug"`
					RealName   string `json:"realName"`
					Reputation int    `json:"reputation"`
				} `json:"profile"`
			} `json:"author"`
			Summary string `json:"summary"`
			Topic   struct {
				Id           int  `json:"id"`
				Subscribed   bool `json:"subscribed"`
				CommentCount int  `json:"commentCount"`
				ViewCount    int  `json:"viewCount"`
				Post         struct {
					Id         int    `json:"id"`
					Status     string `json:"status"`
					VoteStatus int    `json:"voteStatus"`
					IsOwnPost  bool   `json:"isOwnPost"`
				} `json:"post"`
			} `json:"topic"`
			ByLeetcode    bool `json:"byLeetcode"`
			IsMyFavorite  bool `json:"isMyFavorite"`
			IsMostPopular bool `json:"isMostPopular"`
			FavoriteCount int  `json:"favoriteCount"`
			IsEditorsPick bool `json:"isEditorsPick"`
			HitCount      int  `json:"hitCount"`
			VideosInfo    []struct {
				VideoId  string  `json:"videoId"`
				CoverUrl string  `json:"coverUrl"`
				Duration float64 `json:"duration"`
			} `json:"videosInfo"`
			Content string `json:"content"`
			Next    struct {
				Slug  string `json:"slug"`
				Title string `json:"title"`
			} `json:"next"`
			Prev struct {
				Slug  string `json:"slug"`
				Title string `json:"title"`
			} `json:"prev"`
		} `json:"solutionArticle"`
	}
)
