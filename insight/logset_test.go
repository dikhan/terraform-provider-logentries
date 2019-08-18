package insight

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestLogsets_GetLogsets(t *testing.T) {

	expectedLogsets := []*Logset{
		{
			Id:          "log-set-uuid",
			Name:        "MyLogset",
			Description: "some description",
			LogsInfo: []*Info{
				{
					Id:   "logs-info-uuid",
					Name: "MyLog",
					Links: []*Link{
						{
							Href: "https://eu.rest.logs.insight.rapid7.com/management/logs/logs-info-uuid",
							Rel:  "Self",
						},
					},
				},
			},
		},
	}

	requestMatcher := NewRequestMatcher(http.MethodGet, "/management/logsets", nil, http.StatusOK, Logsets{expectedLogsets})
	client := getTestClient(requestMatcher)
	returnedLogsets, err := client.GetLogsets()
	assert.Nil(t, err)
	assert.EqualValues(t, expectedLogsets, returnedLogsets)
}

func TestLogsets_GetLogset(t *testing.T) {

	expectedLogset := &Logset{
		Id:          "log-set-uuid",
		Name:        "MyLogset",
		Description: "some description",
		LogsInfo: []*Info{
			{
				Id:   "logs-info-uuid",
				Name: "Lambda Log",
				Links: []*Link{
					{
						Href: "https://eu.rest.logs.insight.rapid7.com/management/logs/logs-info-uuid",
						Rel:  "Self",
					},
				},
			},
		},
	}

	url := fmt.Sprintf("/management/logsets/%s", expectedLogset.Id)
	requestMatcher := NewRequestMatcher(http.MethodGet, url, nil, http.StatusOK, LogsetRequest{expectedLogset})
	client := getTestClient(requestMatcher)
	returnedLogset, err := client.GetLogset(expectedLogset.Id)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedLogset, returnedLogset)
}

func TestLogsets_GetLogsetErrorsIfLogsetIdIsEmpty(t *testing.T) {
	requestMatcher := NewRequestMatcher(http.MethodGet, "/management/logsets/", nil, http.StatusOK, LogsetRequest{&Logset{}})
	client := getTestClient(requestMatcher)
	_, err := client.GetLogset("")
	assert.NotNil(t, err)
	assert.Error(t, err, "logsetId input parameter is mandatory")
}

func TestLogsets_PostLogset(t *testing.T) {

	p := &Logset{
		Name:        "MyLogset2",
		Description: "some description",
		LogsInfo: []*Info{
			{
				Id: "logs-info-uuid",
			},
		},
		UserData: map[string]string{},
	}

	expectedLogset := &Logset{
		Id:          "log-set-uuid",
		Name:        p.Name,
		Description: p.Description,
		LogsInfo: []*Info{
			{
				Id:   p.LogsInfo[0].Id,
				Name: "mylog",
				Links: []*Link{
					{
						Href: "https://eu.rest.logs.insight.rapid7.com/management/logs/logs-info-uuid",
						Rel:  "Self",
					},
				},
			},
		},
		UserData: map[string]string{},
	}

	requestMatcher := NewRequestMatcher(http.MethodPost, "/management/logsets", LogsetRequest{p}, http.StatusCreated, LogsetRequest{expectedLogset})
	client := getTestClient(requestMatcher)
	err := client.PostLogset(p)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedLogset, p)
}

func TestLogsets_PutLogset(t *testing.T) {

	logsetId := "log-set-uuid"

	p := &Logset{
		Id:          logsetId,
		Name:        "New Name",
		Description: "updated description",
		LogsInfo: []*Info{
			{
				Id:   "logs-info-uuid",
				Name: "Lambda Log",
				Links: []*Link{
					{
						Href: "https://eu.rest.logs.insight.rapid7.com/management/logs/logs-info-uuid",
						Rel:  "Self",
					},
				},
			},
		},
		UserData: map[string]string{},
	}

	expectedLogset := &Logset{
		Id:          logsetId,
		Name:        p.Name,
		Description: p.Description,
		LogsInfo: []*Info{
			{
				Id:   p.LogsInfo[0].Id,
				Name: p.LogsInfo[0].Name,
				Links: []*Link{
					{
						Href: p.LogsInfo[0].Links[0].Href,
						Rel:  p.LogsInfo[0].Links[0].Rel,
					},
				},
			},
		},
		UserData: map[string]string{},
	}

	url := fmt.Sprintf("/management/logsets/%s", logsetId)
	requestMatcher := NewRequestMatcher(http.MethodPut, url, LogsetRequest{p}, http.StatusOK, LogsetRequest{expectedLogset})
	client := getTestClient(requestMatcher)
	err := client.PutLogset(p)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedLogset, p)
}

func TestLogsets_PutLogsetSetErrorsIfLogsetIdIsEmpty(t *testing.T) {
	requestMatcher := NewRequestMatcher(http.MethodGet, "/management/logsets/", nil, http.StatusOK, LogsetRequest{&Logset{}})
	client := getTestClient(requestMatcher)
	err := client.PutLogset(&Logset{})
	assert.NotNil(t, err)
	assert.Error(t, err, "logsetId input parameter is mandatory")
}

func TestLogsets_DeleteLogset(t *testing.T) {
	logSetId := "log-set-uuid"
	url := fmt.Sprintf("/management/logsets/%s", logSetId)
	requestMatcher := NewRequestMatcher(http.MethodDelete, url, nil, http.StatusNoContent, nil)
	client := getTestClient(requestMatcher)
	err := client.DeleteLogset(logSetId)
	assert.Nil(t, err)
}

func TestLogsets_DeleteLogsetSetErrorsIfLogsetIdIsEmpty(t *testing.T) {
	requestMatcher := NewRequestMatcher(http.MethodGet, "/management/logsets/", nil, http.StatusOK, LogsetRequest{&Logset{}})
	client := getTestClient(requestMatcher)
	err := client.DeleteLogset("")
	assert.NotNil(t, err)
	assert.Error(t, err, "logSetId input parameter is mandatory")
}
