package main

import (
	"log"
	"net/http"
	"strings"

	"oly-server/internal/config"
	"oly-server/internal/database"
	"oly-server/internal/handlers"
	"oly-server/internal/middleware"
	"oly-server/internal/services"
)

func main() {
	cfg := config.Load()

	db, err := database.NewDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	authService := services.NewAuthService(db.DB)
	chatService := services.NewChatService(db.DB)

	authHandler := handlers.NewAuthHandler(authService, cfg.JWTSecret)
	chatHandler := handlers.NewChatHandler(chatService)
	jwtMiddleware := middleware.NewJWTMiddleware(cfg.JWTSecret)

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" && r.Method == http.MethodGet {
			w.Write([]byte("Hello"))
			return
		}
		http.NotFound(w, r)
	})

	mux.HandleFunc("/auth/sign-up", authHandler.SignUp)
	mux.HandleFunc("/auth/sign-in", authHandler.SignIn)

	mux.HandleFunc("/chat/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/chat")

		if path == "" || path == "/" {
			if r.Method == http.MethodGet {
				chatHandler.GetUserChats(w, r)
				return
			}
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if strings.HasPrefix(path, "/create") {
			chatHandler.CreateChat(w, r)
			return
		}

		if strings.Contains(path, "/participants") {
			chatHandler.AddParticipant(w, r)
			return
		}

		http.NotFound(w, r)
	})

	protectedMux := http.NewServeMux()
	protectedMux.Handle("/chat/", jwtMiddleware.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/chat")
		if path == "" || path == "/" {
			if r.Method == http.MethodGet {
				chatHandler.GetUserChats(w, r)
				return
			}
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if strings.HasPrefix(path, "/create") {
			chatHandler.CreateChat(w, r)
			return
		}
		if strings.Contains(path, "/participants") {
			chatHandler.AddParticipant(w, r)
			return
		}
		http.NotFound(w, r)
	})))

	log.Println("Server starting on :3000")
	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
