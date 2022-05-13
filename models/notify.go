package models

import "fmt"

// Notification type, used for sending the notification to all subscribers
type Notification struct {
	Token string   `json:"token"`
	Words []string `json:"words"`
}

func (n Notification) String() string {
	return fmt.Sprintf("Notification { Token: %s, URL: %s}", n.Token, n.Words)
}
