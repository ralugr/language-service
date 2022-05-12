package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ralugr/language-service/handlers"
	"github.com/ralugr/language-service/repository"
	"github.com/ralugr/language-service/service"
)

func main() {
	bannedWordsRepo := repository.New("banned-words.json")
	subscribers := repository.New("subscribers.json")

	service := service.New(bannedWordsRepo, subscribers)
	handler := handlers.New(service)

	defer service.BannedWords.Close()
	defer service.Subscribers.Close()

	writePID()

	log.Fatal(http.ListenAndServe(":8081", routes(handler)))
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
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(fmt.Sprintf("%d", pid))

	if err2 != nil {
		log.Fatal(err2)
	}
}
