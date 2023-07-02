package server

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

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
		err := s.do(record.TitleSlug)
		if err != nil {
			panic(err)
		}

		sleep(2, 5)
	}
}

func (s *Server) do(titleSlug string) (err error) {
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
	path := filepath.Join("./data", dir)
	_ = os.MkdirAll(path, os.ModePerm)

	if question.IsPaidOnly {
		log.Println("IsPaidOnly", path)
		return nil
	}

	// 新建各语言solution文件
	solution, err := tmpl.NewSolution(question)
	if err != nil {
		return err
	}
	err = solution.Save(path)
	if err != nil {
		return err
	}

	// 解析题目描述模板
	list := []struct {
		name   string
		parser tmpl.Parser
	}{
		{name: "en", parser: tmpl.NewEN(question)},
		{name: "zh", parser: tmpl.NewZN(question)},
	}
	for _, elem := range list {
		l := tmpl.Lang{Parser: elem.parser}

		err = l.Save(path)
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
