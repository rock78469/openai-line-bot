package mylinebot

import (
	"context"
	"log"
	"openai-line-bot/clients"
	gpt3 "openai-line-bot/clients/gp3"

	"github.com/line/line-bot-sdk-go/linebot"
)

// CheckLineQuota - 檢查使用line是否達到限制
func CheckLineQuota() error {
	_, err := clients.MyLineBot.GetMessageQuota().Do()
	if err != nil {
		return err
	}
	return nil
}

// LineImageReply - 回傳圖片至Line Bot
func LineImageReply(token, response string, err error) {
	// 如果收到的訊息包含錯誤，則將錯誤訊息回傳到LineBot
	if err != nil {
		if _, err2 := clients.MyLineBot.ReplyMessage(token, linebot.NewTextMessage(err.Error())).Do(); err != nil {
			log.Print("LineBot reply message fail", err2)
		}
	}

	// 回傳圖片至Line Bot
	if _, err = clients.MyLineBot.ReplyMessage(token, linebot.NewImageMessage(response, response)).Do(); err != nil {
		log.Print("LineBot reply image fail", err)
	}
}

// LineImageReply - 回傳訊息至Line Bot
func LineMessageReply(token, response string, err error) {
	if err != nil {
		// 如果收到的訊息包含錯誤，則將錯誤訊息回傳到LineBot
		if _, err2 := clients.MyLineBot.ReplyMessage(token, linebot.NewTextMessage(err.Error())).Do(); err != nil {
			log.Print("LineBot reply message fail ", err2)
		}
	}
	// 回傳文字訊息至Line Bot
	if _, err = clients.MyLineBot.ReplyMessage(token, linebot.NewTextMessage(response)).Do(); err != nil {
		log.Print("LineBot reply message fail ", err)
	}
}

func RequestOpenAI(line_Message string) (string, error) {
	ctx := context.Background()

	resp, err := clients.MyOpenAI.CompletionWithEngine(ctx, "text-davinci-003", gpt3.CompletionRequest{
		Prompt:    []string{line_Message},
		MaxTokens: gpt3.IntPtr(512),
		//Stop:      []string{"."},
		//Echo:      false,
	})

	if err != nil {
		log.Fatalln(err)
	}
	return resp.Choices[0].Text, err
}

func RequestImageFromOpenAI(line_Message string) (string, error) {
	ctx := context.Background()

	request := gpt3.ImageRequest{
		Prompt: line_Message,
		Number: 1,
		Size:   "512x512",
	}
	resp, err := clients.MyOpenAI.Image(ctx, request)

	if err != nil {
		log.Fatalln(err)
	}
	return resp.Data[0].Url, err

}
