package cmd

import (
	"fmt"
	"log"
	"openai-line-bot/controller/mybot"
	"openai-line-bot/env"

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

// add command
func init() {
	rootCmd.AddCommand(serverCmd)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")

	}
}

func start() {
	fmt.Println("start function")

	env.LineConn()
	env.Gpt3Conn()

	ginServer := gin.New()
	ginServer.SetTrustedProxies(nil)
	ginServer.POST("/callback", mybot.NewStart)
	ginServer.Run(":8833")
}
