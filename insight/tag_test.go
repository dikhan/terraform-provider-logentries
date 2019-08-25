package insight

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestTags_GetTags(t *testing.T) {

	expectedTags := []*Tag{
		{
			Id:   "tag-uuid",
			Name: "Login Failure",
			Type: "Alert",
			Sources: []*Source{
				{
					Id:              "source-uuid",
					Name:            "auth.log",
					RetentionPeriod: "default",
					StoredDays:      []int{},
				},
			},
			Actions: []*Action{
				{
					Id:               "action-uuid",
					MinMatchesCount:  1,
					MinReportCount:   1,
					MinMatchesPeriod: "Day",
					MinReportPeriod:  "Day",
					Targets: []*Target{
						{
							Id:   "",
							Type: "",
							ParameterSet: &TargetParameterSet{
								Direct: "user@example.com",
								Teams:  "some-team",
								Users:  "user@example.com",
							},
							AlertContentSet: &TargetAlertContentSet{},
						},
					},
					Enabled: true,
					Type:    "Alert",
				},
			},
			Labels: []*Label{
				{
					Id:       "label-uuid",
					Name:     "Login Failure",
					Reserved: false,
					Color:    "007afb",
					SN:       1056,
				},
			},
			Patterns: []string{"Power Button as"},
		},
	}

	requestMatcher := NewRequestMatcher(http.MethodGet, "/management/tags", nil, http.StatusOK, Tags{expectedTags})
	client := getTestClient(requestMatcher)
	returnedTags, err := client.GetTags()
	assert.Nil(t, err)
	assert.EqualValues(t, expectedTags, returnedTags)
}

func TestTags_GetTag(t *testing.T) {

	expectedTag := &Tag{
		Id:   "tag-uuid",
		Name: "Login Failure",
		Type: "Alert",
		Sources: []*Source{
			{
				Id:              "source-uuid",
				Name:            "auth.log",
				RetentionPeriod: "default",
				StoredDays:      []int{},
			},
		},
		Actions: []*Action{
			{
				Id:               "action-uuid",
				MinMatchesCount:  1,
				MinReportCount:   1,
				MinMatchesPeriod: "Day",
				MinReportPeriod:  "Day",
				Targets: []*Target{
					{
						Id:   "",
						Type: "",
						ParameterSet: &TargetParameterSet{
							Direct: "user@example.com",
							Teams:  "some-team",
							Users:  "user@example.com",
						},
						AlertContentSet: &TargetAlertContentSet{},
					},
				},
				Enabled: true,
				Type:    "Alert",
			},
		},
		Labels: []*Label{
			{
				Id:       "label-uuid",
				Name:     "Login Failure",
				Reserved: false,
				Color:    "007afb",
				SN:       1056,
			},
		},
		Patterns: []string{"Power Button as"},
	}

	url := fmt.Sprintf("/management/tags/%s", expectedTag.Id)
	requestMatcher := NewRequestMatcher(http.MethodGet, url, nil, http.StatusOK, TagRequest{expectedTag})
	client := getTestClient(requestMatcher)
	returnedTag, err := client.GetTag(expectedTag.Id)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedTag, returnedTag)
}

func TestTags_GetTagErrorsIfTagIdIsEmpty(t *testing.T) {
	requestMatcher := NewRequestMatcher(http.MethodGet, "/management/tags/", nil, http.StatusOK, TagRequest{&Tag{}})
	client := getTestClient(requestMatcher)
	_, err := client.GetTag("")
	assert.NotNil(t, err)
	assert.Error(t, err, "tagId input parameter is mandatory")
}

func TestTags_PostTag(t *testing.T) {

	p := &Tag{
		Name: "Foo Bar Tag",
		Type: "Alert",
		Sources: []*Source{
			{
				Id: "source-uuid",
			},
		},
		Actions: []*Action{
			{
				MinMatchesCount:  1,
				MinReportCount:   1,
				MinMatchesPeriod: "Day",
				MinReportPeriod:  "Day",
				Targets: []*Target{
					{
						Type: "mailto",
						ParameterSet: &TargetParameterSet{
							Direct: "test@test.com",
						},
						AlertContentSet: &TargetAlertContentSet{
							Context: StringBool(true),
						},
					},
				},
				Enabled: true,
				Type:    "Alert",
			},
		},
		Labels: []*Label{
			{
				Id:       "label-uuid",
				Name:     "Login Failure",
				Reserved: false,
				Color:    "007afb",
				SN:       1056,
			},
		},
		Patterns: []string{"/Foo Bar/"},
	}

	expectedTag := &Tag{
		Id:   "new-tag-uuid",
		Name: p.Name,
		Type: p.Type,
		Sources: []*Source{
			{
				Id:              p.Sources[0].Id,
				Name:            "auth.log",
				RetentionPeriod: "default",
				StoredDays:      []int{},
			},
		},
		Actions: []*Action{
			{
				Id:               "new-action-uuid",
				MinMatchesCount:  p.Actions[0].MinMatchesCount,
				MinReportCount:   p.Actions[0].MinReportCount,
				MinMatchesPeriod: p.Actions[0].MinMatchesPeriod,
				MinReportPeriod:  p.Actions[0].MinReportPeriod,
				Targets: []*Target{
					{
						Id:   "new-target-uuid",
						Type: p.Actions[0].Targets[0].Type,
						ParameterSet: &TargetParameterSet{
							Direct: p.Actions[0].Targets[0].ParameterSet.Direct,
							Teams:  p.Actions[0].Targets[0].ParameterSet.Teams,
							Users:  p.Actions[0].Targets[0].ParameterSet.Users,
						},
						AlertContentSet: p.Actions[0].Targets[0].AlertContentSet,
					},
				},
				Enabled: p.Actions[0].Enabled,
				Type:    p.Actions[0].Type,
			},
		},
		Labels: []*Label{
			{
				Id:       p.Labels[0].Id,
				Name:     p.Labels[0].Name,
				Reserved: p.Labels[0].Reserved,
				Color:    p.Labels[0].Color,
				SN:       p.Labels[0].SN,
			},
		},
		Patterns: p.Patterns,
	}

	requestMatcher := NewRequestMatcher(http.MethodPost, "/management/tags", TagRequest{p}, http.StatusCreated, TagRequest{expectedTag})
	client := getTestClient(requestMatcher)
	err := client.PostTag(p)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedTag, p)
}

