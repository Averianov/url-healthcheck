package method

import (
	"fmt"
	"net/http"
)

// CheckStatusCode checking url by StatusCode from http response
func CheckStatusCode(url string) (err error) {
	//fmt.Printf("Start check %s by status code\n", url)

	var resp *http.Response
	resp, err = http.Get(url)
	if err != nil {
		//fmt.Println(err)
		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("Invalid http status: %d - %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	//fmt.Printf("End check %s by status code with error: %v\n", url, err)
	return
}
