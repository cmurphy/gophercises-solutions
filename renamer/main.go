package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func rename(path string, info fs.FileInfo, err error) error {
	if err != nil {
		return err
	}
	r := regexp.MustCompile(`(?P<subj>.*)_(?P<num>\d\d\d).(?P<ext>\w\w\w)`)
	if matches := r.FindStringSubmatch(info.Name()); len(matches) >= 3 {
		subject := matches[1]
		num, err := strconv.Atoi(matches[2])
		ext := matches[3]
		if err != nil {
			return err
		}
		if err := os.Rename(path, filepath.Join(filepath.Dir(path), fmt.Sprintf("%s (%d of 100).%s", subject, num, ext))); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	err := filepath.Walk("./sample", rename)
	if err != nil {
		panic(err)
	}
}
