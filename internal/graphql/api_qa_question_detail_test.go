package graphql

import (
	"context"
	"fmt"
	"testing"
)

func TestClient_QaQuestionDetail(t *testing.T) {
	var ctx = context.Background()

	type client struct {
		cookie   map[string]string
		endpoint string
	}

	type fields client
	type args struct {
		ctx context.Context
		in  QaQuestionDetailReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *QaQuestionDetailRes
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				cookie:   nil,
				endpoint: EndpointZh,
			},
			args: args{
				ctx: ctx,
				in: QaQuestionDetailReq{
					Uuid: "G0n5iY",
				},
			},
			wantRes: nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(tt.fields.endpoint, WithCookie(tt.fields.cookie))
			gotRes, err := c.QaQuestionDetail(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("QaQuestionDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			fmt.Println(gotRes.QaQuestion.Content)
		})
	}
}
