package insight

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// An ExpectedRequest represents the expected request that will be matched against the incoming http request
type ExpectedRequest struct {
	HttpMethod string
	Url        string
	Payload    interface{}
}

// A Response represents the response to return when the incoming request matches the ExpectedRequest
type Response struct {
	HttpStatusCode int
	Payload        interface{}
	RawContent     bool
}

// A TestRequestMatcher represents the expected behaviour of the mock server
type TestRequestMatcher struct {
	ExpectedRequest ExpectedRequest
	Response        Response
}

// NewRequestMatcher constructs a TestRequestMatcher used to match an http.Request with the expected configured request
// and returns the configured response in JSON and status code
func NewRequestMatcher(expectedHttpMethod, expectedPath string, expectedPayload interface{}, responseStatusCode int, response interface{}) TestRequestMatcher {
	return NewRequestMatcherWithOptions(expectedHttpMethod, expectedPath, expectedPayload, responseStatusCode, response, false)
}

// NewRequestMatcher constructs a TestRequestMatcher used to match an http.Request with the expected configured request
// and return the configured response (either in JSON or RAW depending on the value provided for responseRawContent) and status code
func NewRequestMatcherWithOptions(expectedHttpMethod, expectedPath string, expectedPayload interface{}, responseStatusCode int, response interface{}, responseRawContent bool) TestRequestMatcher {
	return TestRequestMatcher{
		ExpectedRequest: ExpectedRequest{
			HttpMethod: expectedHttpMethod,
			Url:        expectedPath,
			Payload:    expectedPayload,
		},
		Response: Response{responseStatusCode, response, responseRawContent},
	}
}

// match checks whether the http.Request is equal to the expected one - method, url and body must match; otherwise
// an error is returned
func (rm *TestRequestMatcher) match(r *http.Request) error {
	var body []byte
	var err error
	if rm.ExpectedRequest.HttpMethod == r.Method && rm.ExpectedRequest.Url == r.URL.Path {
		body, err = ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}
		if len(body) == 0 {
			return nil
		} else {
			if rm.ExpectedRequest.Payload == nil {
				fmt.Println("Request matcher missing expected request payload, please populate the expected paylaod field")
			}
			expectedRequest, err := json.Marshal(rm.ExpectedRequest.Payload)
			if err != nil {
				return err
			}
			if string(expectedRequest) == string(body) {
				return nil
			}
		}
	}
	return fmt.Errorf("No matching expected request found:\n- ExpectedRequest = %s\n- ActualRequest = %s %s %s\n", rm.ExpectedRequest, r.Method, r.URL.Path, string(body))
}
