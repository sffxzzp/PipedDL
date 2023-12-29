package main

import (
	"io"
	"net/http"
	"time"
)

//lint:ignore U1000 Ignore unused function temporarily for debugging
func httpGet(url string) []byte {
	retry := 3
	for retry > 0 {
		res, err := http.Get(url)
		if err != nil {
			retry--
			continue
		}
		defer res.Body.Close()
		rbytes, err := io.ReadAll(res.Body)
		if err != nil {
			retry--
			continue
		}
		return rbytes
	}
	return []byte{}
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func httpPost(url string, contentType string, body io.Reader) []byte {
	retry := 3
	for retry > 0 {
		res, err := http.Post(url, contentType, body)
		if err != nil {
			retry--
			continue
		}
		defer res.Body.Close()
		rbytes, err := io.ReadAll(res.Body)
		if err != nil {
			retry--
			continue
		}
		return rbytes
	}
	return []byte{}
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func checkAvail(url string) bool {
	client := &http.Client{Timeout: 10 * time.Second}
	req, _ := http.NewRequest("GET", url, nil)
	_, err := client.Do(req)
	return err == nil
}
