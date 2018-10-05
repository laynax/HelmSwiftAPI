package helm

import (
	"net/http"
	"youtab/dashboard/services/assert"

	"bytes"

	"io/ioutil"

	"encoding/json"

	"github.com/kataras/iris/core/errors"
	"gopkg.in/yaml.v2"
)

type chart struct {
	ChartURL string `json:"chart_url"`
}

type update struct {
	chart
	Values interface{} `json:"values"`
}

type Upgrade struct {
	ClusterName string
	NS          string
}

func (m manager) Install(chartName string) error {
	var payload = chart{url + chartName}
	resp := m.request(http.MethodPost, payload)

	_, err := marshal(resp)
	return err
}

func (m manager) Upgrade(chartURL string, values map[string]string) error {
	marshalled, err := yaml.Marshal(values)
	assert.Nil(err)

	inlineMap := string(marshalled)

	payload := update{
		chart{chartURL},
		struct {
			Raw string `json:"raw"`
		}{Raw: inlineMap},
	}
	resp := m.request(http.MethodPut, payload)
	_, err = marshal(resp)
	return err
}

func (m manager) Values() {
	resp := m.request(http.MethodPost, nil)

	_, err := marshal(resp)
	assert.Nil(err)
}

func (m manager) request(method string, body interface{}) []byte {
	client := &http.Client{}
	assert.True(m.Name != "", errors.New("release name cannot be empty"))

	data, err := json.Marshal(body)
	assert.Nil(err)

	req, err := http.NewRequest(method, url+m.Name+`/json`, bytes.NewBuffer(data))
	assert.Nil(err)

	resp, err := client.Do(req)
	assert.Nil(err)
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	assert.Nil(err)

	return respBody
}

func marshal(resp []byte) (Release, error) {
	var payload AutoGenerated
	err := json.Unmarshal(resp, &payload)
	assert.Nil(err)

	if payload.Code != 0 {
		return Release{}, errors.New(payload.Message)
	}

	return payload.Release, nil
}
