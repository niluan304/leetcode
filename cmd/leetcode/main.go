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
	id     = flag.String("i", "", "leetcode problem id")
	slug   = flag.String("s", "", "leetcode problem slug")
	output = flag.String("o", "", "output file path")
	_      = flag.String("e", "", "os.Exit(0)")
	once   = new(sync.Once)
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
	s := server.New("")

	for {
		if *id != "" {
			err := s.BuildById(*id)
			if err != nil {
				fmt.Println(err)

			}
		}

		if *slug != "" {
			err := s.Build(*slug)
			if err != nil {
				fmt.Println(err)
			}
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
	}
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
	}
	return nil
}
