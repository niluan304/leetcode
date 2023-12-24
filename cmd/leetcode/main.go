package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/pkg/errors"

	"leetcode/cmd/leetcode/server"
)

var (
	id      = flag.String("i", "", "leetcode problem id")
	slug    = flag.String("s", "", "leetcode problem slug")
	output  = flag.String("o", "", "output file path")
	today   = flag.String("t", "", "today of problem")
	article = flag.String("a", "", "solution article")
	_       = flag.String("e", "", "os.Exit(0)")
	once    = new(sync.Once)
)

func init() {
	// 自定义使用信息
	flag.Usage = func() {
		_, err := fmt.Fprintln(os.Stderr, "Options:")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fprintln: %v\n", err)
		}
		flag.PrintDefaults()
	}

	// 解析命令行参数
	flag.Parse()
}

func main() {
	s := server.New("config.yml")

	if strings.ToUpper(*today) == "YES" {
		err := s.Today()
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	if *id != "" {
		err := s.Id(*id)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	if *slug != "" {
		err := s.TitleSlug(*slug)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	if *article != "" {
		err := s.Article(*article)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	// reset
	*id = ""
	*slug = ""

	once.Do(flag.Usage)
	fmt.Print("请输入命令参数：")

	// 读取用户输入
	err := Scan()
	if err != nil {
		fmt.Println(err)
	}

	main()
}

func Scan() (err error) {
	// 读取用户输入
	var input = ""
	_, err = fmt.Scan(&input)
	if err != nil {
		return errors.Wrap(err, "fail scan input")
	}

	input = strings.TrimLeft(input, "-")

	// 处理用户输入
	switch input {
	case "exit":
		os.Exit(0)
	case "i":
		fmt.Print("请输入leetcode problem id：")
		_, err = fmt.Scan(id)
		if err != nil {
			return errors.Wrap(err, "fail scan problem id ")
		}
	case "s":
		fmt.Print("请输入leetcode problem slug：")
		_, err = fmt.Scan(slug)
		if err != nil {
			return errors.Wrap(err, "fail scan problem slug")
		}
	case "t":
		fmt.Print("请确认是否获取每日一题 YES/NO：")
		_, err = fmt.Scan(today)
		if err != nil {
			return errors.Wrap(err, "fail scan today")
		}
	case "a":
		fmt.Print("请确认题解链接：")
		_, err = fmt.Scan(article)
		if err != nil {
			return errors.Wrap(err, "fail scan today")
		}
	}
	return nil
}
