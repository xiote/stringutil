package zerolog

import (
	"fmt"
	// "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type StderrWriter struct {
	//
}

func (w StderrWriter) Write(appName string, goodsCode string, loginId string, ticketingId string, stepName string, message string) {

	// fmt.Fprintln(os.Stderr, "hello world")

	go println(fmt.Sprintf("%s | %s | %s | %s | %s | %s | %s",
		appName,
		goodsCode,
		loginId,
		ticketingId,
		// realClock.Now().Format("0102_15:04:05.000"),
		time.Now().Format("0102_15:04:05.000"),
		stepName,
		message,
	))
}

func TestEventErr(t *testing.T) {
	var logger = New(StderrWriter{}).With().Logger()

	logger.Log().
		Err(fmt.Errorf("ABC"))
	<-time.After(time.Second)
}
