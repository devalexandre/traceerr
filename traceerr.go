package traceerr

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var defaultRingBufferSize = 20

// exitFunc allows testing of os.Exit behavior.
var exitFunc = os.Exit

func init() {
	if sizeStr := os.Getenv("LOG_PERSIST_TRACE"); sizeStr != "" {
		if size, err := strconv.Atoi(sizeStr); err == nil && size > 0 {
			defaultRingBufferSize = size
		}
	}
}

type TraceLogger struct {
	buffer *ringBuffer
	debug  bool
}

func NewTraceLogger(bufferSize int) *TraceLogger {
	return &TraceLogger{
		buffer: newRingBuffer(bufferSize),
		debug:  os.Getenv("DEBUG") == "1",
	}
}

var Logger = NewTraceLogger(defaultRingBufferSize)

func Print(v ...interface{}) {
	Logger.log(fmt.Sprint(v...))
}

func Printf(format string, v ...interface{}) {
	Logger.log(fmt.Sprintf(format, v...))
}

func Errorln(v ...interface{}) {
	msg := fmt.Sprint(v...)
	Logger.showTraceAndExit(newTraceError(msg))
}

func Errorf(format string, v ...interface{}) error {
	msg := fmt.Sprintf(format, v...)
	return newTraceError(msg)
}

func (tl *TraceLogger) log(msg string) {
	timestampedMsg := fmt.Sprintf("%s - %s", time.Now().Format("2006-01-02 15:04:05"), msg)
	tl.buffer.add(timestampedMsg)

	if tl.debug {
		fmt.Println(timestampedMsg)
	}
}

func (tl *TraceLogger) showTraceAndExit(err error) {
	tl.buffer.add(err.Error())

	fmt.Fprintln(os.Stderr, "---- LOGS ----")
	for _, log := range tl.buffer.dump() {
		fmt.Fprintln(os.Stderr, log)
	}

	fmt.Fprintln(os.Stderr, "---- ERROR ----")
	if te, ok := err.(stackTracer); ok {
		fmt.Fprintln(os.Stderr, te.Error())
		fmt.Fprintln(os.Stderr, te.StackTrace())
	} else {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	exitFunc(1)
}

type traceError struct {
	message string
	stack   string
}

func newTraceError(msg string) error {
	return &traceError{
		message: msg,
		stack:   generateStackTrace(3),
	}
}

func (te *traceError) Error() string {
	return te.message
}

func (te *traceError) StackTrace() string {
	return te.stack
}

type stackTracer interface {
	Error() string
	StackTrace() string
}

func generateStackTrace(skip int) string {
	var stackTrace strings.Builder
	stackTrace.WriteString("Stack trace:\n")
	for i := skip; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fn := runtime.FuncForPC(pc).Name()
		stackTrace.WriteString(fmt.Sprintf("\t%s:%d (%s)\n", file, line, fn))
	}
	return stackTrace.String()
}
