package spell

import (
	"fmt"
	"gamma-rho-bot/bing"
	"log"
	"strings"
)

type corrector struct {
	spellCheckerAPIClient bing.SpellCheckAPIClient
	error                 chan error
}

func (c *corrector) checkAndCorrect(message string) (bool, string) {
	message = strings.TrimSpace(message)
	if message == "" {
		return false, ""
	}

	log.Printf("spell check request started...")
	checkingResult, err := c.spellCheckerAPIClient.Check(message)
	if err != nil {
		c.error <- err
		return false, ""
	}
	log.Print("spell check request finished")

	if len(checkingResult.FlaggedTokens) == 0 {
		return false, ""
	}

	var oldNewStrings []string
	for _, flaggedToken := range checkingResult.FlaggedTokens {
		if len(flaggedToken.Suggestions) == 0 {
			c.error <- fmt.Errorf("there is no suggestions for token \"%s\"", flaggedToken.Token)
			return false, ""
		}

		suggestion := fmt.Sprintf("_%s_", flaggedToken.Suggestions[0].Suggestion)
		oldNewStrings = append(oldNewStrings, flaggedToken.Token)
		oldNewStrings = append(oldNewStrings, suggestion)
	}
	replacer := strings.NewReplacer(oldNewStrings...)

	message = replacer.Replace(message)

	return true, message
}
