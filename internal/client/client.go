package client

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	BaseURL string
}

func (cl Client) Request(methodFlag string, keyFlag string, valueFlag string) string {
	httpClient := &http.Client{}

	switch strings.ToLower(methodFlag) {
	case "get":
		return cl.getRequest(keyFlag, httpClient)
	case "put":
		return cl.putRequest(valueFlag, keyFlag, httpClient)
	case "delete", "del":
		return cl.deleteRequest(keyFlag, httpClient)
	default:
		return fmt.Sprintf("method not supported\n")
	}
}

func (cl Client) getRequest(keyFlag string, httpClient *http.Client) string {
	bodyString, status, err := cl.httpRequest(http.MethodGet, fmt.Sprintf("%v/data?key=%v", cl.BaseURL, keyFlag), httpClient)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if status == 200 {
		return bodyString
	}
	return statusCodeToMessage(status, bodyString)
}

func (cl Client) deleteRequest(keyFlag string, httpClient *http.Client) string {
	bodyString, status, err := cl.httpRequest(http.MethodDelete, fmt.Sprintf("%v/data?key=%v", cl.BaseURL, keyFlag), httpClient)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return statusCodeToMessage(status, bodyString)
}

func (cl Client) putRequest(valueFlag string, keyFlag string, httpClient *http.Client) string {
	if valueFlag == "" {
		return "no value set"
	}
	bodyString, status, err := cl.httpRequest(http.MethodPut, fmt.Sprintf("%v/data?key=%v&value=%v", cl.BaseURL, keyFlag, valueFlag), httpClient)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return statusCodeToMessage(status, bodyString)
}

func (cl Client) httpRequest(method string, url string, httpClient *http.Client) (string, int, error) {
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
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", resp.StatusCode, err
	}
	sb := string(body)
	return sb, resp.StatusCode, nil
}

func statusCodeToMessage(status int, bodyString string) string {
	switch status {
	case http.StatusOK:
		return "Your value was updated"
	case http.StatusCreated:
		return "Your key value pair was created"
	case http.StatusNoContent:
		return "Deleted"
	case http.StatusRequestEntityTooLarge:
		return "Your key or value is over 200 characters"
	case http.StatusInsufficientStorage:
		return "Database is out of memory"
	case http.StatusBadRequest:
		return fmt.Sprintf("Your Request is invalid: %v\n", bodyString)
	default:
		return fmt.Sprintf("Unexpected server response: [%v] %v\n", status, bodyString)
	}
}
