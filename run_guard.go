package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type runGuard struct {
	directory string
	mu        sync.Mutex
}

func newRunGuard(directory string) *runGuard {
	return &runGuard{directory: directory}
}

// Runは処理を直列化し、正常終了した対象日を記録する。
// Cloud Run側でも最大インスタンス数と同時実行数を1にするが、同一プロセス内の
// 重複リクエストはこのロックで最終的に防止する。
func (g *runGuard) Run(targetDate string, run func() error) (bool, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if err := os.MkdirAll(g.directory, 0o700); err != nil {
		return false, fmt.Errorf("create run state directory: %w", err)
	}

	marker := filepath.Join(g.directory, targetDate+".done")
	if _, err := os.Stat(marker); err == nil {
		return false, nil
	} else if !errors.Is(err, os.ErrNotExist) {
		return false, fmt.Errorf("read run state: %w", err)
	}

	if err := run(); err != nil {
		return false, err
	}
	if err := os.WriteFile(marker, []byte("completed\n"), 0o600); err != nil {
		return false, fmt.Errorf("write run state: %w", err)
	}

	return true, nil
}
