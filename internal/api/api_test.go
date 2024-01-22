package api_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/niluan304/leetcode/internal/api"
)

func TestClient_ConsolePanelConfig(t *testing.T) {
	var ctx = context.Background()

	type client struct {
		cookie   map[string]string
		endpoint string
	}
	type args struct {
		ctx       context.Context
		titleSlug string
	}
	tests := []struct {
		name    string
		fields  client
		args    args
		wantRes *api.ConsolePanelConfigRes
		wantErr bool
	}{
		{
			name: "EndpointZh",
			fields: client{
				cookie: map[string]string{
					//"_bl_uid":                     "qdl04jjhfzwt10nmIk8dd9yb1Oya",
					//"csrftoken":                   "D7wEbUi2MVGjyCFE8Eg68JjS8CnD0FsSAD9c6D6TYXP89NCGLlpf3GoCnFnOLD12",
					//"NEW_QUESTION_DETAIL_PAGE_V2": "1",
					//"NEW_PROBLEMLIST_PAGE":        "1",
					//"LEETCODE_SESSION":            "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJfYXV0aF91c2VyX2lkIjoiMjUyNDM0NCIsIl9hdXRoX3VzZXJfYmFja2VuZCI6ImF1dGhlbnRpY2F0aW9uLmF1dGhfYmFja2VuZHMuUGhvbmVBdXRoZW50aWNhdGlvbkJhY2tlbmQiLCJfYXV0aF91c2VyX2hhc2giOiI3MTEwYWRjMDk4YmNhYmJlYTViZWQwM2YyMDUxOWY2ZDM0OTcwOWM2OGIyZGE4YjZhZDRkZGM0MWU5ZjQ2Mjc2IiwiaWQiOjI1MjQzNDQsImVtYWlsIjoibmlsdWFuMzA0QGdtYWlsLmNvbSIsInVzZXJuYW1lIjoid3UtZmVuZy1odyIsInVzZXJfc2x1ZyI6Ind1LWZlbmctaHciLCJhdmF0YXIiOiJodHRwczovL2Fzc2V0cy5sZWV0Y29kZS5jbi9hbGl5dW4tbGMtdXBsb2FkL3VzZXJzL3d1LWZlbmctaHcvYXZhdGFyXzE2MDY5NzkxMDgucG5nIiwicGhvbmVfdmVyaWZpZWQiOnRydWUsIl90aW1lc3RhbXAiOjE2ODgyMTA1NTEuMTA4MTU3NCwiZXhwaXJlZF90aW1lXyI6MTY5MDc0MzYwMCwidmVyc2lvbl9rZXlfIjoxLCJsYXRlc3RfdGltZXN0YW1wXyI6MTY4ODIzMzAwNX0.iXbgRZbqfvjaycx--ImeMuwxA1bGzB1ElTbfTCmPvo0",
				},
				endpoint: api.EndpointZh,
			},
			args: args{
				ctx:       ctx,
				titleSlug: "fHi6rV",
			},
			wantRes: nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := api.New(tt.fields.endpoint, api.WithCookie(tt.fields.cookie))

			gotRes, err := c.ConsolePanelConfig(tt.args.ctx, tt.args.titleSlug)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConsolePanelConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				//t.Errorf("ConsolePanelConfig() gotRes = %v, want %v", gotRes, tt.wantRes)
			}

			fmt.Println("gotRes.Question.QuestionId:", gotRes.Question.QuestionId)
			fmt.Println("gotRes.Question.QuestionFrontendId:", gotRes.Question.QuestionFrontendId)
			fmt.Println("gotRes.Question.QuestionTitle:", gotRes.Question.QuestionTitle)
			fmt.Println("gotRes.Question.EnableRunCode:", gotRes.Question.EnableRunCode)
			fmt.Println("gotRes.Question.EnableSubmit:", gotRes.Question.EnableSubmit)
			fmt.Println("gotRes.Question.EnableTestMode:", gotRes.Question.EnableTestMode)
			fmt.Println("gotRes.Question.JsonExampleTestcases:", gotRes.Question.JsonExampleTestcases)
			fmt.Println("gotRes.Question.ExampleTestcases:", gotRes.Question.ExampleTestcases)
			fmt.Println("gotRes.Question.MetaData:", gotRes.Question.MetaData)
			fmt.Println("gotRes.Question.SampleTestCase:", gotRes.Question.SampleTestCase)
		})
	}
}

