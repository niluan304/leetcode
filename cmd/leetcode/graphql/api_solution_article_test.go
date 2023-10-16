package graphql

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestClient_SolutionArticle(t *testing.T) {
	var ctx = context.Background()

	type client struct {
		cookie   map[string]string
		endpoint string
	}
	type fields client

	type args struct {
		ctx context.Context
		in  SolutionArticleReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *SolutionArticleRes
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
				endpoint: EndpointZh,
			},
			args: args{
				ctx: ctx,
				in: SolutionArticleReq{
					Slug: "shu-zu-zhong-de-di-kge-zui-da-yuan-su-by-leetcode-",
				},
			},
			wantRes: nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(tt.fields.endpoint, WithCookie(tt.fields.cookie))

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
			for idx := strings.Index(content, Flag); idx >= 0; {
				idx2 := strings.Index(content[idx+len(Flag):], Flag)

				prefix := content[:idx-1]
				suffix := content[idx+9+idx2:]
				content = prefix + suffix
			}

			fmt.Println(content)
		})
	}
}
