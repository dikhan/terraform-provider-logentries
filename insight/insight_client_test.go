package insight

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func getTestClient(requestMatcher TestRequestMatcher) *InsightClient {
	testClientServer := TestClientServer{
		RequestMatcher: requestMatcher,
	}
	httpClient, httpServer := testClientServer.TestClientServer()
	c := &InsightClient{InsightUrl: httpServer.URL, ApiKey: "apikey", HttpClient: httpClient}
	return c
}

func TestInsightClient_NewInsightClient(t *testing.T) {
	_, err := NewInsightClient("apiKey", "eu")
	assert.Nil(t, err)
}

func TestInsightClient_NewInsightClientMissing(t *testing.T) {
	_, err := NewInsightClient("", "eu")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "ApiKey is mandatory to initialize Insight client")
	_, err = NewInsightClient("apiKey", "")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Region is mandatory to initialize Insight client")
}

type mockObject struct {
	Data string `json:"data"`
}

func TestInsightClient_ClientGet(t *testing.T) {
	mockResponse := &mockObject{Data: "some data..."}
	requestMatcher := NewRequestMatcher(http.MethodGet, "/api/testing", nil, http.StatusOK, mockResponse)
	c := getTestClient(requestMatcher)
	expectedResponse := &mockObject{}
	err := c.get("/api/testing", expectedResponse)

	assert.Nil(t, err)
	assert.Equal(t, expectedResponse.Data, mockResponse.Data)
}

func TestInsightClient_ClientGetResponseNotStatusOk(t *testing.T) {
	requestMatcher := NewRequestMatcher(http.MethodGet, "/api/testing", nil, http.StatusUnauthorized, &mockObject{})
	c := getTestClient(requestMatcher)
	err := c.get("/api/testing", &mockObject{})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf("Received a non expected response status code %d", http.StatusUnauthorized))
}

func TestInsightClient_ClientPost(t *testing.T) {
	mockRequestPayload := &mockObject{Data: "some req data..."}
	mockResponse := &mockObject{Data: "some data..."}
	requestMatcher := NewRequestMatcher(http.MethodPost, "/api/testing", mockRequestPayload, http.StatusCreated, mockResponse)

	c := getTestClient(requestMatcher)
	expectedResponse := &mockObject{}
	body, err := c.post("/api/testing", mockRequestPayload)
	err = json.Unmarshal(body, &expectedResponse)
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse.Data, mockResponse.Data)
}

func TestInsightClient_ClientPostResponseNotStatusCreated(t *testing.T) {
	requestMatcher := NewRequestMatcher(http.MethodPost, "/api/testing", &mockObject{}, http.StatusUnauthorized, &mockObject{})

	c := getTestClient(requestMatcher)
	_, err := c.post("/api/testing", &mockObject{})

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf("Received a non expected response status code %d", http.StatusUnauthorized))
}

func TestInsightClient_ClientPut(t *testing.T) {
	mockRequestPayload := &mockObject{Data: "some req data..."}
	mockResponse := &mockObject{Data: "some data..."}
	requestMatcher := NewRequestMatcher(http.MethodPut, "/api/testing", mockRequestPayload, http.StatusOK, mockResponse)

	c := getTestClient(requestMatcher)
	expectedResponse := &mockObject{}
	body, err := c.put("/api/testing", mockRequestPayload)
	err = json.Unmarshal(body, &expectedResponse)
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse.Data, mockResponse.Data)
}

func TestInsightClient_ClientPutResponseNotStatusCreated(t *testing.T) {
	requestMatcher := NewRequestMatcher(http.MethodPut, "/api/testing", &mockObject{}, http.StatusUnauthorized, &mockObject{})

	c := getTestClient(requestMatcher)
	_, err := c.put("/api/testing", &mockObject{})

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf("Received a non expected response status code %d", http.StatusUnauthorized))
}

func TestInsightClient_ClientDelete(t *testing.T) {
	requestMatcher := NewRequestMatcher(http.MethodDelete, "/api/testing", nil, http.StatusNoContent, nil)
	c := getTestClient(requestMatcher)
	err := c.delete("/api/testing")
	assert.Nil(t, err)
}

func TestInsightClient_ClientGetResponseNotStatusNoContent(t *testing.T) {
	requestMatcher := NewRequestMatcher(http.MethodDelete, "/api/testing", nil, http.StatusUnauthorized, nil)
	c := getTestClient(requestMatcher)
	err := c.delete("/api/testing")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf("Received a non expected response status code %d", http.StatusUnauthorized))
}
