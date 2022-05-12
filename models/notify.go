package models

import "fmt"

type Notification struct {
	Token string   `json:"token"`
	Words []string `json:"words"`
}

func (n Notification) String() string {
	return fmt.Sprintf("Notification { Token: %s, URL: %s}", n.Token, n.Words)
}
