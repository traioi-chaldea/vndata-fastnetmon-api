package util

import (
  "io"
  "io/ioutil"
  "net/http"
)

func Curl(method string, url string, body io.Reader, authorization string) string {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		ErrorLogger.Print(err)
	}

	req.Header.Set("Accept", "application/json")

	if authorization != "" {
		req.Header.Set("Content-type", "application/json")
		req.Header.Set("Authorization", authorization)
	} else {
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ErrorLogger.Print(err)
	}

	defer resp.Body.Close()

	t_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ErrorLogger.Print(err)
	}

	return string(t_body)
}

