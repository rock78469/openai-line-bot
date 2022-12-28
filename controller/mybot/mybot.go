package mybot

import (
	"fmt"
	"log"
	"openai-line-bot/clients"
	"openai-line-bot/module/mylinebot"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

// MessageRespondent - 接收Line傳進來的訊息，並在之後進行回應處理
func MessageRespondent(r *gin.Context) {
	// 解析收到的訊息事件events
	events, err := clients.MyLineBot.ParseRequest(r.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			// w.WriteHeader(400)
			fmt.Println(400)
		} else {
			// w.WriteHeader(500)
			fmt.Println(500)
		}
		return
	}

	// 分類收到的訊息事件
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			// 傳入的訊息為文字時，則執行下列動作
			case *linebot.TextMessage:
				// 檢查line 是否達到使用限制
				if mylinebot.CheckLineQuota() != nil {
					log.Println("Quota err:", err)
				}
				// 印出收到的訊息
				log.Println("receive message :: ", message.Text)

				// 如果機器人沒有被tag，則會略過訊息，因為將LineBot加入群組中，如果沒有tag機制，機器人會一直回應訊息，因此需要加上tag才會處發回應功能
				if !strings.Contains(message.Text, "@bot") {
					log.Println("dont container @bot , just return")
					continue
				}

				// 將訊息中有 圖、@botimg、照片 的內容，加入isImage變數內
				isImage := strings.Contains(message.Text, "圖") ||
					strings.Contains(message.Text, "@botimg") ||
					strings.Contains(message.Text, "照片")

				// 如果在群組中被Tag，則將Tag的前綴給取代掉為空值
				txt := strings.Replace(message.Text, "@botimg", "", 1)
				txt = strings.Replace(message.Text, "@bot", "", 1)

				// 如果有關圖片的訊息，則將訊息丟至RequestImageFromOpenAI，讓其回應照片訊息，其他一率回應文字訊息
				if isImage {
					log.Println("send image ...")
					aiReplyImage, err := mylinebot.RequestImageFromOpenAI(txt)
					if err != nil {
						// 回傳收到的錯誤訊息到 Line Bot
						mylinebot.LineImageReply(event.ReplyToken, "", err)
					}
					// 回傳收到的訊息到 Line Bot
					mylinebot.LineImageReply(event.ReplyToken, aiReplyImage, nil)
				} else {
					log.Println("send message ...")
					// 回傳文字訊息至Line Bot
					aiReplyMsg, err := mylinebot.RequestOpenAI(txt)
					if err != nil {
						// 回傳收到的錯誤訊息到 Line Bot
						mylinebot.LineMessageReply(event.ReplyToken, "", err)
					}
					// 回傳收到的訊息到 Line Bot
					mylinebot.LineMessageReply(event.ReplyToken, aiReplyMsg, nil)
				}
			}
		}
	}
}
