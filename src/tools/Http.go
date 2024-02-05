package tools

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

// HttpGet 获取http内容
func HttpGet(url string) (string, error) {
	resp, err := http.Get(url)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))

	return string(body), err
}

// HttpPost 0 json 1 form
func HttpPost(url string, requestBody []byte, bodyType int) (string, error) {
	var bt string
	if bodyType == 0 {
		bt = "application/json"
	} else if bodyType == 1 {
		bt = "application/form"
	} else {
		return "", errors.New("不支持的body类型")
	}

	resp, err := http.Post(url, bt, bytes.NewBuffer(requestBody))
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}
