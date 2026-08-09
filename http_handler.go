package main

import (
	"context"
	"cover-utamita/consts"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type botRunner func(context.Context) error

func newHTTPHandler(guard *runGuard, run botRunner, now func() time.Time) http.Handler {
	if now == nil {
		now = time.Now
	}

	mux := http.NewServeMux()
	// ヘルスチェックでは、外部APIやDiscordに接続しない。
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", http.MethodGet)
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, "ok")
	})
	// 実行エンドポイントは、Cloud Run IAMで検証されるBearerトークンを必須とする。
	mux.HandleFunc("/run", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodPost)
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if !hasBearerToken(r.Header.Get("Authorization")) {
			w.Header().Set("WWW-Authenticate", "Bearer")
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		targetDate := now().In(jst()).AddDate(0, 0, consts.BeforeDay).Format("2006-01-02")
		executed, err := guard.Run(targetDate, func() error {
			return run(r.Context())
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if !executed {
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprintf(w, "target date %s was already completed", targetDate)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintf(w, "target date %s completed", targetDate)
	})

	return mux
}

func hasBearerToken(header string) bool {
	scheme, token, found := strings.Cut(header, " ")
	return found && strings.EqualFold(scheme, "Bearer") && strings.TrimSpace(token) != ""
}

func jst() *time.Location {
	location, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return time.FixedZone("JST", 9*60*60)
	}
	return location
}
