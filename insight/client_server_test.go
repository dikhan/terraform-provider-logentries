// Package testutils provides utilities for testing purposes
package insight

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
)

// TestClientServer represents a mock server configured with a matcher and a working client that is able to
// send requests to the server
type TestClientServer struct {
	RequestMatcher TestRequestMatcher
}

// serveResponse checks if the incoming request matches the expected request and if so returns
// the mock response along with the status code configured in RequestMatcher
func (t *TestClientServer) serveResponse(w http.ResponseWriter, r *http.Request) error {
	err := t.RequestMatcher.match(r)
	if err != nil {
		return err
	} else {
		var resp []byte
		var contentType string
		if t.RequestMatcher.Response.RawContent {
			resp = t.RequestMatcher.Response.Payload.([]byte)
			contentType = "application/octet-stream"
		} else {
			resp, err = json.Marshal(t.RequestMatcher.Response.Payload)
			if err != nil {
				return fmt.Errorf("error thrown while marshalling the mock repsonse [%s] - Error: %s", resp, err)
			}
			contentType = "application/json"
		}
		if resp != nil {
			w.WriteHeader(t.RequestMatcher.Response.HttpStatusCode)
			w.Header().Set("Content-Type", contentType)
			w.Write([]byte(resp))
		}
		return nil
	}
}

// TestClientServer creates a client and a server that the user can then use in unit tests. The server returned
// will serve a response only if the incoming request matches the configured expected request.
func (t *TestClientServer) TestClientServer() (*http.Client, *httptest.Server) {

	httpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if err := t.serveResponse(w, r); err != nil {
			fmt.Println("\nAn unexpected error occurred => " + err.Error())
		}
	}))

	httpClient := &http.Client{
		Transport: &http.Transport{
			Proxy: func(request *http.Request) (*url.URL, error) {
				return url.Parse(httpServer.URL)
			},
		},
	}

	return httpClient, httpServer
}
