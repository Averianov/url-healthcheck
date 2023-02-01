package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type URL struct {
	Url    string
	Checks []string
	Count  int `json:"min_checks_cnt"`
}

type URLConfig struct {
	Urls []URL
}

// ReadJSON search and unmarshal data from json file
func ReadJSON(fileName string) (urls URLConfig, err error) {
	var filePath string
	filePath, err = os.Getwd()
	if err != nil {
		return
	}

	filePath = filePath + "/" + fileName
	//fmt.Printf("%s\n", filePath)
	var raw []byte
	raw, err = ioutil.ReadFile(filePath)
	if err != nil {
		//fmt.Println(err)
		return
	}

	//fmt.Printf("%v\n", raw)
	err = json.Unmarshal([]byte(raw), &urls)
	if err != nil {
		//fmt.Println(err)
	}
	return
}
