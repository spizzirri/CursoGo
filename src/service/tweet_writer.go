package service

import (
	"github.com/twitteer-go/src/domain"
	"os"
)

type TweetWriter interface {
	Write(tweet domain.Tweet)
}

type MemoryTweetWrite struct{
	lastTweet domain.Tweet
}

func NewMemoryTweetWriter() *MemoryTweetWrite{
	return &MemoryTweetWrite{}
}

func (m *MemoryTweetWrite) Write(tweet domain.Tweet){
	m.lastTweet = tweet
}

func (m *MemoryTweetWrite) GetLastSavedTweet() domain.Tweet{
	return m.lastTweet
}

type FileTweetWriter struct {
	file *os.File
}

func (f *FileTweetWriter) NewFileTweetWrite() *FileTweetWriter{
	file, _:= os.OpenFile(
		"tweet.txt",
		os.O_WRONLY | os.O_TRUNC|os.O_CREATE,
		0666,
	)
	writer := new(FileTweetWriter)
	writer.file = file
	return writer
}

func (writer *FileTweetWriter) Write(tweet domain.Tweet){
	go func(){
		if writer.file != nil{
			byteSlice := []byte(tweet.PrintableTweet() + "\n")
			_,_ = writer.file.Write(byteSlice)
		}
	}()
}