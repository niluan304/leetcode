package server

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/skratchdot/open-golang/open"
	"gopkg.in/yaml.v3"

	"leetcode/cmd/leetcode/graphql"
	"leetcode/cmd/leetcode/tmpl"
)

type Server struct {
	configPath string
}

func New(configPath string) *Server {
	return &Server{
		configPath: configPath,
	}
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

type Config struct {
	Parse []struct {
		Name string `yaml:"name"`
		Open bool   `yaml:"open"`
		App  string `yaml:"app"`
	} `yaml:"parse"`

	Url struct {
		Open bool   `yaml:"open"`
		App  string `yaml:"app"`
	} `yaml:"url"`
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

	file, err := os.Open(s.configPath)
	if err != nil {
		return err
	}

	var config Config
	err = yaml.NewDecoder(file).Decode(&config)
	if err != nil {
		return err
	}

	// 解析模板
	for _, cfg := range config.Parse {
		var parser tmpl.Parser
		switch strings.ToLower(cfg.Name) {
		case "en":
			parser = tmpl.NewParserEN(question)
		case "zh":
			parser = tmpl.NewParserZH(question)
		case "leetcode":
			parser = tmpl.NewParserLeetcode(question.Pkg())
		case "solution":
			parser = tmpl.NewParserSolution(question)
		case "test":
			parser = tmpl.NewParserEndlessTest(question)
		case "sample":
			parser = tmpl.NewParserSamples(question)
		}

		p := tmpl.Parse{
			Open:   cfg.Open,
			App:    cfg.App,
			Parser: parser,
		}
		err = p.Save(path)
		if err != nil {
			return err
		}
	}

	if config.Url.Open {
		err = open.Run(fmt.Sprintf("https://leetcode.cn/problems/%s/description/", titleSlug))
		if err != nil {
			return err
		}
	}

	fmt.Printf("submit: %s. %s\n", question.QuestionFrontendId, question.TitleSlug)
	time.Sleep(time.Second)
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

func (s *Server) Article(article string) (err error) {
	if article[len(article)-1] == '/' {
		article = article[:len(article)-1]
	}

	i := strings.LastIndex(article, "/")
	article = article[i+1:]

	var ctx = context.Background()
	client := graphql.New(graphql.EndpointZh)
	res, err := client.SolutionArticle(ctx, graphql.SolutionArticleReq{
		Slug: article,
	})
	if err != nil {
		return err
	}
	const Flag = "```"
	content := res.SolutionArticle.Content
	for idx := strings.Index(content, Flag); idx >= 0; {
		idx2 := strings.Index(content[idx+len(Flag):], Flag+"\n")
		if idx2 == -1 {
			break
		}

		prefix := content[:idx-1]
		suffix := content[idx+9+idx2:]
		content = prefix + suffix
	}

	fmt.Println(content)
	return nil
}
