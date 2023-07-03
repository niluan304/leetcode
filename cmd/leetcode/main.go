package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"

	"leetcode/cmd/leetcode/server"
)

var (
	id     = flag.String("i", "", "leetcode problem id")
	slug   = flag.String("s", "", "leetcode problem slug")
	output = flag.String("o", "", "output file path")
)

func init() {
	// 自定义使用信息
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
	}

	// 解析命令行参数
	flag.Parse()

	// 打印用法
	flag.Usage()
}

func main() {
	var err error
	defer func() {
		if err != nil {
			fmt.Println(err)
		}
	}()

	if *id != "" {
		err = GetQuestionId(*id)
		return
	}

	if *slug != "" {
		err = GetTitleSlug(*slug)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		// 读取用户输入
		fmt.Print("请输入命令参数：")
		flag.Usage()
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取输入出错：", err)
			continue
		}

		// 去除输入中的换行符
		input = strings.TrimSpace(input)
		input = strings.TrimLeft(input, "-")

		// 处理用户输入
		if input == "exit" {
			fmt.Println("程序退出")
			os.Exit(0)
		}

		switch input {
		case "i":
			fmt.Print("请输入leetcode problem id：")
			id, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("读取输入出错：", err)
				continue
			}
			err = GetQuestionId(strings.TrimSpace(id))
			if err != nil {
				fmt.Println("获取题目出错：", err)
				continue
			}
		case "s":
			fmt.Print("请输入leetcode problem slug：")
			slug, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("读取输入出错：", err)
				continue
			}
			err = GetTitleSlug(strings.TrimSpace(slug))
			if err != nil {
				fmt.Println("获取题目出错：", err)
				continue
			}
		}
	}
}

type Problem struct {
	FrontendId string `json:"frontend_id"`
	TitleSlug  string `json:"title_slug"`
}

func GetQuestionId(id string) (err error) {
	var v struct {
		Records []Problem `json:"RECORDS"`
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
		if record.FrontendId == id {
			return GetTitleSlug(record.TitleSlug)
		}
	}

	return errors.New("not found: " + id)
}

func GetTitleSlug(slug string) (err error) {
	s := server.New("")
	return s.Do(slug)
}
