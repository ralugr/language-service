package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ralugr/language-service/handlers"
	"github.com/ralugr/language-service/logger"
	"github.com/ralugr/language-service/repository"
	"github.com/ralugr/language-service/service"
	"net/http"
	"os"
)

func main() {
	bannedWordsRepo := repository.New("banned-words.json")
	subscribers := repository.New("subscribers.json")

	srv := service.New(bannedWordsRepo, subscribers)
	handler := handlers.New(srv)

	defer srv.BannedWords.Close()
	defer srv.Subscribers.Close()

	writePID()

	logger.Warning.Fatal(http.ListenAndServe(":8081", routes(handler)))
}

func routes(h handlers.Handler) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	mux.Get("/", h.Home)
	mux.Get("/list", h.List)
	mux.Post("/list", h.UpdateList)
	mux.Post("/subscribe", h.Subscribe)

	return mux
}

func writePID() {
	pid := os.Getpid()

	f, err := os.Create("language_service.pid")

	if err != nil {
		logger.Warning.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(fmt.Sprintf("%d", pid))

	if err2 != nil {
		logger.Warning.Fatal(err2)
	}
}
