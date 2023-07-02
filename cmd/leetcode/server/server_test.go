package server

import "testing"

func TestLoad(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.titleSlug, func(t *testing.T) {
			s := New("")
			if err := s.do(tt.titleSlug); err != nil {
				t.Logf("do() error = %v", err)
			}
		})

		sleep(1, 4)
	}
}
