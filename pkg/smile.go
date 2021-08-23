package bot

import (
	"strings"

	"github.com/slack-go/slack"
)

type Smile struct{}

func (s Smile) Response(msg *slack.MessageEvent) string {
	if !strings.Contains(strings.ToLower(msg.Text), ":smile:") {
		return ""
	}

	return ":smile:"
}
