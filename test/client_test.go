package test

import (
	"simple-http-key-value-store-client/internal/client"
	"testing"
)
import "github.com/jarcoal/httpmock"

func TestMethodNotSupported(t *testing.T) {
	client := client.Client{BaseURL: "http://localhost:8080"}
	res := client.Request("abc", "a", "")

	assertEqual(t, "method not supported\n", res)
}

func TestGet(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://localhost:8080/data?key=a",
		httpmock.NewStringResponder(200, `AAA`))

	client := client.Client{BaseURL: "http://localhost:8080"}
	res := client.Request("get", "a", "")

	assertEqual(t, "[ 200 ]: AAA\n", res)
}

func TestDelete(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("DELETE", "http://localhost:8080/data?key=a",
		httpmock.NewStringResponder(200, ``))

	client := client.Client{BaseURL: "http://localhost:8080"}
	res := client.Request("delete", "a", "")

	assertEqual(t, "[ 200 ]: \n", res)
}

func TestDel(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("DELETE", "http://localhost:8080/data?key=a",
		httpmock.NewStringResponder(200, ``))

	client := client.Client{BaseURL: "http://localhost:8080"}
	res := client.Request("del", "a", "")

	assertEqual(t, "[ 200 ]: \n", res)
}

func TestPut(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("PUT", "http://localhost:8080/data?key=a&value=BBB",
		httpmock.NewStringResponder(201, ``))

	client := client.Client{BaseURL: "http://localhost:8080"}
	res := client.Request("PUT", "a", "BBB")

	assertEqual(t, "[ 201 ]: \n", res)
}

func TestPutNoValueSet(t *testing.T) {
	client := client.Client{BaseURL: "http://localhost:8080"}
	res := client.Request("PUT", "a", "")

	assertEqual(t, "no value set", res)
}

func TestNoConnection(t *testing.T) {
	client := client.Client{BaseURL: "http://localhost:8081"}
	res := client.Request("GET", "a", "")

	assertEqual(t, "", res)
}

func assertEqual(t *testing.T, expected string, value string) {
	if value != expected {
		t.Errorf("expected \"%v\", got \"%v\" exists", expected, value)
	}
}
