package bing

import "errors"

func NewSpellCheckAPIClient(apiKey string) (SpellCheckAPIClient, error) {
	if apiKey == "" {
		return nil, errors.New("token shouldn't be empty")
	}
	return &client{apiKey: apiKey}, nil
}

type SpellCheckAPIClient interface {
	Check(text string) (*CheckResult, error)
}

type CheckResult struct {
	Type          string         `json:"_type"`
	FlaggedTokens []FlaggedToken `json:"flaggedTokens"`
}

type FlaggedToken struct {
	Offset      int          `json:"offset"`
	Token       string       `json:"token"`
	Type        string       `json:"type"`
	Suggestions []Suggestion `json:"suggestions"`
}

type Suggestion struct {
	Suggestion string  `json:"suggestion"`
	Score      float64 `json:"score"`
}
