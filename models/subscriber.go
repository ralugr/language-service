package models

import "fmt"

type Subscriber struct {
	Token string `json:"token"`
	URL   string `json:"url"`
}

func (s Subscriber) String() string {
	return fmt.Sprintf("Subscriber { Token: %s, URL: %s}", s.Token, s.URL)
}
