package examples

import (
	"strings"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
)

// Example about setting timeout
func TestTimeout(t *testing.T) {
	ts := createHttpbinServer(0)
	// ts := createHttpbinServer(0)
	defer ts.Close()

	req := resty.New().SetTimeout(10 * time.Millisecond).R()
	resp, err := req.Get(ts.URL + "/sleep/15")
	if err == nil {
		t.Fatal("expected timeout error, body:", resp.String())
	}
	assertEqual(t, true, strings.Contains(err.Error(), "Client.Timeout exceeded"))
}
