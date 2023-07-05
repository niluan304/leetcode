package server

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"

	"leetcode/cmd/leetcode/graphql"
	"leetcode/cmd/leetcode/tmpl"
)

type Server struct{}

func New(configPath string) *Server {
	return &Server{}
}

func (s *Server) Run() {
	var v struct {
		Records []struct {
			TitleSlug string `json:"title_slug"`
		} `json:"RECORDS"`
	}

	file, err := os.Open("D:/Users/Feng/Desktop/problem.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&v)
	if err != nil {
		panic(err)
	}

	list := v.Records
	for _, record := range list {
		err := s.Do(record.TitleSlug)
		if err != nil {
			panic(err)
		}

		sleep(2, 5)
	}
}

func (s *Server) Do(titleSlug string) (err error) {
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
	list := []tmpl.Parser{
		tmpl.NewParserEN(question),
		tmpl.NewParserZN(question),
		tmpl.NewParserCase(question),
		tmpl.NewParserSolution(question),
		tmpl.NewParserUnitCase(question),
		tmpl.NewParserLeetcode(question.Pkg()),
	}
	for _, elem := range list {
		p := &tmpl.Parse{Parser: elem}

		err = p.Save(path)
		if err != nil {
			return err
		}
	}

	return nil
}

func sleep(min int, max int) {
	d := min + rand.Intn(max-min)

	time.Sleep(time.Millisecond * time.Duration(d*100))
}
