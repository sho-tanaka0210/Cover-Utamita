# Cover-Utamita
ホロライブ・ホロスタの歌ってみたをDiscordに通知する

## Build
`docer build . -t utamita-bot`

## ローカル実行方法

`docker run --rm -v $(pwd)/:/code -it utamita-bot /bin/bash`

## Discord Server
[UtamitaBot](https://discord.gg/BcEK9UU7nA)

## HTTP endpoints

- `GET /healthz`: ヘルスチェック。YouTube API や Discord には接続しない。
- `POST /run`: 前日分の検索・投稿を実行する。Cloud Run IAM で認証された Bearer トークンが必要。

検索対象期間は日本時間の前日 00:00 以上、当日 00:00 未満。実行ログの
`YouTube APIクォータ使用量` で、1 回の定期実行における `search.list` の
リクエスト回数とクォータ消費量を確認できる。

Cloud Run は非公開、最大インスタンス数と同時実行数はともに 1 でデプロイする。
デプロイ前に GitHub Actions の Repository variables に、既存の Cloud Scheduler
ジョブ名 `SCHEDULER_JOB` と、`roles/run.invoker` を付与する呼び出し元サービス
アカウント `SCHEDULER_SERVICE_ACCOUNT` を設定する。デプロイ時に Scheduler の対象を
`POST /run` に変更し、Cloud Run の URL を audience とする OIDC 認証を設定する。

同じ対象日の完了状態は Cloud Storage に永続化する。書き込み可能な専用バケットを
用意し、その名前を Repository variable `RUN_STATE_BUCKET` に設定する。Cloud Run の
サービスアカウントには、このバケットの `roles/storage.objectUser` が必要となる。
再試行時は直近 100 件の Discord 投稿も確認し、既に投稿済みの動画 URL は送信しない。
