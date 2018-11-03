package bing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const apiEndpointUri = "https://api.cognitive.microsoft.com/bing/v7.0/spellcheck?mkt=en-US&mode=proof"

type client struct {
	apiKey string
}

func (c *client) Check(text string) (*CheckResult, error) {
	client := &http.Client{}
	form := url.Values{}
	form.Add("text", text)
	request, err := http.NewRequest(http.MethodPost, apiEndpointUri, strings.NewReader(form.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Ocp-Apim-Subscription-Key", c.apiKey)

	response, err := client.Do(request)

	defer response.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("Bing Spell API error: %s", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Bing Spell API returns %d", response.StatusCode)
	}

	result := &CheckResult{}
	if err = json.NewDecoder(response.Body).Decode(result); err != nil {
		return nil, fmt.Errorf(
			"cann't decode JSON from response of Bing Spell API: %s",
			err.Error(),
		)
	}

	return result, nil
}
