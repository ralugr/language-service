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

// Handler stores all the available handlers
type Handler struct {
	s service.Service
}

// New constructor
func New(s service.Service) Handler {
	return Handler{s}
}

// Home route handlers
func (h Handler) Home(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Welcome to the language service!!")); err != nil {
		respond.Error(w, 500, "Encountered internal error")
	}
}

// List returns the list of banned words
func (h Handler) List(w http.ResponseWriter, r *http.Request) {
	words, err := h.s.BannedWords.Read()

	if err != nil {
		logger.Warning.Printf("Could not get banned words %v", err)
		respond.Error(w, 500, "Encountered internal error")
		return
	}

	w.WriteHeader(200)
	if _, err := w.Write(words); err != nil {
		respond.Error(w, 500, "Encountered internal error")
	}
}

// Subscribe adds a new subscribers for banned list updates
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

// UpdateList used to simulate a change to the banned list
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
