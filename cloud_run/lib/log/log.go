package log

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type severity string

// ref: https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry?hl=ja#LogSeverity
const (
	severityError severity = "ERROR"
	severityInfo  severity = "INFO"
)

func Infof(ctx context.Context, format string, v ...interface{}) {
	infof(ctx, fmt.Sprintf(format, v...))
}

func Errorf(ctx context.Context, format string, v ...interface{}) {
	errorf(ctx, fmt.Sprintf(format, v...))
}

func infof(ctx context.Context, msg string) {
	logPrintf(ctx, severityInfo, msg)
}

func errorf(ctx context.Context, msg string) {
	logPrintf(ctx, severityError, msg)
}

func logPrintf(ctx context.Context, s severity, msg string) {
	logger := log.New(getWriter(), "", 0)

	// The info for struct logging are listed below.

	// ref: https://cloud.google.com/logging/docs/agent/configuration?hl=ja#special-fields
	entry := map[string]interface{}{
		"message":  msg,
		"severity": s,
	}
	payload, _ := json.Marshal(entry)
	logger.Println(string(payload))
}

var getWriter = func() io.Writer {
	return os.Stdout
}
