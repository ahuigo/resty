package examples

import (
	"strings"
	"testing"

	"github.com/go-resty/resty/v2"
)

// Example about sending Authentication header
func TestAuth(t *testing.T) {
	ts := createHttpbinServer(0)
	// ts := createEchoServer()
	defer ts.Close()
	// Test authentication usernae and password
	client := resty.New()
	resp, err := client.R().
		SetBasicAuth("USER", "PASSWORD").
		Get(ts.URL + "/get")

	if err != nil {
		t.Fatal(err)
	}
	curlCmdExecuted := resp.Request.GenerateCurlCommand()

	if !strings.Contains(curlCmdExecuted, "Authorization: Basic ") {
		t.Fatal("bad curl:", curlCmdExecuted)
	}
	if !strings.Contains(string(resp.Body()), "Authorization: Basic ") {
		t.Fatal("bad auth body:\n" + resp.String())
	}
	t.Log(curlCmdExecuted)
}