func TestTags_PutTag(t *testing.T) {

	tagId := "tagId"

	putTag := &Tag{
		Id:   tagId,
		Name: "Foo Bar Tag",
		Type: "Alert",
		Sources: []*Source{
			{
				Id: "source-uuid",
			},
		},
		Actions: []*Action{
			{
				MinMatchesCount:  0,
				MinReportCount:   1,
				MinMatchesPeriod: "Hour",
				MinReportPeriod:  "Hour",
				Targets: []*Target{
					{
						Type: "mailto",
						ParameterSet: &TargetParameterSet{
							Direct: "test@test.com",
						},
						AlertContentSet: &TargetAlertContentSet{
							Context: StringBool(true),
						},
					},
				},
				Enabled: true,
				Type:    "Alert",
			},
		},
		Labels: []*Label{
			{
				Id:       "label-uuid",
				Name:     "Test Label",
				Reserved: false,
				Color:    "3498db",
				SN:       1025,
			},
		},
		Patterns: []string{"/Foo Bar/"},
	}

	expectedTag := &Tag{
		Id:   "new-tag-uuid",
		Name: putTag.Name,
		Type: putTag.Type,
		Sources: []*Source{
			{
				Id:              putTag.Sources[0].Id,
				Name:            "auth.log",
				RetentionPeriod: "default",
				StoredDays:      []int{},
			},
		},
		Actions: []*Action{
			{
				Id:               "new-action-uuid",
				MinMatchesCount:  putTag.Actions[0].MinMatchesCount,
				MinReportCount:   putTag.Actions[0].MinReportCount,
				MinMatchesPeriod: putTag.Actions[0].MinMatchesPeriod,
				MinReportPeriod:  putTag.Actions[0].MinReportPeriod,
				Targets: []*Target{
					{
						Id:   "new-target-uuid",
						Type: putTag.Actions[0].Targets[0].Type,
						ParameterSet: &TargetParameterSet{
							Direct: putTag.Actions[0].Targets[0].ParameterSet.Direct,
							Teams:  putTag.Actions[0].Targets[0].ParameterSet.Teams,
							Users:  putTag.Actions[0].Targets[0].ParameterSet.Users,
						},
						AlertContentSet: putTag.Actions[0].Targets[0].AlertContentSet,
					},
				},
				Enabled: putTag.Actions[0].Enabled,
				Type:    putTag.Actions[0].Type,
			},
		},
		Labels: []*Label{
			{
				Id:       putTag.Labels[0].Id,
				Name:     putTag.Labels[0].Name,
				Reserved: putTag.Labels[0].Reserved,
				Color:    putTag.Labels[0].Color,
				SN:       putTag.Labels[0].SN,
			},
		},
		Patterns: putTag.Patterns,
	}

	url := fmt.Sprintf("/management/tags/%s", tagId)
	requestMatcher := NewRequestMatcher(http.MethodPut, url, TagRequest{putTag}, http.StatusOK, TagRequest{expectedTag})
	client := getTestClient(requestMatcher)
	err := client.PutTag(putTag)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedTag, putTag)
}

func TestTags_PutTagErrorsIfTagIdIsEmpty(t *testing.T) {
	requestMatcher := NewRequestMatcher(http.MethodGet, "/management/tags/", nil, http.StatusOK, TagRequest{&Tag{}})
	client := getTestClient(requestMatcher)
	err := client.PutTag(&Tag{})
	assert.NotNil(t, err)
	assert.Error(t, err, "tagId input parameter is mandatory")
}

func TestTags_DeleteTag(t *testing.T) {
	tagId := "tag-uuid"
	url := fmt.Sprintf("/management/tags/%s", tagId)
	requestMatcher := NewRequestMatcher(http.MethodDelete, url, nil, http.StatusNoContent, nil)
	client := getTestClient(requestMatcher)
	err := client.DeleteTag(tagId)
	assert.Nil(t, err)
}

func TestTags_DeleteTagErrorsIfTagIdIsEmpty(t *testing.T) {
	requestMatcher := NewRequestMatcher(http.MethodGet, "/management/tags/", nil, http.StatusOK, TagRequest{&Tag{}})
	client := getTestClient(requestMatcher)
	err := client.DeleteTag("")
	assert.NotNil(t, err)
	assert.Error(t, err, "tagId input parameter is mandatory")
}
