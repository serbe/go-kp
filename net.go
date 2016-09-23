package gokp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func getJSON(url string) ([]byte, error) {
	fmt.Println(url)
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
