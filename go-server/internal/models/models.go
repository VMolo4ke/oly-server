package models

import "time"

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type Chat struct {
	ID        string    `json:"id"`
	CreatorID string    `json:"creatorId"`
	CreatedAt time.Time `json:"createdAt"`
}

type ChatParticipant struct {
	ID       string    `json:"id"`
	ChatID   string    `json:"chatId"`
	UserID   string    `json:"userId"`
	JoinedAt time.Time `json:"joinedAt"`
}

type Message struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	ChatID    string    `json:"chatId"`
	SenderID  string    `json:"senderId"`
	CreatedAt time.Time `json:"createdAt"`
}

type VoiceCall struct {
	ID        string    `json:"id"`
	ChatID    string    `json:"chatId"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

type VoiceParticipant struct {
	ID       string    `json:"id"`
	CallID   string    `json:"callId"`
	UserID   string    `json:"userId"`
	JoinedAt time.Time `json:"joinedAt"`
}
