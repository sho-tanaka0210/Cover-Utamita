package api

import (
	"fmt"
	"io"
	"net/http"
)

func SearchVideoes(url string) (results []string, err error) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		// エラーハンドリング
		fmt.Println("Error:", err)
		return nil, err
	}

	fmt.Println(string(body))
	return results, nil
}
