package graphql

import "time"

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
