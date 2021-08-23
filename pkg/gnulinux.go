package bot

import (
	"strings"

	"github.com/slack-go/slack"
)

type GNULinux struct{}

func (g GNULinux) Response(msg *slack.MessageEvent) string {
	if strings.Contains(strings.ToLower(msg.Text), "gnu linux") {
		return ""
	}

	if !strings.Contains(strings.ToLower(msg.Text), "linux") {
		return ""
	}

	return "Do you mean GNU Linux?"
}
