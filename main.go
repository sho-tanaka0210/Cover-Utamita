package main

import (
	"cover-utamita/config"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func main() {
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

	// defer discord.Close()

	// fmt.Println("Listening...")
	// stopBot := make(chan os.Signal, 1)
	// signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	// <-stopBot
	discord.Close()
}
