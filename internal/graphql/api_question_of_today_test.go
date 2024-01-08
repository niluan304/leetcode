package graphql

import (
	"context"
	"reflect"
	"testing"
)

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
		wantRes *QuestionOfTodayRes
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
			},
			wantRes: &QuestionOfTodayRes{
				TodayRecord: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(tt.fields.endpoint, WithCookie(tt.fields.cookie))

			gotRes, err := c.QuestionOfToday(tt.args.ctx)
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
