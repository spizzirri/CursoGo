package service_test

import (
	"github.com/twitteer-go/src/domain"
	"github.com/twitteer-go/src/service"
	"testing"
)

func TestPublishedTweetIsSavedToExternalResource(t *testing.T){

	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter()

	var tweetManager service.TweetManager
	tweetManager.InitializeService()

	var tweetGenerator domain.TextTweet
	tweet := tweetGenerator.NewTweet("damian", "holaMundo", "", nil)
	id, _ := tweetManager.PublishTweet(tweet)
	tweetWriter.Write(tweet)

	memoryWriter := (tweetWriter).(*service.MemoryTweetWrite)
	savedTweet := memoryWriter.GetLastSavedTweet()

	if savedTweet == nil {
		t.Errorf(" El tweet %s es nulo", savedTweet)
	}

	if savedTweet.GetId() != id {
		t.Errorf(" El tweet %s tiene un Id distinto al %d", savedTweet, id )
	}
}