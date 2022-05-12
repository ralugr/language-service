package service

import (
	"bytes"
	"encoding/json"
	"github.com/ralugr/language-service/adapters"
	"github.com/ralugr/language-service/logger"
	"github.com/ralugr/language-service/models"
	"github.com/ralugr/language-service/repository"
	"net/http"
)

type Service struct {
	BannedWords repository.Base
	Subscribers repository.Base
}

func New(bannedWords repository.Base, subscribers repository.Base) Service {
	return Service{
		BannedWords: bannedWords,
		Subscribers: subscribers,
	}
}

func (s Service) Notify() {
	logger.Info.Printf("Sending notifications to all subscribers")
	var subscribers []models.Subscriber

	data, err := s.Subscribers.Read()
	if err != nil {
		logger.Warning.Printf("Could not read subscriber list %v", err)
		return
	}

	if err := json.Unmarshal(data, &subscribers); err != nil {
		logger.Warning.Printf("Could not unmarshall subscriber list %v, error %v", string(data), err)
	}

	var notification models.Notification
	list, err := s.BannedWords.Read()

	if err != nil {
		logger.Warning.Printf("Could not read banned words %v", err)
	}
	notification.Words, err = adapters.ConvertFromByteToStringArray(list)

	for _, s := range subscribers {
		notification.Token = s.Token

		data, err = json.Marshal(notification)
		if err != nil {
			logger.Info.Printf("Could not marshall notification %v", notification)
		}

		_, err := http.Post(s.URL, "application/json", bytes.NewReader(data))

		if err != nil {
			logger.Warning.Printf("POST failed for subscriber %v, error %v", s, err)
		}
	}

	logger.Info.Printf("Notified all subscribers")
}

func (s Service) AddSubscriber(subscriber models.Subscriber) error {
	data, err := s.Subscribers.Read()
	var existingSubscribers []models.Subscriber

	if len(data) != 0 {
		if err := json.Unmarshal(data, &existingSubscribers); err != nil {
			logger.Warning.Printf("Could not unmarshall existing subscriber %s. Please provide token(string) and url(string) keys. %v", data, err)
			return err
		}
	}
	if s.tokenExists(subscriber.Token, existingSubscribers) {
		logger.Warning.Printf("Subscriber with token %s, already exists.", subscriber.Token)
		return nil
	}

	existingSubscribers = append(existingSubscribers, subscriber)
	encoded, err := json.Marshal(existingSubscribers)

	if err != nil {
		logger.Warning.Printf("Could not save new subscriber %v, error %v", subscriber, err)
		return err
	}

	err = s.Subscribers.Write(encoded, true)

	if err != nil {
		logger.Warning.Printf("Could not write subscribers to file %v", err)
		return err
	}

	return nil
}

func (s Service) tokenExists(token string, list []models.Subscriber) bool {
	for _, s := range list {
		if s.Token == token {
			return true
		}
	}

	return false
}
