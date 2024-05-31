# sample-go-rest-api-echo

Go言語で実装するrest apiのサンプルコード。（フレームワークはEchoを使用）
アーキテクチャはCleanArchitectureを採用。

## 機能概要
簡単な認証とタスク管理をするAPI

| パス | HttpMethod | 処理内容 |
| ---- | ---- | ---- |
| /signup | POST | ユーザー登録 |
| /login | POST | ログイン。成功した場合、JWTトークンの取得 |
| /logout | POST | ログアウト。JWTトークンの破棄 |
| /csrf | GET | csrfトークンの取得 |
| /tasks | GET | ユーザーに紐づくタスクの一覧情報を取得 |
| /tasks/:taskId | GET | タスクIDに紐づくタスクを一件取得 |
| /tasks | POST | タスクの登録 |
| /tasks/:taskId | PUT | タスクの更新 |
| /tasks/:taskId | DELETE | タスクの削除 |

## 主な使用ライブラリ

- Echo
- Gorm

# リリースバージョン毎の対応表

## v1.0

機能概要を満たすアプリケーションの実装

## v2.0

テストコードの追加