func TestClient_ProblemsQuestionList(t *testing.T) {
	var ctx = context.Background()

	type client struct {
		cookie   map[string]string
		endpoint string
	}

	type fields client
	type args struct {
		ctx context.Context
		in  api.ProblemQuestionReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *api.ProblemQuestionRes
		wantErr bool
	}{
		{
			name: "EndpointZh",
			fields: fields{
				cookie:   nil,
				endpoint: api.EndpointZh,
			},
			args: args{
				ctx: ctx,
				in: api.ProblemQuestionReq{
					CategorySlug: "",
					Skip:         0,
					Limit:        20,
					Filters: api.ProblemQuestionFilters{
						Difficulty:     "",
						Status:         "",
						SearchKeywords: "188",
						Tags:           nil,
						ListId:         "",
					},
				},
			},
			wantRes: nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := api.New(tt.fields.endpoint, api.WithCookie(tt.fields.cookie))

			gotRes, err := c.ProblemsetQuestionList(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProblemsetQuestionList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("ProblemsetQuestionList() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestClient_QaQuestionDetail(t *testing.T) {
	var ctx = context.Background()

	type client struct {
		cookie   map[string]string
		endpoint string
	}

	type fields client
	type args struct {
		ctx context.Context
		in  api.QaQuestionDetailReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *api.QaQuestionDetailRes
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				cookie:   nil,
				endpoint: api.EndpointZh,
			},
			args: args{
				ctx: ctx,
				in: api.QaQuestionDetailReq{
					Uuid: "G0n5iY",
				},
			},
			wantRes: nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := api.New(tt.fields.endpoint, api.WithCookie(tt.fields.cookie))
			gotRes, err := c.QaQuestionDetail(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("QaQuestionDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			fmt.Println(gotRes.QaQuestion.Content)
		})
	}
}

func TestClient_QuestionData(t *testing.T) {
	var ctx = context.Background()

	type client struct {
		cookie   map[string]string
		endpoint string
	}
	type args struct {
		ctx       context.Context
		titleSlug string
	}
	tests := []struct {
		name    string
		fields  client
		args    args
		wantRes *api.ConsolePanelConfigRes
		wantErr bool
	}{
		{
			name: "EndpointZh",
			fields: client{
				cookie: map[string]string{
					//"_bl_uid":                     "qdl04jjhfzwt10nmIk8dd9yb1Oya",
					//"csrftoken":                   "D7wEbUi2MVGjyCFE8Eg68JjS8CnD0FsSAD9c6D6TYXP89NCGLlpf3GoCnFnOLD12",
					//"NEW_QUESTION_DETAIL_PAGE_V2": "1",
					//"NEW_PROBLEMLIST_PAGE":        "1",
					//"LEETCODE_SESSION":            "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJfYXV0aF91c2VyX2lkIjoiMjUyNDM0NCIsIl9hdXRoX3VzZXJfYmFja2VuZCI6ImF1dGhlbnRpY2F0aW9uLmF1dGhfYmFja2VuZHMuUGhvbmVBdXRoZW50aWNhdGlvbkJhY2tlbmQiLCJfYXV0aF91c2VyX2hhc2giOiI3MTEwYWRjMDk4YmNhYmJlYTViZWQwM2YyMDUxOWY2ZDM0OTcwOWM2OGIyZGE4YjZhZDRkZGM0MWU5ZjQ2Mjc2IiwiaWQiOjI1MjQzNDQsImVtYWlsIjoibmlsdWFuMzA0QGdtYWlsLmNvbSIsInVzZXJuYW1lIjoid3UtZmVuZy1odyIsInVzZXJfc2x1ZyI6Ind1LWZlbmctaHciLCJhdmF0YXIiOiJodHRwczovL2Fzc2V0cy5sZWV0Y29kZS5jbi9hbGl5dW4tbGMtdXBsb2FkL3VzZXJzL3d1LWZlbmctaHcvYXZhdGFyXzE2MDY5NzkxMDgucG5nIiwicGhvbmVfdmVyaWZpZWQiOnRydWUsIl90aW1lc3RhbXAiOjE2ODgyMTA1NTEuMTA4MTU3NCwiZXhwaXJlZF90aW1lXyI6MTY5MDc0MzYwMCwidmVyc2lvbl9rZXlfIjoxLCJsYXRlc3RfdGltZXN0YW1wXyI6MTY4ODIzMzAwNX0.iXbgRZbqfvjaycx--ImeMuwxA1bGzB1ElTbfTCmPvo0",
				},
				endpoint: api.EndpointZh,
			},
			args: args{
				ctx:       ctx,
				titleSlug: "movement-of-robots",
			},
			wantRes: nil,
			wantErr: false,
		},
		{
			name: "EndpointZh",
			fields: client{
				cookie: map[string]string{
					//"_bl_uid":                     "qdl04jjhfzwt10nmIk8dd9yb1Oya",
					//"csrftoken":                   "D7wEbUi2MVGjyCFE8Eg68JjS8CnD0FsSAD9c6D6TYXP89NCGLlpf3GoCnFnOLD12",
					//"NEW_QUESTION_DETAIL_PAGE_V2": "1",
					//"NEW_PROBLEMLIST_PAGE":        "1",
					//"LEETCODE_SESSION":            "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJfYXV0aF91c2VyX2lkIjoiMjUyNDM0NCIsIl9hdXRoX3VzZXJfYmFja2VuZCI6ImF1dGhlbnRpY2F0aW9uLmF1dGhfYmFja2VuZHMuUGhvbmVBdXRoZW50aWNhdGlvbkJhY2tlbmQiLCJfYXV0aF91c2VyX2hhc2giOiI3MTEwYWRjMDk4YmNhYmJlYTViZWQwM2YyMDUxOWY2ZDM0OTcwOWM2OGIyZGE4YjZhZDRkZGM0MWU5ZjQ2Mjc2IiwiaWQiOjI1MjQzNDQsImVtYWlsIjoibmlsdWFuMzA0QGdtYWlsLmNvbSIsInVzZXJuYW1lIjoid3UtZmVuZy1odyIsInVzZXJfc2x1ZyI6Ind1LWZlbmctaHciLCJhdmF0YXIiOiJodHRwczovL2Fzc2V0cy5sZWV0Y29kZS5jbi9hbGl5dW4tbGMtdXBsb2FkL3VzZXJzL3d1LWZlbmctaHcvYXZhdGFyXzE2MDY5NzkxMDgucG5nIiwicGhvbmVfdmVyaWZpZWQiOnRydWUsIl90aW1lc3RhbXAiOjE2ODgyMTA1NTEuMTA4MTU3NCwiZXhwaXJlZF90aW1lXyI6MTY5MDc0MzYwMCwidmVyc2lvbl9rZXlfIjoxLCJsYXRlc3RfdGltZXN0YW1wXyI6MTY4ODIzMzAwNX0.iXbgRZbqfvjaycx--ImeMuwxA1bGzB1ElTbfTCmPvo0",
				},
				endpoint: api.EndpointZh,
			},
			args: args{
				ctx:       ctx,
				titleSlug: "apply-operations-to-make-two-strings-equal",
			},
			wantRes: nil,
			wantErr: false,
		},
		{
			name: "EndpointZh",
			fields: client{
				cookie: map[string]string{
					//"_bl_uid":                     "qdl04jjhfzwt10nmIk8dd9yb1Oya",
					//"csrftoken":                   "D7wEbUi2MVGjyCFE8Eg68JjS8CnD0FsSAD9c6D6TYXP89NCGLlpf3GoCnFnOLD12",
					//"NEW_QUESTION_DETAIL_PAGE_V2": "1",
					//"NEW_PROBLEMLIST_PAGE":        "1",
					//"LEETCODE_SESSION":            "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJfYXV0aF91c2VyX2lkIjoiMjUyNDM0NCIsIl9hdXRoX3VzZXJfYmFja2VuZCI6ImF1dGhlbnRpY2F0aW9uLmF1dGhfYmFja2VuZHMuUGhvbmVBdXRoZW50aWNhdGlvbkJhY2tlbmQiLCJfYXV0aF91c2VyX2hhc2giOiI3MTEwYWRjMDk4YmNhYmJlYTViZWQwM2YyMDUxOWY2ZDM0OTcwOWM2OGIyZGE4YjZhZDRkZGM0MWU5ZjQ2Mjc2IiwiaWQiOjI1MjQzNDQsImVtYWlsIjoibmlsdWFuMzA0QGdtYWlsLmNvbSIsInVzZXJuYW1lIjoid3UtZmVuZy1odyIsInVzZXJfc2x1ZyI6Ind1LWZlbmctaHciLCJhdmF0YXIiOiJodHRwczovL2Fzc2V0cy5sZWV0Y29kZS5jbi9hbGl5dW4tbGMtdXBsb2FkL3VzZXJzL3d1LWZlbmctaHcvYXZhdGFyXzE2MDY5NzkxMDgucG5nIiwicGhvbmVfdmVyaWZpZWQiOnRydWUsIl90aW1lc3RhbXAiOjE2ODgyMTA1NTEuMTA4MTU3NCwiZXhwaXJlZF90aW1lXyI6MTY5MDc0MzYwMCwidmVyc2lvbl9rZXlfIjoxLCJsYXRlc3RfdGltZXN0YW1wXyI6MTY4ODIzMzAwNX0.iXbgRZbqfvjaycx--ImeMuwxA1bGzB1ElTbfTCmPvo0",
				},
				endpoint: api.EndpointZh,
			},
			args: args{
				ctx:       ctx,
				titleSlug: "stock-price-fluctuation", // 系统设计
			},
			wantRes: nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := api.New(tt.fields.endpoint, api.WithCookie(tt.fields.cookie))

			gotRes, err := c.QuestionData(tt.args.ctx, api.QuestionDataReq{TitleSlug: tt.args.titleSlug})
			if (err != nil) != tt.wantErr {
				t.Errorf("ConsolePanelConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				//t.Errorf("ConsolePanelConfig() gotRes = %+v, want %+v", gotRes, tt.wantRes)
			}

			data, err := gotRes.Question.ReUnmarshal()
			if err != nil {
				t.Errorf("ConsolePanelConfig() error = %v", err)
				return
			}
			fmt.Println(data)
		})
	}
}

func TestClient_QuestionOfToday(t *testing.T) {
	var ctx = context.Background()

	type client struct {
		cookie   map[string]string
		endpoint string
	}
	type fields client

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *api.QuestionOfTodayRes
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				cookie: map[string]string{
					//"_bl_uid":                     "qdl04jjhfzwt10nmIk8dd9yb1Oya",
					//"csrftoken":                   "D7wEbUi2MVGjyCFE8Eg68JjS8CnD0FsSAD9c6D6TYXP89NCGLlpf3GoCnFnOLD12",
					//"NEW_QUESTION_DETAIL_PAGE_V2": "1",
					//"NEW_PROBLEMLIST_PAGE":        "1",
					//"LEETCODE_SESSION":            "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJfYXV0aF91c2VyX2lkIjoiMjUyNDM0NCIsIl9hdXRoX3VzZXJfYmFja2VuZCI6ImF1dGhlbnRpY2F0aW9uLmF1dGhfYmFja2VuZHMuUGhvbmVBdXRoZW50aWNhdGlvbkJhY2tlbmQiLCJfYXV0aF91c2VyX2hhc2giOiI3MTEwYWRjMDk4YmNhYmJlYTViZWQwM2YyMDUxOWY2ZDM0OTcwOWM2OGIyZGE4YjZhZDRkZGM0MWU5ZjQ2Mjc2IiwiaWQiOjI1MjQzNDQsImVtYWlsIjoibmlsdWFuMzA0QGdtYWlsLmNvbSIsInVzZXJuYW1lIjoid3UtZmVuZy1odyIsInVzZXJfc2x1ZyI6Ind1LWZlbmctaHciLCJhdmF0YXIiOiJodHRwczovL2Fzc2V0cy5sZWV0Y29kZS5jbi9hbGl5dW4tbGMtdXBsb2FkL3VzZXJzL3d1LWZlbmctaHcvYXZhdGFyXzE2MDY5NzkxMDgucG5nIiwicGhvbmVfdmVyaWZpZWQiOnRydWUsIl90aW1lc3RhbXAiOjE2ODgyMTA1NTEuMTA4MTU3NCwiZXhwaXJlZF90aW1lXyI6MTY5MDc0MzYwMCwidmVyc2lvbl9rZXlfIjoxLCJsYXRlc3RfdGltZXN0YW1wXyI6MTY4ODIzMzAwNX0.iXbgRZbqfvjaycx--ImeMuwxA1bGzB1ElTbfTCmPvo0",
				},
				endpoint: api.EndpointZh,
			},
			args: args{
				ctx: ctx,
			},
			wantRes: &api.QuestionOfTodayRes{
				TodayRecord: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := api.New(tt.fields.endpoint, api.WithCookie(tt.fields.cookie))

			gotRes, err := c.QuestionOfToday(tt.args.ctx, api.QuestionOfTodayReq{})
			if (err != nil) != tt.wantErr {
				t.Errorf("QuestionOfToday() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("QuestionOfToday() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestClient_SolutionArticle(t *testing.T) {
	var ctx = context.Background()

	type client struct {
		cookie   map[string]string
		endpoint string
	}
	type fields client

	type args struct {
		ctx context.Context
		in  api.SolutionArticleReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *api.SolutionArticleRes
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				cookie: map[string]string{
					//"_bl_uid":                     "qdl04jjhfzwt10nmIk8dd9yb1Oya",
					//"csrftoken":                   "D7wEbUi2MVGjyCFE8Eg68JjS8CnD0FsSAD9c6D6TYXP89NCGLlpf3GoCnFnOLD12",
					//"NEW_QUESTION_DETAIL_PAGE_V2": "1",
					//"NEW_PROBLEMLIST_PAGE":        "1",
					//"LEETCODE_SESSION":            "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJfYXV0aF91c2VyX2lkIjoiMjUyNDM0NCIsIl9hdXRoX3VzZXJfYmFja2VuZCI6ImF1dGhlbnRpY2F0aW9uLmF1dGhfYmFja2VuZHMuUGhvbmVBdXRoZW50aWNhdGlvbkJhY2tlbmQiLCJfYXV0aF91c2VyX2hhc2giOiI3MTEwYWRjMDk4YmNhYmJlYTViZWQwM2YyMDUxOWY2ZDM0OTcwOWM2OGIyZGE4YjZhZDRkZGM0MWU5ZjQ2Mjc2IiwiaWQiOjI1MjQzNDQsImVtYWlsIjoibmlsdWFuMzA0QGdtYWlsLmNvbSIsInVzZXJuYW1lIjoid3UtZmVuZy1odyIsInVzZXJfc2x1ZyI6Ind1LWZlbmctaHciLCJhdmF0YXIiOiJodHRwczovL2Fzc2V0cy5sZWV0Y29kZS5jbi9hbGl5dW4tbGMtdXBsb2FkL3VzZXJzL3d1LWZlbmctaHcvYXZhdGFyXzE2MDY5NzkxMDgucG5nIiwicGhvbmVfdmVyaWZpZWQiOnRydWUsIl90aW1lc3RhbXAiOjE2ODgyMTA1NTEuMTA4MTU3NCwiZXhwaXJlZF90aW1lXyI6MTY5MDc0MzYwMCwidmVyc2lvbl9rZXlfIjoxLCJsYXRlc3RfdGltZXN0YW1wXyI6MTY4ODIzMzAwNX0.iXbgRZbqfvjaycx--ImeMuwxA1bGzB1ElTbfTCmPvo0",
				},
				endpoint: api.EndpointZh,
			},
			args: args{
				ctx: ctx,
				in: api.SolutionArticleReq{
					Slug: "wu-nao-zuo-fa-er-wei-qian-zhui-he-er-wei-zwiu",
				},
			},
			wantRes: nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := api.New(tt.fields.endpoint, api.WithCookie(tt.fields.cookie))

			gotRes, err := c.SolutionArticle(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("SolutionArticle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				//t.Errorf("SolutionArticle() gotRes = %v, want %v", gotRes, tt.wantRes)
			}

			//fmt.Println(gotRes.SolutionArticle.Content)

			const Flag = "```"
			content := gotRes.SolutionArticle.Content
			//for idx := strings.Index(content, Flag); idx >= 0; {
			//	idx2 := strings.Index(content[idx+len(Flag):], Flag+"\n")
			//	if idx2 == -1 {
			//		break
			//	}
			//
			//	prefix := content[:idx-1]
			//	suffix := content[idx+9+idx2:]
			//	content = prefix + suffix
			//}

			fmt.Println(content)
		})
	}
}
