#!/bin/bash

# マイグレーション
go run migrate/migrate.go

# ホットリロードでの実行
air

# デバッグでの実行
# dlv debug ./main.go --headless --listen=:2345 --log --api-version=2