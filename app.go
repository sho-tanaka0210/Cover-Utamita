package main

import (
	"cover-utamita/consts"
	channelids "cover-utamita/consts/hololive/channel_ids"
	coreDomain "cover-utamita/domain"
	hololiveDomain "cover-utamita/domain/hololive"
	"cover-utamita/infrastructure"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

func App(d coreDomain.DiscordMessenger) error {

	// hololive
	today := time.Now().Format("20060102")
	filename := today + ".json"
	loc, _ := time.LoadLocation("Asia/Tokyo")
	fmt.Println("実行日： ", time.Now().In(loc).AddDate(0, 0, consts.BeforeDay))
	if _, err := os.Stat(filename); err == nil {
	} else if os.IsNotExist(err) {
	} else {
		fmt.Println("ファイルの状態確認に失敗しました:", err)
		return err
	}

	searcher := infrastructure.UtamitaSearcher{ApiKey: os.Getenv("YOUTUBE_API_KEY")}
	r, quotaUsage, err := hololiveDomain.SearchVideos(searcher, channelids.SearchGroups)
	fmt.Printf("YouTube APIクォータ使用量: search.list=%d回, %d単位\n", quotaUsage.SearchListRequests, quotaUsage.Units)
	if err != nil {
		fmt.Printf("YouTubeAPIによる取得に失敗しました。 : %v", err)
		return err
	}

	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		fmt.Println("JSON変換に失敗しました:", err)
		return err
	}
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		fmt.Println("ファイルへの書き込みに失敗しました:", err)
		return err
	}

	var postErrors []error
	for _, v := range r {
		c := infrastructure.Channel{Url: v.Url, DiscordId: v.DiscordId}
		err = c.SendMessage(d)
		if err != nil {
			postErrors = append(postErrors, err)
		}
	}

	return errors.Join(postErrors...)
}
