package configs

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	t.Setenv("ENV", "prod")
	c, err := New()
	if err != nil {
		t.Errorf("error in creating config object %v", err)
	}
	if c.Database.Host == "" {
		t.Error("host is empty")
	}
	if c.Slack.WebHookUrl == "" {
		t.Error("webhook url is empty")
	}
}
