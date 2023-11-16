package problem

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	"leetcode/cmd/leetcode/server"
)

func TestUpdateFilename(t *testing.T) {
	var dirs []string

	err := filepath.Walk("./", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			t.Log(err)
			return err
		}

		if !info.IsDir() || strings.HasPrefix(path, ".git") {
			return nil
		}

		//var dir = filepath.Dir(path)
		//// 不修改 problem下的文件
		//if dir == filepath.Dir("./") {
		//	return nil
		//}

		dirs = append(dirs, path)
		return nil
	})
	if err != nil {
		t.Errorf("TestUpdateFilename() error = %v", err)
	}

	var toInt = func(s string) int {
		n, _, _ := strings.Cut(s, "_")
		i, _ := strconv.Atoi(n)
		return i
	}

	sort.Slice(dirs, func(i, j int) bool {
		x, y := dirs[i], dirs[j]
		return toInt(x) < toInt(y)
	})

	for _, dir := range dirs {
		if !strings.Contains(dir, "_") ||
			strings.HasPrefix(dir, "problem") ||
			strings.Contains(dir, "A") {
			continue
		}

		fmt.Println(dir)

		solution, err := os.ReadFile(filepath.Join(dir, "solution.go"))
		if err != nil {
			solution = []byte("package main\n")
			fmt.Println(err)
		}
		leetcode, err := os.ReadFile(filepath.Join(dir, "solution_leetcode.go"))
		if err != nil {
			leetcode = []byte("package main\n")
			fmt.Println(err)
		}
		//os.Remove(filepath.Join(dir, "case.go"))

		s := server.New("D:\\__Project\\leetcode\\config.yml")
		err = s.Id(strconv.Itoa(toInt(dir)))
		if err != nil {
			fmt.Println(err)
			continue
		}

		a := filepath.Join("problem", dir, "solution.go")
		b := filepath.Join("problem", dir, "solution_leetcode.go")

		err = os.WriteFile(a, solution, 644)
		if err != nil {
			fmt.Println("open", a, err)
		}

		time.Sleep(50 * time.Millisecond)
		err = os.WriteFile(b, leetcode, 644)
		if err != nil {
			fmt.Println("open", b, err)
		}
	}
}
