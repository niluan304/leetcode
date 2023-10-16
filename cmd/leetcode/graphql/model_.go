package graphql

import "time"

type SolutionArticleReq struct {
	Slug string
}

type SolutionArticleRes struct {
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
