package main

import (
	"fmt"
	"github.com/abiosoft/ishell"
	"github.com/gin-gonic/gin"
	"github.com/twitteer-go/src/domain"
	"github.com/twitteer-go/src/service"
	"net/http"
	"strconv"
)

func funcionQueHaceGet(c *gin.Context){

	parameter := c.Param("parametro")
	c.String(http.StatusOK, "ok")
	fmt.Println("Parametro GET " + parameter)
}

func funcionQueHacePost(c *gin.Context){
	var user domain.Usuario
	if err := c.ShouldBindJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
	}else{
		c.String(http.StatusOK,"ok")
		fmt.Println("Parametros POST " + user.Correo)
		fmt.Println("Parametros POST " + user.Pass)
	}
	return
}

func tweetsByUser(c *gin.Context){
	var tweets []domain.Tweet

	tm := c.MustGet("tm").(*service.TweetManager)

	tweets = tm.GetTweetsByUser(c.Param("user"))

	c.JSON(http.StatusOK, tweets)
	return
}

func newTweet(c *gin.Context){
	var textTweet domain.UsuarioTextTweet
	var tm *service.TweetManager
	tm = c.MustGet("tm").(*service.TweetManager)

	if err := c.ShouldBindJSON(&textTweet); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
	}else{
		var tweetGenerator *domain.TextTweet
		fmt.Println( "JSON: " +  textTweet.Usuario + ", " +  textTweet.Texto)
		tweet := tweetGenerator.NewTweet(textTweet.Usuario, textTweet.Texto, "", nil)
		id,_ := tm.PublishTweet(tweet)
		c.String(http.StatusOK,"Tweet numero %d", id)
	}
}

func myMiddleware(tm *service.TweetManager) gin.HandlerFunc {
	return func(c *gin.Context){
		c.Set("tm", tm)
		c.Next()
	}
}

func initServer(tm *service.TweetManager){

	router := gin.Default()
	router.Use(myMiddleware(tm))

	router.GET("/unGet/:parametro", funcionQueHaceGet)
	router.POST("/unPost", funcionQueHacePost)
	router.GET("/tweets/:user", tweetsByUser)
	router.POST("/tweet", newTweet)

	router.Run()
}

func main() {

	var tweetManager service.TweetManager
	tweetManager.InitializeService()

	go initServer(&tweetManager)

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTextTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your username: ")

			user := c.ReadLine()

			c.Print("Write your tweet: ")

			text := c.ReadLine()

			var tweet domain.Tweet
			var tweetGenerator *domain.TextTweet

			tweet = tweetGenerator.NewTweet(user, text, "", nil)

			_,_ = tweetManager.PublishTweet(tweet)

			c.Print("Tweet sent\n")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "publishImageTweet",
		Help: "Publishes a image tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your username: ")
			user := c.ReadLine()

			c.Print("Write your tweet: ")
			text := c.ReadLine()

			c.Print("Write your image: ")
			urlImage := c.ReadLine()

			var tweet domain.Tweet
			var tweetGenerator *domain.ImageTweet

			tweet = tweetGenerator.NewTweet(user, text, urlImage, nil)

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
			c.Println(tweet.String())
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
				c.Println("Nro Tweet ", nroTweet)
				c.Println("=====================")
				c.Println(tweet.String())
				c.Println()
			}
			return
		},
	})

	shell.Run()
}
