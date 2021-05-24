package auth

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
)

func TestHealthCheckHandler(t *testing.T) {

	r := chi.NewRouter()
	r.Use(Middleware())

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("bye"))
	})

	ts := httptest.NewServer(r)

	defer ts.Close()

	// GET /
	if _, body := testRequest(t, ts, "GET", "/", nil); body != "bye" {
		t.Fatalf(body)
	}

	if resp, body := testRequest2(t, ts, "GET", "/", nil); resp.StatusCode != http.StatusForbidden {
		t.Fatalf(body)
	}

}

func testRequest(t *testing.T, ts *httptest.Server, method, path string, body io.Reader) (*http.Response, string) {
	req, err := http.NewRequest(method, ts.URL+path, body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	req.Header.Add("Authorization", "Basic c2VjdXJld29ya3M6c3VwZXJzZWNyZXQ=")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}
	defer resp.Body.Close()

	return resp, string(respBody)
}

func testRequest2(t *testing.T, ts *httptest.Server, method, path string, body io.Reader) (*http.Response, string) {
	req, err := http.NewRequest(method, ts.URL+path, body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}
	defer resp.Body.Close()

	return resp, string(respBody)
}
