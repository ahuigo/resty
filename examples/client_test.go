package examples

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-resty/resty/v2"
)

type Payload struct {
	Status int `json:"status"`
}

func TestRestyClient(t *testing.T) {
	// Mock server returns HTTP status specified in request body.
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "error reading request body: %s", err)
			return
		}

		var payload *Payload
		if err := json.Unmarshal(b, &payload); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "error deserializing request body: %s, body=%s", err, string(b))
			return
		}

		println("request body: "+ string(b)+", status: "+ strconv.Itoa(payload.Status))
		w.WriteHeader(payload.Status)
	}))

	client := resty.New().
		AddRetryCondition(
			func(r *resty.Response, err error) bool {
				return err != nil || r.StatusCode() > 499
			},
		).
		SetRetryCount(1)

	// This issue only occurs intermittently in a nondeterministic manner, so I run it a few times
	for i := 0; i < 10; i++ {
		t.Run("RequestBodyIsWrittenOnce", func(t *testing.T) {
			// Trigger some retries
			resp, err := client.
				R().
				SetBody(Payload{http.StatusInternalServerError}).
				Execute(http.MethodPost, srv.URL)
			assertNil(t, err)
			assertEqual(t, http.StatusInternalServerError, resp.StatusCode())
			assertEqual(t, "", string(resp.Body()))

			// Expect 200 OK
			resp, err = client.
				R().
				SetBody(Payload{http.StatusOK}).
				Execute(http.MethodPost, srv.URL)
			assertNil(t, err)
			assertEqual(t, http.StatusOK, resp.StatusCode())
		})
	}
}