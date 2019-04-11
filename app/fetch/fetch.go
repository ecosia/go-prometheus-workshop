package fetch

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
)

// exported so visible to main package
type Response struct {
	Count int     `json:"count"`
	Rate  float32 `json:"rate"`
}

var (
	treesAPIURL = "https://api.ecosia.org/v1/trees/count"
)

// NewRequest creates a http client and makes the request, it returns a response and an error
func NewRequest() (resp *http.Response, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", treesAPIURL, nil)
	if err != nil {
		return resp, err
	}
	resp, err = client.Do(req)
	return resp, err
}

// Fetch returns response from Trees API and stores it in the response variable.
func Fetch(makeRequest func() (*http.Response, error)) (Response, int, error) {
	var treeData Response
	// randomly returns a 500 status
	statusCodes := map[int]int{503: 20}
	number := rand.Intn(100)
	if number < statusCodes[503] {
		return treeData, 503, errors.New("Fake 503(service unavailable) repsonse")
	}

	resp, err := makeRequest()
	defer resp.Body.Close()

	if err != nil {
		return treeData, http.StatusInternalServerError, err
	}

	if resp.StatusCode != 200 {
		return treeData, resp.StatusCode, nil
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return treeData, http.StatusInternalServerError, err
	}

	err = json.Unmarshal(respBytes, &treeData)
	if err != nil {
		return treeData, http.StatusInternalServerError, err
	}
	return treeData, resp.StatusCode, err
}
