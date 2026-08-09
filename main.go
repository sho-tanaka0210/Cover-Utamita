package main

import (
	"context"
	"cover-utamita/config"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
)

func main() {
	config.LoadEnv()

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	stateDir := os.Getenv("RUN_STATE_DIR")
	if stateDir == "" {
		stateDir = "/tmp/cover-utamita"
	}

	handler := newHTTPHandler(newRunGuard(stateDir), runBot, nil)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	log.Printf("Server is listening on port %s...", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func runBot(_ context.Context) error {
	botToken := os.Getenv("BOT_TOKEN")
	discord, err := discordgo.New("Bot " + botToken)
	if err != nil {
		return fmt.Errorf("BOTのログインに失敗しました: %w", err)
	}

	if err := discord.Open(); err != nil {
		return fmt.Errorf("Discordとの疎通に失敗しました: %w", err)
	}
	defer discord.Close()

	if err := App(discord); err != nil {
		return fmt.Errorf("YouTube APIによる取得、もしくはDiscordへの投稿に失敗しました: %w", err)
	}

	return nil
}
