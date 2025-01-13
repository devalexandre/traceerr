package traceerr

import (
	"os"
	"strings"
	"testing"
)

func TestErrorln(t *testing.T) {
	Logger.buffer = newRingBuffer(3) // Reset the buffer
	Print("log1")
	Print("log2")

	// Mock exitFunc to prevent program termination
	var exited bool
	exitFunc = func(code int) {
		exited = true
	}

	defer func() { exitFunc = os.Exit }() // Restore exitFunc after test

	Errorln("test error")

	if !exited {
		t.Error("Errorln() did not call exitFunc")
	}
}

func TestErrorf(t *testing.T) {
	Logger.buffer = newRingBuffer(3) // Reset the buffer
	Print("logA")
	Print("logB")

	err := Errorf("formatted %d", 123)
	if err == nil {
		t.Fatalf("Errorf() returned nil")
	}

	if !strings.Contains(err.Error(), "formatted 123") {
		t.Errorf("expected formatted message, got: %s", err.Error())
	}
}
