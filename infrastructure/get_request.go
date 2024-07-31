package infrastructure

import (
	"io"
	"net/http"
)

// 引数で受け取ったエンドポイントに対してGetリクエストを送信し、結果をbyte[]型の変数として返却する。
func Get(url string) ([]byte, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
