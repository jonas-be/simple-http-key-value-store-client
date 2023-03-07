package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const baseURL = "http://localhost:8080"

func main() {
	methodFlag := flag.String("m", "", "method to execute")
	keyFlag := flag.String("key", "", "key to access")
	valueFlag := flag.String("value", "", "value to set")

	flag.Parse()

	request(methodFlag, keyFlag, valueFlag)
}

func request(methodFlag *string, keyFlag *string, valueFlag *string) {
	client := &http.Client{}

	switch strings.ToLower(*methodFlag) {
	case "get":
		bodyString, status, done := httpRequest(http.MethodGet, fmt.Sprintf("%v/data?key=%v", baseURL, *keyFlag), client)
		if done {
			return
		}
		fmt.Printf("[ %v ]: %v\n", status, bodyString)
		break
	case "put":
		if *valueFlag == "" {
			fmt.Println("no value set")
			return
		}
		bodyString, status, done := httpRequest(http.MethodPut, fmt.Sprintf("%v/data?key=%v&value=%v", baseURL, *keyFlag, *valueFlag), client)
		if done {
			return
		}
		fmt.Printf("[ %v ]: %v\n", status, bodyString)
		break
	case "delete", "del":
		bodyString, status, done := httpRequest(http.MethodDelete, fmt.Sprintf("%v/data?key=%v&value=%v", baseURL, *keyFlag, *valueFlag), client)
		if done {
			return
		}
		fmt.Printf("[ %v ]: %v\n", status, bodyString)
		break
	default:
		fmt.Println("method not supported")
	}
}

func httpRequest(method string, url string, client *http.Client) (string, int, bool) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return "", 0, true
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", resp.StatusCode, true
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", resp.StatusCode, true
	}
	sb := string(body)
	return sb, resp.StatusCode, false
}
