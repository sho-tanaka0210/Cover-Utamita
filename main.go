package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	// もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		fmt.Printf("環境変数の読み込みが出来ませんでした: %v", err)
	}

	botToken := os.Getenv("BOT_TOKEN")
	discord, err := discordgo.New(botToken)
	if err != nil {
		fmt.Printf("BOTのログインに失敗しました: %v", err)
	}

	err = discord.Open()
	if err != nil {
		fmt.Printf("疎通に失敗しました。: %v", err)
	}

	defer discord.Close()

	fmt.Println("Listening...")
	stopBot := make(chan os.Signal, 1)
	signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stopBot
}
