package model

import "time"

type User struct {
	// GORMの機能でマイグレーション時に自動で複数形にしてくれる。
	// →usersテーブルを作成
	ID        uint      `json:"id" gorm:"primaryKey"` //intのprimaryKeyを指定することでAuotIncrementも設定される。（GORM）
	Email     string    `json:"email" gorm:"unique; not null;"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}
