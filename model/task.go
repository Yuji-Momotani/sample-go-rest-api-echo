package model

import "time"

type Task struct {
	// GORMの機能でマイグレーション時に自動で複数形にしてくれる。
	// →tasksテーブルを作成
	ID        uint      `json:"id" gorm:"primaryKey"` //intのprimaryKeyでAutoIncrementも自動で付加（Gorm）
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uint      `json:"user_id"`
	User      User      `json:"user" gorm:"forenKey:UserID; constraint:OnDelete:CASCADE"`
}

type TaskResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
