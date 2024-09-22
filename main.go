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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		config.LoadEnv()

		botToken := os.Getenv("BOT_TOKEN")
		fmt.Println("BOTを実行します。")
		discord, err := discordgo.New("Bot " + botToken)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "BOTのログインに失敗しました: %v", err)
			discord.Close()
			return
		}

		err = discord.Open()
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "疎通に失敗しました。: %v", err)
			discord.Close()
			return
		}

		err = App(discord)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "YouTubeAPIによる取得、もしくはDiscordへの投稿に失敗しました。 : %v", err)
			discord.Close()
			return
		}

		discord.Close()

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "BOTを終了します。")
	})

}
