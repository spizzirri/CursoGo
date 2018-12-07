package main

import (
	"fmt"
	"github.com/abiosoft/ishell"
	"github.com/twitteer-go/src/domain"
	"github.com/twitteer-go/src/service"
	"strconv"
)

func main() {

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	var tweetManager service.TweetManager

	tweetManager.InitializeService()

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your username: ")

			user := c.ReadLine()

			c.Print("Write your tweet: ")

			text := c.ReadLine()

			var tweet domain.Tweet
			var tweetGenerator *domain.TextTweet

			tweet = tweetGenerator.NewTweet(user, text, "")

			_,_ = tweetManager.PublishTweet(tweet)

			c.Print("Tweet sent\n")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Show a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your Tweet id: ")
			var idTweet int
			idTweet, _= strconv.Atoi(c.ReadLine())
			tweet := tweetManager.GetTweetById(idTweet)

			c.Println("Tweet id: ", tweet.GetId())
			c.Println("Tweet text: ", tweet.GetText())
			c.Println("User account: ", tweet.GetUser())
			c.Println("Tweeted at: ", tweet.GetDate())

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTotalTweets",
		Help: "Show a total Tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			var totalTweets int
			totalTweets = tweetManager.GetTotalTweets()
			c.Println("Total Tweet: ", totalTweets)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showAllUserTweets",
		Help: "Show all tweets by user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			var user string
			var tweets []domain.Tweet

			fmt.Print("Write Twitter username: ")
			user = c.ReadLine()
			tweets = tweetManager.GetTweetsByUser(user)

			if len(tweets) < 1{
				fmt.Println("No hay Tweets de ", user)
			}

			for nroTweet, tweet := range tweets {
				c.Println(tweet.String())
				c.Println("Nro Tweet ", nroTweet)
				c.Println("=====================")
				c.Println("Tweet id: ", tweet.GetId())
				c.Println("Tweet text: ", tweet.GetText())
				c.Println("User account: ", tweet.GetUser())
				c.Println("Tweeted at: ", tweet.GetDate())
				c.Println()
			}
			return
		},
	})

	shell.Run()
}
