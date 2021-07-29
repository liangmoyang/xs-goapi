package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// PostStruct application/json
func PostStruct(url string, data interface{}) (b []byte, err error) {

	jsonStr, err := json.Marshal(data)
	if err != nil {
		return
	}

	return Post(url, jsonStr, "application/json")
}

func Post(url string, data []byte, contentType string) (b []byte, err error) {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}

	rep, err := client.Do(req)
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
