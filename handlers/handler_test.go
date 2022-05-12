package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/ralugr/language-service/logger"
	"github.com/ralugr/language-service/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

type postData struct {
	key   string
	value string
}

var tests = []struct {
	name               string
	url                string
	method             string
	data               []byte
	expectedStatusCode int
}{
	{"home", "/", "GET", []byte{}, http.StatusOK},
	{"get-list", "/list", "GET", []byte{}, http.StatusOK},
	{"subscribe", "/subscribe", "POST", createSubscriber(), http.StatusOK},
	{"set-list", "/list", "POST", createBannedWords(), http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()

	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for tc, tt := range tests {
		logger.Info.Println(" Starting test case ", tc)

		if tt.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + tt.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != tt.expectedStatusCode {
				t.Errorf("for %s expected %d but got %d", tt.name, tt.expectedStatusCode, resp.StatusCode)
			}
		} else {
			resp, err := ts.Client().Post(ts.URL+tt.url, "application/json", bytes.NewReader(tt.data))
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != tt.expectedStatusCode {
				t.Errorf("for %s expected %d but got %d", tt.name, tt.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}

func createSubscriber() []byte {
	subscriber := models.Subscriber{
		Token: "fh343728hdn5dj5sk234ds",
		URL:   "http://localhost:8082/notify",
	}

	data, err := json.Marshal(subscriber)
	if err != nil {
		logger.Info.Printf("Could not marshall subscriber %v", subscriber)
		return nil
	}

	return data
}

func createBannedWords() []byte {
	return []byte("[\"dog\", \"cat\"]")
}
