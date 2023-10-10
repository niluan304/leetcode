package server

import "testing"

func Test_server_do(t *testing.T) {
	tests := []struct {
		titleSlug string
	}{
		{titleSlug: "moving-stones-until-consecutive"},
		{titleSlug: "coloring-a-border"},
		{titleSlug: "uncrossed-lines"},
		{titleSlug: "escape-a-large-maze"},
		{titleSlug: "minimum-swaps-to-group-all-1s-together"},
		{titleSlug: "analyze-user-website-visit-pattern"},
		{titleSlug: "minimum-score-triangulation-of-polygon"},
		{titleSlug: "find-words-that-can-be-formed-by-characters"},
		{titleSlug: "moving-stones-until-consecutive-ii"},
		{titleSlug: "simplify-path"},
		{titleSlug: "h-index"},
	}
	for _, tt := range tests {
		t.Run(tt.titleSlug, func(t *testing.T) {
			s := Server{}
			if err := s.TitleSlug(tt.titleSlug); err != nil {
				t.Logf("TitleSlug() error = %v", err)
			}
		})

		sleep(1, 4)
	}
}

func TestServer_BuildById(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "1", args: args{id: "1"}, wantErr: false},
		{name: "2", args: args{id: "2"}, wantErr: false},
		{name: "3", args: args{id: "3"}, wantErr: false},
		{name: "4", args: args{id: "4"}, wantErr: false},
		{name: "5", args: args{id: "5"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Server{}
			if err := s.Id(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Id() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_TitleSlug(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name          string
		args          args
		wantTitleSlug string
		wantErr       bool
	}{
		{name: "1", args: args{id: "1"}, wantTitleSlug: "two-sum", wantErr: false},
		{name: "2", args: args{id: "2"}, wantTitleSlug: "add-two-numbers", wantErr: false},
		{name: "3", args: args{id: "3"}, wantTitleSlug: "longest-substring-without-repeating-characters", wantErr: false},
		{name: "4", args: args{id: "4"}, wantTitleSlug: "median-of-two-sorted-arrays", wantErr: false},
		{name: "5", args: args{id: "5"}, wantTitleSlug: "longest-palindromic-substring", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Server{}
			gotTitleSlug, err := s.idToSlug(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("TitleSlug() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTitleSlug != tt.wantTitleSlug {
				t.Errorf("TitleSlug() gotTitleSlug = %v, want %v", gotTitleSlug, tt.wantTitleSlug)
			}
		})
	}
}
