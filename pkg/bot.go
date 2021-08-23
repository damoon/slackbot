package bot

import (
	"log"

	"github.com/slack-go/slack"
)

type bot struct {
	api        *slack.Client
	responders []Reponder
	status     *Status
}

type Reponder interface {
	Response(msg *slack.MessageEvent) string
}

func NewBot(token string, responders []Reponder) *bot {
	status := &Status{}
	responders = append(responders, status)

	return &bot{
		api:        slack.New(token),
		responders: responders,
		status:     status,
	}
}

func (b *bot) Run() {
	rtm := b.api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		b.processEvent(msg, rtm)
	}
}

func (b *bot) processEvent(msg slack.RTMEvent, rtm *slack.RTM) {

	switch ev := msg.Data.(type) {

	case *slack.ConnectingEvent:
		log.Println("connecting")

	case *slack.ConnectedEvent:
		log.Println("connected")

	case *slack.HelloEvent:
		log.Println("hello")

	case *slack.UserTypingEvent:
		// ignore this event

	case *slack.PresenceChangeEvent:
		// ignore this event

	case *slack.MessageEvent:
		answer := b.answer(ev)
		if answer != "" {
			rtm.SendMessage(rtm.NewOutgoingMessage(answer, ev.Channel))
		}

	case *slack.DesktopNotificationEvent:
		// ignore this event

	case *slack.LatencyReport:
		b.status.latency = ev.Value
		log.Printf("latency: %v", ev.Value)

	case *slack.RTMError:
		log.Printf("real time messaging error: %v", ev.Error())

	case *slack.InvalidAuthEvent:
		log.Printf("Invalid credentials")
		return

	case *slack.AckMessage:
		// ignore this event

	default:
		log.Printf("unexpected event: %v", msg)
	}

}

func (b *bot) answer(msg *slack.MessageEvent) string {
	for _, responder := range b.responders {
		response := responder.Response(msg)
		if response != "" {
			log.Printf("answering: %s", response)
			return response
		}
	}

	return ""
}
