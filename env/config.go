package env

import (
	"fmt"
	"log"
	"os"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/line/line-bot-sdk-go/linebot"
)

var MyLineBot *linebot.Client

var MyOpenAI gpt3.Client

func LineConn() *linebot.Client {
	var err error
	MyLineBot, err = linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("line connection success")
	return MyLineBot
}

func Gpt3Conn() {
	MyOpenAI = gpt3.NewClient(os.Getenv("OPEN_AI_TOKEN"))
}
