package client

import (
	"github.com/jarcoal/httpmock"
	"net/http"
	"testing"
)

func TestMethodNotSupported(t *testing.T) {
	client := Client{BaseURL: "http://localhost:8080"}
	res := client.Request("abc", "a", "")

	assertEqual(t, "method not supported\n", res)
}

func TestGet(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://localhost:8080/data?key=a",
		httpmock.NewStringResponder(200, "AAA"))

	client := Client{BaseURL: "http://localhost:8080"}
	res := client.Request("get", "a", "")

	assertEqual(t, "AAA", res)
}

func TestDelete(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("DELETE", "http://localhost:8080/data?key=a",
		httpmock.NewStringResponder(204, ""))

	client := Client{BaseURL: "http://localhost:8080"}
	res := client.Request("delete", "a", "")

	assertEqual(t, "Deleted", res)
}

func TestDel(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("DELETE", "http://localhost:8080/data?key=a",
		httpmock.NewStringResponder(204, ""))

	client := Client{BaseURL: "http://localhost:8080"}
	res := client.Request("del", "a", "")

	assertEqual(t, "Deleted", res)
}

func TestPut(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("PUT", "http://localhost:8080/data?key=a&value=BBB",
		httpmock.NewStringResponder(201, ""))

	client := Client{BaseURL: "http://localhost:8080"}
	res := client.Request("PUT", "a", "BBB")

	assertEqual(t, "Your key value pair was created", res)
}

func TestPutNoValueSet(t *testing.T) {
	client := Client{BaseURL: "http://localhost:8080"}
	res := client.Request("PUT", "a", "")

	assertEqual(t, "no value set", res)
}

func TestNoConnection(t *testing.T) {
	client := Client{BaseURL: "http://localhost:8081"}
	res := client.Request("GET", "a", "")

	assertEqual(t, "", res)
}

func assertEqual(t *testing.T, expected string, value string) {
	if value != expected {
		t.Errorf("expected \"%v\", got \"%v\"", expected, value)
	}
}

func Test_statusCodeToMessage(t *testing.T) {
	type args struct {
		status     int
		bodyString string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Ok", args: struct {
				status     int
				bodyString string
			}{
				status: http.StatusOK, bodyString: ""},
			want: "Your value was updated",
		},
		{
			name: "Created", args: struct {
				status     int
				bodyString string
			}{
				status: http.StatusCreated, bodyString: ""},
			want: "Your key value pair was created",
		},
		{
			name: "Deleted", args: struct {
				status     int
				bodyString string
			}{
				status: http.StatusNoContent, bodyString: ""},
			want: "Deleted",
		},
		{
			name: "To large key or value", args: struct {
				status     int
				bodyString string
			}{
				status: http.StatusRequestEntityTooLarge, bodyString: ""},
			want: "Your key or value is over 200 characters",
		},
		{
			name: "Database out of memory", args: struct {
				status     int
				bodyString string
			}{
				status: http.StatusInsufficientStorage, bodyString: ""},
			want: "Database is out of memory",
		},
		{
			name: "BadRequest", args: struct {
				status     int
				bodyString string
			}{
				status: http.StatusBadRequest, bodyString: "your error"},
			want: "Your Request is invalid: your error\n",
		},
		{
			name: "Created", args: struct {
				status     int
				bodyString string
			}{
				status: http.StatusForbidden, bodyString: "no access"},
			want: "Unexpected server response: [403] no access\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := statusCodeToMessage(tt.args.status, tt.args.bodyString); got != tt.want {
				t.Errorf("statusCodeToMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
