package handlers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ralugr/language-service/repository"
	"github.com/ralugr/language-service/service"
	"net/http"
)

func getRoutes() http.Handler {

	bannedWordsRepo := repository.New("banned-words-test.json")
	subscribers := repository.New("subscribers-test.json")

	service := service.New(bannedWordsRepo, subscribers)
	handler := New(service)

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	mux.Get("/", handler.Home)
	mux.Get("/list", handler.List)
	mux.Post("/list", handler.UpdateList)
	mux.Post("/subscribe", handler.Subscribe)

	return mux
}
