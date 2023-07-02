package graphql

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
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
		wantRes *ConsolePanelConfigRes
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
				endpoint: EndpointZh,
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
			c := New(tt.fields.endpoint, WithCookie(tt.fields.cookie))

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
		wantRes *ConsolePanelConfigRes
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
				endpoint: EndpointZh,
			},
			args: args{
				ctx:       ctx,
				titleSlug: "two-sum",
			},
			wantRes: nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(tt.fields.endpoint, WithCookie(tt.fields.cookie))

			gotRes, err := c.QuestionData(tt.args.ctx, tt.args.titleSlug)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConsolePanelConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				//t.Errorf("ConsolePanelConfig() gotRes = %+v, want %+v", gotRes, tt.wantRes)
			}

			data, err := json.Marshal(gotRes.Question)
			if err != nil {
				t.Errorf("ConsolePanelConfig() error = %v", err)
				return
			}
			fmt.Println(string(data))
		})
	}
}
