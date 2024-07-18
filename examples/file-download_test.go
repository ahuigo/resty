package examples

import (
	"path/filepath"
	"testing"

	"github.com/go-resty/resty/v2"
)

/*
An example about post `file` with `form data`:
curl "https://www.httpbin.org/post" -F 'file1=@./test-file.txt'  -F 'name=alex'
*/
func TestPostFile1(t *testing.T) {
	ts := createHttpbinServer(0)
	defer ts.Close()

	download_path := filepath.Join(getTestDataPath(), "tmp/download.txt") // auto create parent directories, overwrite if exists
	r := resty.New().R().
		SetFormData(MapString{ "name": "Alex", }).
		SetOutput(download_path)
		// SetResult(&data)// SetResult and SetOutput are in conflict

	// 2. Post file
	_, err := r.Get(ts.URL + "/post")
	if err != nil {
		t.Fatal(err)
	}


}
