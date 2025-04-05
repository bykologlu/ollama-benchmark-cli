package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ollama-benchmark/internal/i18n"
)

type ModelInfo struct {
	Name string `json:"name"`
}

type ModelListResponse struct {
	Models []ModelInfo `json:"models"`
}

func GetModelList(apiURL string) ([]string, error) {
	url := fmt.Sprintf("%s/api/tags", apiURL)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf(i18n.T("err_api_request"), err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(i18n.T("err_api_status"), resp.Status)
	}

	var parsed ModelListResponse
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return nil, fmt.Errorf(i18n.T("err_api_parse"), err)
	}

	var modelNames []string
	for _, model := range parsed.Models {
		modelNames = append(modelNames, model.Name)
	}

	return modelNames, nil
}
