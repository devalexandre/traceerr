package traceerr

import (
	"testing"
)

func TestRingBufferAddAndDump(t *testing.T) {
	rb := newRingBuffer(3)
	rb.add("log1")
	rb.add("log2")
	rb.add("log3")
	logs := rb.dump()
	if len(logs) != 3 {
		t.Errorf("expected 3 logs, got %d", len(logs))
	}
	if logs[0] != "log1" || logs[1] != "log2" || logs[2] != "log3" {
		t.Errorf("logs order mismatch: %v", logs)
	}
}

func TestRingBufferOverflow(t *testing.T) {
	rb := newRingBuffer(2)
	rb.add("logA")
	rb.add("logB")
	rb.add("logC") // should overwrite the oldest
	logs := rb.dump()
	if len(logs) != 2 {
		t.Errorf("expected 2 logs, got %d", len(logs))
	}
	// "logA" should have been overwritten by "logC"
	if logs[0] != "logB" || logs[1] != "logC" {
		t.Errorf("overflow mismatch, logs: %v", logs)
	}
}
