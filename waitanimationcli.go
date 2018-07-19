package cliwait

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

func DoFuncWithWaitAnimation(text string, f func()) {
	done := make(chan struct{})
	go func() {
		f()
		done <- struct{}{}
	}()

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	green := color.New(color.FgHiGreen).FprintFunc()
	clockStates := [...]string{"-", "\\", "|", "/"}
	currClockState := 0

	stdout := bufio.NewWriter(os.Stdout)
	for {
		select {
		case <-ticker.C:
			green(color.Output, fmt.Sprintf("\r%s %s", text, clockStates[currClockState]))
			stdout.Flush()
			currClockState = (currClockState + 1) % len(clockStates)
		case <-done:
			fmt.Fprintf(color.Output, "\r%s\r", strings.Repeat(" ", len(text)+4))
			stdout.Flush()
			return
		}
	}
}
