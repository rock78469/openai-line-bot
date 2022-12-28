package cmd

import (
	"log"

	"openai-line-bot/clients"
	"openai-line-bot/controller/mybot"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// server command
var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

var PORT string

// add command
func init() {
	rootCmd.AddCommand(serverCmd)
	// 啟動時帶入參數 -p 可輸入自訂port，預設為8833
	serverCmd.Flags().StringVarP(&PORT, "port", "p", "8833", "server port")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .clients file")

	}
}

func start() {
	clients.LineConn()
	clients.Gpt3Conn()

	ginServer := gin.New()
	ginServer.SetTrustedProxies(nil)
	ginServer.POST("/callback", mybot.MessageRespondent)
	ginServer.GET("/", func(r *gin.Context) {
		r.JSONP(200, gin.H{"message": "ai bot ready", "code": 0})
	})
	ginServer.Run(":" + PORT)
}
