package bot

import (
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/slack-go/slack"
)

type Status struct {
	latency time.Duration
}

func (s Status) Response(msg *slack.MessageEvent) string {
	if !strings.Contains(strings.ToLower(msg.Text), "how do you feel?") {
		return ""
	}

	text := "my memory consumption is at %.2f%% (%d MB), my response time is %.2fs"

	abs, perc, err := memoryUsage()
	if err != nil {
		log.Printf("calculate memory usage: %v", err)
		return ""
	}

	return fmt.Sprintf(text, perc, byteToMegabyte(abs), s.latency.Seconds())
}

func memoryUsage() (uint64, float32, error) {
	max, err := memoryLimit()
	if err != nil {
		return 0, 0.0, fmt.Errorf("lookup memory limit: %v", err)
	}

	var now runtime.MemStats
	runtime.ReadMemStats(&now)

	return now.Alloc, float32(now.Alloc) * 100.0 / float32(max), nil
}

func byteToMegabyte(b uint64) uint64 {
	return b / 1024 / 1024
}

func memoryLimit() (int, error) {
	out, err := ioutil.ReadFile("/sys/fs/cgroup/memory/memory.limit_in_bytes")
	if err != nil {
		return 0, fmt.Errorf("reading memory limit from /sys/fs/cgroup/memory/memory.limit_in_bytes: %v", err)
	}

	i, err := strconv.Atoi(strings.TrimSpace(string(out)))
	if err != nil {
		return 0, fmt.Errorf("convert cgroup memory limit to an int : %v", err)
	}

	return i, nil
}
