package services

import (
	"database/sql"

	"oly-server/internal/models"
)

type ChatService struct {
	db *sql.DB
}

func NewChatService(db *sql.DB) *ChatService {
	return &ChatService{db: db}
}

type CreateChatInput struct {
	CreatorID string `json:"creatorId"`
}

type ChatResponse struct {
	Success bool          `json:"success"`
	Message string        `json:"message,omitempty"`
	Chat    *models.Chat  `json:"chat,omitempty"`
	Chats   []models.Chat `json:"chats,omitempty"`
}

func (s *ChatService) Create(input CreateChatInput) (*ChatResponse, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	id, err := generateID()
	if err != nil {
		return nil, err
	}

	var newChat models.Chat
	err = tx.QueryRow(
		`INSERT INTO chat (id, creator_id) VALUES ($1, $2) RETURNING id, creator_id, created_at`,
		id, input.CreatorID,
	).Scan(&newChat.ID, &newChat.CreatorID, &newChat.CreatedAt)

	if err != nil {
		return nil, err
	}

	participantID, err := generateID()
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(
		`INSERT INTO chat_participant (id, chat_id, user_id) VALUES ($1, $2, $3)`,
		participantID, newChat.ID, input.CreatorID,
	)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &ChatResponse{
		Success: true,
		Message: "Чат успешно создан",
		Chat:    &newChat,
	}, nil
}

func (s *ChatService) GetUserChats(userID string) (*ChatResponse, error) {
	rows, err := s.db.Query(
		`SELECT c.id, c.creator_id, c.created_at 
		 FROM chat_participant cp 
		 INNER JOIN chat c ON cp.chat_id = c.id 
		 WHERE cp.user_id = $1`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chats []models.Chat
	for rows.Next() {
		var chat models.Chat
		if err := rows.Scan(&chat.ID, &chat.CreatorID, &chat.CreatedAt); err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &ChatResponse{
		Success: true,
		Chats:   chats,
	}, nil
}

func (s *ChatService) AddParticipant(chatID, userID string) (*ChatResponse, error) {
	var existingID string
	err := s.db.QueryRow(
		`SELECT id FROM chat_participant WHERE chat_id = $1 AND user_id = $2 LIMIT 1`,
		chatID, userID,
	).Scan(&existingID)

	if err == nil {
		return &ChatResponse{
			Success: false,
			Message: "Пользователь уже является участником этого чата",
		}, nil
	}
	if err != sql.ErrNoRows {
		return nil, err
	}

	id, err := generateID()
	if err != nil {
		return nil, err
	}

	_, err = s.db.Exec(
		`INSERT INTO chat_participant (id, chat_id, user_id) VALUES ($1, $2, $3)`,
		id, chatID, userID,
	)
	if err != nil {
		return nil, err
	}

	return &ChatResponse{
		Success: true,
		Message: "Пользователь успешно добавлен в чат",
	}, nil
}
