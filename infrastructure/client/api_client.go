package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"taskmanager/domain/model"
)

type apiClient struct {
	client    *http.Client
	Token     string
	ProjectId string
}

type ApiClient interface {
	RequestAllTasks(t []*model.Task) ([]*model.Task, error)
}

func NewApiClient(token string, projectId string) ApiClient {
	httpClient := &http.Client{}
	return &apiClient{
		client:    httpClient,
		Token:     token,
		ProjectId: projectId,
	}
}

func (api *apiClient) RequestAllTasks(t []*model.Task) ([]*model.Task, error) {
	url := "https://api.todoist.com/rest/v1/tasks"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+api.Token)

	query := req.URL.Query()
	query.Add("project_id", api.ProjectId)
	req.URL.RawQuery = query.Encode()

	res, err := api.client.Do(req)
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(responseData, &t)

	return t, nil
}
