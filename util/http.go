package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Get(cUrl string, params url.Values, headers map[string]string) (b []byte, err error) {

	req, err := http.NewRequest("GET", cUrl, nil)
	if err != nil {
		return
	}

	req.URL.RawQuery = params.Encode()

	return exec(req, headers)
}

// PostJson application/json
func PostJson(url string, data interface{}, headers map[string]string) (b []byte, err error) {

	jsonStr, err := json.Marshal(data)
	if err != nil {
		return
	}

	//Post的Content-Type默认值是application/x-www-form-urlencoded
	if _, ok := headers["Content-Type"]; !ok {
		if headers == nil {
			headers = map[string]string{"Content-Type": "application/json"}
		} else {
			headers["Content-Type"] = "application/json"
		}
	}

	return Post(url, jsonStr, headers)
}

func Post(url string, data []byte, headers map[string]string) (b []byte, err error) {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return
	}

	return exec(req, headers)
}

func exec(request *http.Request, headers map[string]string) (b []byte, err error) {

	// 遍历设置多个header
	for k, v := range headers {
		request.Header.Set(k, v)
	}

	client := &http.Client{}
	rep, err := client.Do(request)
	if err != nil {
		return
	}
	defer rep.Body.Close()

	b, err = ioutil.ReadAll(rep.Body)
	if err != nil {
		return
	}

	return
}
