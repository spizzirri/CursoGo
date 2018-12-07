package service_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/twitteer-go/src/domain"
	"github.com/twitteer-go/src/service"
	"testing"
)

func TestPublishedTweetIsSaved(t *testing.T) {
	var tweetManager service.TweetManager
	tweetManager.InitializeService()
	assert := assert.New(t)
	var tweet domain.Tweet
	var tweetGenerator *domain.TextTweet

	user := "Meli"
	text := "Este es un tweet"
	tweet = tweetGenerator.NewTweet(user, text, "")
	var idTweet int
	idTweet, _ = tweetManager.PublishTweet(tweet)

	publishedTweet := tweetManager.GetTweetById(idTweet)
	if publishedTweet.GetUser() != user &&
		publishedTweet.GetText() != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s", user, text, publishedTweet.GetUser(), publishedTweet.GetText())
	}

	assert.Equal(publishedTweet.GetUser(), user)
	assert.Equal(publishedTweet.GetText(), text)
	assert.NotEqual(idTweet, -1)
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	var tweetManager service.TweetManager
	tweetManager.InitializeService()
	assert := assert.New(t)
	var tweet domain.Tweet
	var tweetGenerator *domain.TextTweet

	var user string
	text := "Este es un tweet"

	tweet = tweetGenerator.NewTweet(user, text, "")

	var err error
	_, err = tweetManager.PublishTweet(tweet)

	assert.NotNil(err, "Expected error should be not nil")
	assert.Equal(err.Error(), "user is required")
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	var tweetManager service.TweetManager
	tweetManager.InitializeService()
	assert := assert.New(t)
	var tweet domain.Tweet
	var tweetGenerator *domain.TextTweet

	user := "Meli"
	var text string

	tweet = tweetGenerator.NewTweet(user, text, "")

	var err error
	_, err = tweetManager.PublishTweet(tweet)

	assert.NotNil(err, "Expected error should be not nil")
	assert.Equal(err.Error(), "text is required")
}

func TestTweetWith140CharacterIsNotPublished(t *testing.T) {
	var tweetManager service.TweetManager
	tweetManager.InitializeService()
	assert := assert.New(t)
	var tweet domain.Tweet
	var tweetGenerator *domain.TextTweet

	user := "Meli"
	text := "BV8D8UBv8wgnNBio4fmBBAQBPyAzf0um3tWNUkYcUmnrYGIlJyoHxms3se5nbm1tTfEof0inyPaEJVUrr5EbNHlYXurKYZi0M2fxNofI1OirYVJyJKk7pzwF68rXGxrgziwxvG67jZgz1"

	tweet = tweetGenerator.NewTweet(user, text, "")

	var err error
	_, err = tweetManager.PublishTweet(tweet)

	assert.NotNil(err, "Expected error should be not nil")
	assert.Equal(err.Error(), "character exceeded")
}


func TestCanPublisheAndRetrieveMoreThanOneTweet(t *testing.T) {
	var tweetManager service.TweetManager
	tweetManager.InitializeService()
	assert := assert.New(t)
	tweetManager.InitializeService()
	var firstTweet, secondTweet domain.Tweet
	var tweetGenerator *domain.TextTweet

	firstTweet = tweetGenerator.NewTweet("Damian", "Hola Mundo", "")
	secondTweet = tweetGenerator.NewTweet("Damian", "Hola Mundo2", "")

	var id1, id2 int

	id1, _ = tweetManager.PublishTweet(firstTweet)
	id2, _ = tweetManager.PublishTweet(secondTweet)

	publishedTweets1 := tweetManager.GetTweetById(id1)
	publishedTweets2 := tweetManager.GetTweetById(id2)

	assert.Equal(publishedTweets1.GetId() + 1, publishedTweets2.GetId())

}

func TestCanRetrieveTheTextTweetsSentByAnUser(t *testing.T) {
	var tweetManager service.TweetManager
	tweetManager.InitializeService()
	assert := assert.New(t)
	// Initialization
	tweetManager.InitializeService()
	var tweet, secondTweet, thirdTweet domain.Tweet
	var tweetGenerator *domain.TextTweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = tweetGenerator.NewTweet(user, text, "")
	secondTweet = tweetGenerator.NewTweet(user, secondText, "")
	thirdTweet = tweetGenerator.NewTweet(anotherUser, text, "")
	// publish the 3 tweets

	_,_  = tweetManager.PublishTweet(tweet)
	_,_  = tweetManager.PublishTweet(secondTweet)
	_,_  = tweetManager.PublishTweet(thirdTweet)

	// Operation
	tweets := tweetManager.GetTweetsByUser(user)

	// Validation
	if len(tweets) != 2 {
		/* handle error */
		panic("Total tweets is not correct")
	}

	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]
	// check if isValidTweet for firstPublishedTweet and secondPublishedTweet
	assert.Equal(firstPublishedTweet.GetUser(), secondPublishedTweet.GetUser())
}

