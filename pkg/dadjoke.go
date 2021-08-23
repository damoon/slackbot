package bot

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/slack-go/slack"
)

type Dadjoke struct {
}

func (s Dadjoke) Response(msg *slack.MessageEvent) string {
	if !strings.Contains(strings.ToLower(msg.Text), "joke") {
		return ""
	}

	joke, err := fetchJoke()
	if err != nil {
		log.Printf("fetching joke: %v", err)
		return ""
	}

	return joke
}

func fetchJoke() (string, error) {
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: time.Second,
		}).Dial,
		TLSHandshakeTimeout: time.Second,
	}

	var netClient = &http.Client{
		Timeout:   time.Second,
		Transport: netTransport,
	}

	req, err := http.NewRequest(http.MethodGet, "https://icanhazdadjoke.com/", nil)
	if err != nil {
		return "", fmt.Errorf("prepare http request: %v", err)
	}

	req.Header.Add("Accept", "text/plain")

	resp, err := netClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("fetching icanhazdadjoke.com: %v", err)
	}
	defer resp.Body.Close()

	joke, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reading response from icanhazdadjoke.com: %v", err)
	}

	return string(joke), nil
}
