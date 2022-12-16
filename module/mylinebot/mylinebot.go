package mylinebot

import (
	"context"
	"fmt"
	"log"
	"openai-line-bot/env"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/line/line-bot-sdk-go/linebot"
)

func LineBotTemplate(events []*linebot.Event) {
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				// quota, err := env.MyLineBot.GetMessageQuota().Do()
				_, err := env.MyLineBot.GetMessageQuota().Do()
				if err != nil {
					log.Println("Quota err:", err)
				}
				// 將訊息丟給OpenAI
				openAIresp := requestOpenAI(message.Text)
				// message.ID: Msg unique ID
				// message.Text: Msg text
				// 官方預設回覆
				// if _, err := env.MyLineBot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("msg ID:"+message.ID+":"+"Get:"+message.Text+" , \n OK! remain message:"+strconv.FormatInt(quota.Value, 10))).Do(); err != nil {
				// 	log.Print(err)
				// }
				// 回覆同樣的訊息
				// if _, err = env.MyLineBot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
				// 	log.Print(err)
				// }
				// 回覆OpenAI回應的訊息
				if _, err = env.MyLineBot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(openAIresp)).Do(); err != nil {
					log.Print(err)
				}
			case *linebot.StickerMessage:
				var kw string
				for _, k := range message.Keywords {
					kw = kw + "," + k
				}

				outStickerResult := fmt.Sprintf("收到貼圖訊息: %s, pkg: %s kw: %s", message.StickerID, message.PackageID, kw)
				if _, err := env.MyLineBot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(outStickerResult)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}

func requestOpenAI(line_Message string) string {
	ctx := context.Background()
	resp, err := env.MyOpenAI.CompletionWithEngine(ctx, "text-davinci-003", gpt3.CompletionRequest{
		Prompt:    []string{line_Message},
		MaxTokens: gpt3.IntPtr(150),
		Stop:      []string{"."},
		Echo:      false,
	})
	if err != nil {
		log.Fatalln(err)
	}
	return resp.Choices[0].Text
}
