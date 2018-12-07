package domain

import (
	"time"
)

type Stringer interface {
	String() string
}

type Tweet interface {
	NewTweet(string, string, string, Tweet) Tweet
	PrintableTweet() string
	GetUser() string
	GetId() int
	GetDate() *time.Time
	GetText() string
	SetUser(string)
	SetId(int)
	SetDate(*time.Time)
	SetText(string)
	String() string
}

type headerTweet struct {
	Id int
	User string
	Date *time.Time
}

type TextTweet struct{
	headerTweet
	Text string
}

type ImageTweet struct{
	TextTweet
	UrlImage string
}

type QuoteTweet struct {
	TextTweet
	headerTweet
	TextoSiguiente string
}

func (tweet *TextTweet) NewTweet(user, text, urlImage string, quoteTweet Tweet) Tweet{
	timeNow := time.Now().Local()
	var newTweet Tweet = &TextTweet{ headerTweet{-1, user, &timeNow,}, text }
	return newTweet
}

func (tweet *ImageTweet) NewTweet(user, text, urlImage string, quoteTweet Tweet) Tweet{
	timeNow := time.Now().Local()
	var newTweet Tweet = &ImageTweet{ TextTweet{headerTweet{-1, user, &timeNow,}, text,}, urlImage,}
	return newTweet
}

func (tweet *QuoteTweet) NewTweet(user, text, urlImage string, quoteTweet Tweet) Tweet{
	timeNow := time.Now().Local()
	var newTweet Tweet = &QuoteTweet{ TextTweet{headerTweet{-1, user, &timeNow,}, text,}, headerTweet{-1, user, &timeNow,}, quoteTweet.PrintableTweet(),}
	return newTweet
}


func (tweet *TextTweet) PrintableTweet() string{
	return "@" + tweet.User + ": " + tweet.Text
}

func (tweet *ImageTweet) PrintableTweet() string{
	return "@" + tweet.User + ": " + tweet.Text + "\n" + tweet.UrlImage
}

func (tweet *QuoteTweet) PrintableTweet() string{
	return "@" + tweet.User + ": " + tweet.Text + " \"" + tweet.TextoSiguiente + "\""
}

func (tweet *TextTweet) GetUser() string{
	return tweet.User
}

func (tweet *TextTweet) GetDate() *time.Time{
	return tweet.Date
}

func (tweet *TextTweet) GetId() int{
	return tweet.Id
}

func (tweet *TextTweet) GetText() string{
	return tweet.Text
}

func (tweet *TextTweet) SetUser(user string) {
	tweet.User = user
}

func (tweet *TextTweet) SetDate(date *time.Time){
	tweet.Date = date
}

func (tweet *TextTweet) SetId(id int){
	tweet.Id = id
}

func (tweet *TextTweet) SetText(text string){
	tweet.Text = text
}

func (tweet *TextTweet) String() string{
	return tweet.PrintableTweet()
}


func (tweet *ImageTweet) String() string{
	return tweet.PrintableTweet()
}

func (tweet *QuoteTweet) String() string{
	return tweet.PrintableTweet()
}
