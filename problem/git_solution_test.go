package problem

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
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

	for _, dir := range dirs[1:] {
		invalid := true
		var fileList []string

		filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
			if path == dir {
				return nil
			}

			fileList = append(fileList, path)

			if !invalid {
				return nil
			}

			if strings.HasPrefix(filepath.Base(path), "README") {
				return nil
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			fd := bufio.NewReader(file)
			lines := 0
			for {
				_, err := fd.ReadString('\n')
				if err != nil {
					break
				}
				lines++
			}

			if lines > 3 {
				invalid = false
			}

			return nil
		})

		if invalid || len(fileList) < 4 {
			fmt.Println(dir, fileList)

			err := os.RemoveAll(dir)
			if err != nil {
				t.Error(err)
			}
		}
	}
}
