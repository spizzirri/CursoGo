package domain_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/twitteer-go/src/domain"
	"testing"
)

func TestCanGetAPrintableTweet(t *testing.T) {
	assert := assert.New(t)

	var tweet domain.Tweet
	var tweetGenerator *domain.TextTweet

	// Initialization
	tweet = tweetGenerator.NewTweet("grupoesfera", "This is my tweet", "")

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

	assert.Equal(expectedText, text)
}


func TestCanGetAStringFromATweet(t *testing.T){

	assert := assert.New(t)
	var tweet domain.Tweet
	var tweetGenerator *domain.TextTweet

	tweet = tweetGenerator.NewTweet("damian", "HolaMundo", "")
	text := tweet.String()

	expectedText := "@damian: HolaMundo"
	if text != expectedText{
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

	assert.Equal(expectedText, text)
}
