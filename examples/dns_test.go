package examples

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
)

func TestDnsSet(t *testing.T) {
	ts := createHttpbinServer(0)
	defer ts.Close()

	client := resty.New()
	httpclient := client.GetClient()

	// 1. set dns
	dns := "8.8.8.8"
	transport, err := client.Transport()
	if err != nil {
		t.Fatal(err)
	}
	transport.DialContext = (&net.Dialer{
		Timeout:   httpclient.Timeout,
		KeepAlive: 30 * time.Second,
		Resolver: &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				d := net.Dialer{
					Timeout: time.Duration(5000) * time.Millisecond,
				}
				return d.DialContext(ctx, "udp", fmt.Sprintf("%s:53", dns))
			},
		},
	}).DialContext

	// 2. send request
	r := client.R()
	resp, err := r.Delete(ts.URL + "/delete")
	if err == nil {
		fmt.Println(resp.String())
	}
}