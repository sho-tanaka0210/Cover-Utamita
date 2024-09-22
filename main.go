package main

import (
	"cover-utamita/config"
	"fmt"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	http.ListenAndServe(":"+port, nil)

	config.LoadEnv()

	botToken := os.Getenv("BOT_TOKEN")
	fmt.Println("BOTを実行します。")
	discord, err := discordgo.New("Bot " + botToken)
	if err != nil {
		fmt.Printf("BOTのログインに失敗しました: %v", err)
	}

	err = discord.Open()
	if err != nil {
		fmt.Printf("疎通に失敗しました。: %v", err)
	}
	err = App(discord)
	if err != nil {
		fmt.Printf("YouTubeAPIによる取得、もしくはDiscordへの投稿に失敗しました。 : %v", err)
	}

	fmt.Println("BOTを終了します。")
	discord.Close()
}
