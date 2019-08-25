package insight

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestLabels_GetLabels(t *testing.T) {

	expectedLabels := []*Label{
		{
			Id:       "label-uuid",
			Name:     "Login Failure",
			Reserved: false,
			Color:    "007afb",
			SN:       1056,
		},
	}

	requestMatcher := NewRequestMatcher(http.MethodGet, "/management/labels", nil, http.StatusOK, Labels{expectedLabels})
	client := getTestClient(requestMatcher)
	returnedLabels, err := client.GetLabels()
	assert.Nil(t, err)
	assert.EqualValues(t, expectedLabels, returnedLabels)
}

func TestTags_GetLabel(t *testing.T) {

	expectedLabel := &Label{
		Id:       "label-uuid",
		Name:     "Login Failure",
		Reserved: false,
		Color:    "007afb",
		SN:       1056,
	}

	url := fmt.Sprintf("/management/labels/%s", expectedLabel.Id)
	requestMatcher := NewRequestMatcher(http.MethodGet, url, nil, http.StatusOK, LabelRequest{expectedLabel})
	client := getTestClient(requestMatcher)
	returnedLabel, err := client.GetLabel(expectedLabel.Id)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedLabel, returnedLabel)
}

func TestTags_GetLabelErrorsIfTagIdIsEmpty(t *testing.T) {
	requestMatcher := NewRequestMatcher(http.MethodGet, "/management/labels/", nil, http.StatusOK, LabelRequest{&Label{}})
	client := getTestClient(requestMatcher)
	_, err := client.GetLabel("")
	assert.NotNil(t, err)
	assert.Error(t, err, "labelId input parameter is mandatory")
}

func TestLabels_DeleteLabel(t *testing.T) {
	labelId := "log-set-uuid"
	url := fmt.Sprintf("/management/labels/%s", labelId)
	requestMatcher := NewRequestMatcher(http.MethodDelete, url, nil, http.StatusNoContent, nil)
	client := getTestClient(requestMatcher)
	err := client.DeleteLabel(labelId)
	assert.Nil(t, err)
}

func TestLabels_PostLabel(t *testing.T) {

	p := &Label{
		Name:  "My Label",
		Color: "ff0000",
	}

	expectedLabel := &Label{
		Id:       "label-uuid",
		Name:     p.Name,
		Color:    p.Color,
		Reserved: false,
		SN:       1021,
	}

	requestMatcher := NewRequestMatcher(http.MethodPost, "/management/labels", LabelRequest{p}, http.StatusCreated, LabelRequest{expectedLabel})
	client := getTestClient(requestMatcher)
	err := client.PostLabel(p)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedLabel, p)
}
