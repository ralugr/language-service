package handlers

import (
	"encoding/json"
	"github.com/ralugr/language-service/logger"
	"github.com/ralugr/language-service/models"
	"github.com/ralugr/language-service/respond"
	"github.com/ralugr/language-service/service"
	"io"
	"net/http"
)

type Handler struct {
	s service.Service
}

func New(s service.Service) Handler {
	return Handler{s}
}

func (h Handler) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the language service!!</h1>"))
}

func (h Handler) List(w http.ResponseWriter, r *http.Request) {
	words, err := h.s.BannedWords.Read()

	if err != nil {
		logger.Warning.Printf("Could not get banned words %v", err)
		respond.Error(w, 500, "Encountered internal error")
		return
	}

	w.WriteHeader(200)
	w.Write(words)
}

func (h Handler) Subscribe(w http.ResponseWriter, r *http.Request) {
	subscriber, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Warning.Printf("Failed to read subscribers from file %v", err)
		respond.Error(w, 500, "Encountered internal error")
		return
	}

	newSubscriber := models.Subscriber{}
	if err := json.Unmarshal(subscriber, &newSubscriber); err != nil {
		logger.Warning.Printf("Could not unmarshall subscriber %s. Please provide token(string) and url(string) keys.", subscriber)
		respond.Error(w, 400, "Validation failed. Please provide object with token(string) url(string)")
		return
	}

	if err := h.s.AddSubscriber(newSubscriber); err != nil {
		logger.Warning.Printf("Failed to add subscribers to file %v", err)
		respond.Error(w, 500, "Encountered internal error")
		return
	}

	respond.Success(w, "Subscribers were updated")
}

func (h Handler) UpdateList(w http.ResponseWriter, r *http.Request) {
	list, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Warning.Printf("Failed to read words from file %v", err)
		respond.Error(w, 500, "Encountered internal error")
		return
	}

	var wordList []string
	if err := json.Unmarshal(list, &wordList); err != nil {
		logger.Warning.Printf("Could not unmarshall word list %s. Please provide a json array of strings.", list)
		respond.Error(w, 400, "Validation failed. Provide an array of strings")
		return
	}

	err = h.s.BannedWords.Write(list, true)

	if err != nil {
		logger.Warning.Printf("Failed to write to words file %v", err)
		respond.Error(w, 500, "Encountered internal error")
		return
	}

	h.s.Notify()
	respond.Success(w, "List was updated")
}
