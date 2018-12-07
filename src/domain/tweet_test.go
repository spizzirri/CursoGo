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
	tweet = tweetGenerator.NewTweet("grupoesfera", "This is my tweet", "", nil)

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

	assert.Equal(expectedText, text)
}

func TestImageTweetPrintsUserTextAndImageURL(t *testing.T) {

	var tweetGenerator domain.ImageTweet
	// Initialization
	tweet := tweetGenerator.NewTweet("grupoesfera", "This is my image",
		"http://www.grupoesfera.com.ar/common/img/grupoesfera.png", nil)
	// Operation
	text := tweet.PrintableTweet()
	// Validation
	expectedText := "@grupoesfera: This is my image\nhttp://www.grupoesfera.com.ar/common/img/grupoesfera.png"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}
}


func TestQuoteTweetPrintsUserTextAndQuotedTweet(t *testing.T) {

	var quoteTweetGenerator domain.QuoteTweet
	var tweetGenerator domain.TextTweet
	// Initialization
	quotedTweet := tweetGenerator.NewTweet("grupoesfera", "This is my tweet", "", nil)
	tweet := quoteTweetGenerator.NewTweet("nick", "Awesome", "", quotedTweet)

	text := tweet.PrintableTweet()

	// Validation
	expectedText := "@nick: Awesome \"@grupoesfera: This is my tweet\""
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}
}


func TestCanGetAStringFromATweet(t *testing.T){

	assert := assert.New(t)
	var tweet domain.Tweet
	var tweetGenerator *domain.TextTweet

	tweet = tweetGenerator.NewTweet("damian", "HolaMundo", "", nil)
	text := tweet.String()

	expectedText := "@damian: HolaMundo"
	if text != expectedText{
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

	assert.Equal(expectedText, text)
}
