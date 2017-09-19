package main

import
// "encoding/json"
// "fmt"

(
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// "math/rand"
// "time"

func TestMeli(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

}
func TestDownloadError(t *testing.T) {
	_, err := Download("args", func(s string) (*http.Response, error) { return nil, errors.New("error") })
	if err == nil {
		t.Error("Download sould have thrown error")
	}
}

func TestDownloadParsingError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		resp, _ := json.Marshal("")
		io.WriteString(w, string(resp))

	}))
	defer ts.Close()
	_, err := Download("args", func(s string) (*http.Response, error) { return http.Get(ts.URL) })
	if err == nil {
		t.Error("Download sould have thrown error")
	}
	ts2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		resp, _ := json.Marshal(make([]int, 200))
		io.WriteString(w, string(resp))

	}))
	defer ts2.Close()
	_, err = Download("args", func(s string) (*http.Response, error) { return http.Get(ts.URL) })
	if err == nil {
		t.Error("Download sould have thrown error")
	}
	ts3 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		resp, _ := json.Marshal("{  \"total_items_in_this_category\": 2222238}")
		io.WriteString(w, string(resp))

	}))
	defer ts3.Close()
	_, err = Download("args", func(s string) (*http.Response, error) { return http.Get(ts.URL) })
	if err == nil {
		t.Error("Download sould have thrown error")
	}
}
func TestDownload(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		resp, _ := json.Marshal(GenerarMismo(400, 20, 30))
		io.WriteString(w, string(resp))
	}))
	defer ts.Close()
	resp, err := Download("args", func(s string) (*http.Response, error) { return http.Get(ts.URL) })
	if err != nil {
		t.Error("Download sould not have thrown error")
	}
	results, ok := resp["results"].([]interface{})
	if !ok {
		t.Error("There was no result!!")
	}
	for i := range results {
		item := results[i].(map[string]interface{})
		price, ok := item["price"].(float64)
		if !ok {
			continue
		}
		if price != 20 {
			t.Error("price should be 20 but got", price)
		}
		sold, ok := item["sold_quantity"].(float64)
		if !ok {
			continue
		}
		if sold != 30 {
			t.Error("sold should be 30 but got", sold)
		}
	}
}
