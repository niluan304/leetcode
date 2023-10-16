package server

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
	"github.com/skratchdot/open-golang/open"

	"leetcode/cmd/leetcode/graphql"
	"leetcode/cmd/leetcode/tmpl"
)

type Server struct{}

func New(configPath string) *Server {
	return &Server{}
}

func (s *Server) idToSlug(id string) (titleSlug string, err error) {
	var ctx = context.Background()
	client := graphql.New(graphql.EndpointZh)

	res, err := client.ProblemsetQuestionList(ctx, graphql.ProblemQuestionReq{
		CategorySlug: "",
		Skip:         0,
		Limit:        50,
		Filters: graphql.ProblemQuestionFilters{
			Difficulty:     "",
			Status:         "",
			SearchKeywords: id,
			Tags:           nil,
			ListId:         "",
		},
	})
	if err != nil {
		return "", err
	}

	list := res.ProblemsetQuestionList.Questions
	if len(list) == 0 {
		return "", errors.New("未找到题目")
	}

	for _, item := range list {
		if item.FrontendQuestionId == id {
			return item.TitleSlug, nil
		}
	}

	return "", errors.New("未找到题目")
}

func (s *Server) Id(id string) (err error) {
	titleSlug, err := s.idToSlug(id)
	if err != nil {
		return err
	}

	return s.TitleSlug(titleSlug)
}

func (s *Server) TitleSlug(titleSlug string) (err error) {
	var ctx = context.Background()
	client := graphql.New(graphql.EndpointZh)

	res, err := client.QuestionData(ctx, titleSlug)
	if err != nil {
		return err
	}
	question, err := res.Question.ReUnmarshal()
	if err != nil {
		return err
	}

	dir := question.Dir()
	path := filepath.Join("./problem", dir)
	_ = os.MkdirAll(path, os.ModePerm)

	if question.IsPaidOnly {
		log.Println("IsPaidOnly", path)
		return errors.New("会员题目, 无法查看")
	}

	// 解析模板
	list := []tmpl.Parse{
		{Open: false, Parser: tmpl.NewParserEN(question)},
		{Open: false, Parser: tmpl.NewParserZH(question)},
		{Open: false, Parser: tmpl.NewParserLeetcode(question.Pkg())},
		{Open: false, Parser: tmpl.NewParserSamples(question)},
		{Open: true, Parser: tmpl.NewParserSolution(question)},
		{Open: true, Parser: tmpl.NewParserEndlessTest(question)},
	}
	for _, p := range list {
		err = p.Save(path)
		if err != nil {
			return err
		}
	}
	err = open.Run(fmt.Sprintf("https://leetcode.cn/problems/%s/description/", titleSlug))
	if err != nil {
		return err
	}

	return nil
}

func sleep(min int, max int) {
	d := min + rand.Intn(max-min)

	time.Sleep(time.Millisecond * time.Duration(d*100))
}

func (s *Server) Today() (err error) {
	var ctx = context.Background()
	client := graphql.New(graphql.EndpointZh)
	res, err := client.QuestionOfToday(ctx)
	if err != nil {
		return err
	}

	for _, record := range res.TodayRecord {
		err = s.TitleSlug(record.Question.TitleSlug)
		if err != nil {
			return err
		}
	}

	return nil
}
