package client

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Client struct {
	BaseURL string
}

func (client Client) Request(methodFlag string, keyFlag string, valueFlag string) string {
	httpClient := &http.Client{}

	switch strings.ToLower(methodFlag) {
	case "get":
		return client.getRequest(keyFlag, httpClient)
	case "put":
		return client.putRequest(valueFlag, keyFlag, httpClient)
	case "delete", "del":
		return client.deleteRequest(keyFlag, httpClient)
	default:
		return fmt.Sprintf("method not supported\n")
	}
}

func (client Client) getRequest(keyFlag string, httpClient *http.Client) string {
	bodyString, status, err := client.httpRequest(http.MethodGet, fmt.Sprintf("%v/data?key=%v", client.BaseURL, keyFlag), httpClient)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return fmt.Sprintf("[ %v ]: %v\n", status, bodyString)
}

func (client Client) deleteRequest(keyFlag string, httpClient *http.Client) string {
	bodyString, status, err := client.httpRequest(http.MethodDelete, fmt.Sprintf("%v/data?key=%v", client.BaseURL, keyFlag), httpClient)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return fmt.Sprintf("[ %v ]: %v\n", status, bodyString)
}

func (client Client) putRequest(valueFlag string, keyFlag string, httpClient *http.Client) string {
	if valueFlag == "" {
		return "no value set"
	}
	bodyString, status, err := client.httpRequest(http.MethodPut, fmt.Sprintf("%v/data?key=%v&value=%v", client.BaseURL, keyFlag, valueFlag), httpClient)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return fmt.Sprintf("[ %v ]: %v\n", status, bodyString)
}

func (client Client) httpRequest(method string, url string, httpClient *http.Client) (string, int, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return "", 0, err
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", 0, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", resp.StatusCode, err
	}
	sb := string(body)
	return sb, resp.StatusCode, nil
}
