package service

import (
	"errors"
	"github.com/twitteer-go/src/domain"
)

type TweetManager struct{
	Tweets []domain.Tweet
	mapTweets map[string][]domain.Tweet
}

func (tm *TweetManager) InitializeService(){
	tm.Tweets = make([]domain.Tweet, 0)
	tm.mapTweets = make(map[string][]domain.Tweet)
}

func (tm *TweetManager) PublishTweet(tweet domain.Tweet) (int, error) {
	var err error

	if tweet.GetUser() == "" {
		return -1, errors.New("user is required")
	}
	if tweet.GetText() == "" {
		return -1, errors.New("text is required")
	}
	if len(tweet.GetText()) > 140 {
		return -1, errors.New("character exceeded")
	}

	tweet.SetId(len(tm.Tweets))
	tm.Tweets = append(tm.Tweets, tweet)

	tm.mapTweets[tweet.GetUser()] = append(tm.mapTweets[tweet.GetUser()], tweet)

	return len(tm.Tweets)-1, err
}

func (tm *TweetManager) GetTweetById(id int) domain.Tweet {
	if tm.Tweets != nil && id < len(tm.Tweets) && id >= 0 {
		return tm.Tweets[id]
	} else{
		return nil}
}

func (tm *TweetManager) GetTotalTweets() int {
	return len(tm.Tweets)
}

func (tm *TweetManager) GetTweetsByUser(user string) []domain.Tweet {
	var elements []domain.Tweet
	var exists bool

	elements, exists = tm.mapTweets[user]

	if exists == true {
		return elements
	}else{
		return nil
	}
}