package graphql

import (
	"context"
	"reflect"
	"testing"
)

func TestClient_ProblemsQuestionList(t *testing.T) {
	var ctx = context.Background()

	type client struct {
		cookie   map[string]string
		endpoint string
	}

	type fields client
	type args struct {
		ctx context.Context
		in  ProblemQuestionReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *ProblemQuestionRes
		wantErr bool
	}{
		{
			name: "EndpointZh",
			fields: fields{
				cookie:   nil,
				endpoint: EndpointZh,
			},
			args: args{
				ctx: ctx,
				in: ProblemQuestionReq{
					CategorySlug: "",
					Skip:         0,
					Limit:        20,
					Filters: ProblemQuestionFilters{
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
			c := New(tt.fields.endpoint, WithCookie(tt.fields.cookie))

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
