package traceerr

import (
	"os"
	"strings"
	"testing"
)

func TestTraceLoggerPrint(t *testing.T) {
	Logger.buffer = newRingBuffer(3) // Reset the buffer
	Print("test message")

	logs := Logger.buffer.dump()
	if len(logs) != 1 || !strings.Contains(logs[0], "test message") {
		t.Errorf("expected log 'test message', got %v", logs)
	}
}

func TestTraceLoggerFatal(t *testing.T) {
	Logger.buffer = newRingBuffer(3) // Reset the buffer
	Print("test message")

	// Mock exitFunc to prevent termination
	var exited bool
	exitFunc = func(code int) {
		exited = true
	}

	defer func() { exitFunc = os.Exit }() // Restore exitFunc after the test

	Errorln("fatal error")

	if !exited {
		t.Errorf("Errorln did not call exitFunc as expected")
	}
}
