package bot

import (
	"strings"
	"time"

	"github.com/slack-go/slack"
)

type CurrentTime struct{}

func (s CurrentTime) Response(msg *slack.MessageEvent) string {
	if !strings.Contains(strings.ToLower(msg.Text), "time") {
		return ""
	}

	t := time.Now()

	return t.String()
}
