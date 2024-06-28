package web

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestGetRequest(t *testing.T) {
	req := HttpRequest{
		Method: http.MethodGet,
		Url:    "http://pie.dev/get",
	}
	res, err := req.Call()
	if err != nil {
		t.Fatal(err)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(b))
}

func TestPostRequest(t *testing.T) {
	body := make(map[string]string)
	body["hello"] = "world"
	b, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	req := HttpRequest{
		Method:  http.MethodPost,
		Url:     "http://pie.dev/post",
		Body:    b,
		Headers: headers,
	}
	res, err := req.Call()
	if err != nil {
		t.Fatal(err)
	}

	b, err = io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(b))
}
