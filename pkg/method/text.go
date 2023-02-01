package method

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// CheckText checking url by available text from http response
func CheckText(url string) (err error) {
	//fmt.Printf("Start check %s by text\n", url)

	var resp *http.Response
	resp, err = http.Get(url)
	if err != nil {
		//fmt.Println(err)
		return
	}

	var raw []byte
	_, err = resp.Body.Read(raw)
	if err != nil {
		//fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		//fmt.Println(err)
		return
	}
	text := string(bodyBytes)
	//fmt.Printf("text from body: %s\n", text)
	if !strings.Contains(text, "ok") {
		err = fmt.Errorf("ok not found in text")
	}

	//fmt.Printf("End check %s by text with error: %v\n", url, err)
	return
}
