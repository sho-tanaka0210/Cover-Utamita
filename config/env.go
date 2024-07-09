package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

// 環境変数の設定・読み込みを行う
// .envファイルが存在する場合のみ、そこに定義されている値を環境変数として利用する(開発環境のみ)
// それ以外の場合は環境変数としてデプロイする際に設定する。
func LoadEnv() {
	fmt.Println(".envファイルを読み込みます...")
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("開発環境でないため、.envファイルは読み込みません。")
	}
}
