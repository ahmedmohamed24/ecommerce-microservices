package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/ahmedmohamed24/ecommerce-microservices/product/configs"
)

type SlackClient struct {
	WebHookUrl string
	UserName   string
	Channel    string
}

func NewSlackClient(c *configs.Config) *SlackClient {
	return &SlackClient{
		WebHookUrl: c.Slack.WebHookUrl,
		Channel:    c.Slack.Channel,
		UserName:   c.Slack.UserName,
	}
}
func (sc SlackClient) Write(p []byte) (n int, err error) {
	message := struct {
		Username  string `json:"username,omitempty"`
		IconEmoji string `json:"icon_emoji,omitempty"`
		Channel   string `json:"channel,omitempty"`
		Text      string `json:"text,omitempty"`
	}{
		Text:      string(p),
		IconEmoji: ":hammer_and_wrench",
		Channel:   sc.Channel,
		Username:  sc.UserName,
	}
	slackBody, _ := json.Marshal(message)
	if sc.WebHookUrl == "" {
		// in case of testing
		return 0, nil
	}
	req, err := http.NewRequest(http.MethodPost, sc.WebHookUrl, bytes.NewBuffer(slackBody))
	if err != nil {
		return 0, err
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Second * 20}
	resp, err := client.Do(req)
	if err != nil {

		return 0, err
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return 0, err
	}
	msg := buf.String()
	if msg != "ok" {
		return 0, errors.New("non-ok response returned from slack " + msg)
	}
	return len(p), nil

}
