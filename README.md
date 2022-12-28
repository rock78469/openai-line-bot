# openai-line-bot
## 這是一個串接LineBot以及OpenAI的服務

### Env Prepare
+ Line Bot 相關參數
  * [申請相關連結](https://developers.line.biz/zh-hant/)
    - Line bot secret
    - Line bot token

+ OpenAI 相關參數
  * [申請相關連結](https://beta.openai.com/account/api-keys)
    - openai token
+ Ngrok 相關參數
  * [申請相關連結](https://dashboard.ngrok.com/get-started/setup)
    - ngrok authtoken

### Quick Start
```
docker-compose up -d --build
```
### Start Communicate with ai bot
1. 到ngrok 查看URL
2. 將ngrok產生的URL更新到 LineBot webhook
3. 開始使用@bot 與機器人對話


