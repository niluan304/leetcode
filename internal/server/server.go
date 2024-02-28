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

	"github.com/niluan304/leetcode/internal/api"
	"github.com/niluan304/leetcode/internal/tmpl"
)

type Config struct {
	Leetcode struct {
		Path string `yaml:"path"`
	} `yaml:"leetcode"`

	Template []tmpl.Config `yaml:"template"`

	Url struct {
		Open bool   `yaml:"open"`
		App  string `yaml:"app"`
	} `yaml:"url"`
}

type Server struct {
	config *Config
}

func NewServer(configPath string) (*Server, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, errors.Wrap(err, "fail to open config file: "+configPath)
	}

	var config *Config
	err = yaml.NewDecoder(file).Decode(&config)
	if err != nil {
		return nil, errors.Wrap(err, "fail to decode yaml file: "+configPath)
	}

	return &Server{
		config: config,
	}, nil
}

func (s *Server) idToSlug(id string) (titleSlug string, err error) {
	ctx := context.Background()
	client := api.New(api.EndpointZh)

	res, err := client.ProblemsetQuestionList(ctx, api.ProblemQuestionReq{
		CategorySlug: "",
		Skip:         0,
		Limit:        50,
		Filters: api.ProblemQuestionFilters{
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
	ctx := context.Background()
	client := api.New(api.EndpointZh)

	res, err := client.QuestionData(ctx, api.QuestionDataReq{TitleSlug: titleSlug})
	if err != nil {
		return err
	}
	question, err := res.Question.ReUnmarshal()
	if err != nil {
		return err
	}

	dir := question.Dir()
	path := filepath.Join(s.config.Leetcode.Path, dir)
	_ = os.MkdirAll(path, os.ModePerm)

	if question.IsPaidOnly {
		log.Println("IsPaidOnly", path)
		return errors.New("会员题目, 无法查看")
	}

	// 解析模板
	data := tmpl.NewData(question)
	for _, cfg := range s.config.Template {
		err = tmpl.NewTemplate(cfg, data).Save(path)
		if err != nil {
			return err
		}
	}

	if s.config.Url.Open {
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
	ctx := context.Background()
	client := api.New(api.EndpointZh)
	res, err := client.QuestionOfToday(ctx, api.QuestionOfTodayReq{})
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

	ctx := context.Background()
	client := api.New(api.EndpointZh)
	res, err := client.SolutionArticle(ctx, api.SolutionArticleReq{
		Slug: article,
	})
	if err != nil {
		return err
	}
	const Flag = "```"
	content := res.SolutionArticle.Content
	//for idx := strings.Index(content, Flag); idx >= 0; {
	//	idx2 := strings.Index(content[idx+len(Flag):], Flag+"\n")
	//	if idx2 == -1 {
	//		break
	//	}
	//
	//	prefix := content[:idx-1]
	//	suffix := content[idx+9+idx2:]
	//	content = prefix + suffix
	//}

	fmt.Println(content)
	return nil
}
