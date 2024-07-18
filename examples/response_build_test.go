package examples

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty/v2"
)

// Example about building response
func TestResponseBuilder(t *testing.T) {
	var err error
	var data = 1
	responseBytes, _ := json.Marshal(data)

	respRecorder := httptest.NewRecorder()
	respRecorder.Write(responseBytes)

	// build response
	resp := resty.Response{
		// Request: resty.New().R(),
		RawResponse :respRecorder.Result(),
		// body: []byte("abc"),
	}
	// if resp.body, err = io.ReadAll(resp.RawResponse.Body); err != nil {
	// 	t.Fatalf("err:%v", err)
	// }
	ndata, err := io.ReadAll(resp.RawResponse.Body) 
	if err != nil {
		t.Fatalf("err:%v", err)
	}

	if !bytes.Equal(ndata , responseBytes) {
		t.Fatalf("expect response:%v", data)
	}

}
