package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type URL struct {
	Url          string   `json:"url"`
	Checks       []string `json:"checks"`
	MinChecksCnt int32    `json:"min_checks_cnt"`
}

type URLs struct {
	Urls []URL `json:"urls"`
}

func ReadJSON(dir, file string) (urls []URLs, err error) {
	var filePath string
	filePath, err = findFileWithPath(dir, file)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	var raw []byte
	raw, err = readFile(filePath)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	err = json.Unmarshal([]byte(raw), &urls)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	return
}

func findFileWithPath(dir string, searchFile string) (file string, err error) {
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && path == searchFile {
			fmt.Printf("%s\n", path)
			file = path
			return nil
		}
		return err
	})
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	return
}

func readFile(filePath string) (raw []byte, err error) {
	var target *os.File
	target, err = os.Open(filePath)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	defer func() {
		if err := target.Close(); err != nil {
			fmt.Printf("%v\n", err)
			return
		}
	}()

	_, err = target.Read(raw)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	return
}
