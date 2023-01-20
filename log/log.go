package log

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	level = INFO
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

type Event struct {
	LogLevel LogLevel
	Payload  string
}

func (e *Event) Type() string {
	return e.LogLevel.String()
}

func Infoln(format string, v ...any) {
	event := newLog(INFO, format, v...)
	fmt.Println(event)
}

func Warnln(format string, v ...any) {
	event := newLog(WARNING, format, v...)
	fmt.Println(event)

}

func Errorln(format string, v ...any) {
	event := newLog(ERROR, format, v...)
	fmt.Println(event)
}

func Debugln(format string, v ...any) {
	event := newLog(DEBUG, format, v...)
	fmt.Println(event)
}

func Fatalln(format string, v ...any) {
	log.Fatalf(format, v...)
}

func Level() LogLevel {
	return level
}

func SetLevel(newLevel LogLevel) {
	level = newLevel
}

func SetStrLevel(newLevel string) {
	switch newLevel {
	case "debug", "DEBUG":
		SetLevel(DEBUG)
	case "info", "INFO":
		SetLevel(INFO)
	case "warning", "WARNING":
		SetLevel(WARNING)
	case "error", "ERROR":
		SetLevel(ERROR)
	case "silent", "SILENT":
		SetLevel(SILENT)
	}
	return
}

func print(data Event) {
	if data.LogLevel < level {
		return
	}

	switch data.LogLevel {
	case INFO:
		log.Infoln(data.Payload)
	case WARNING:
		log.Warnln(data.Payload)
	case ERROR:
		log.Errorln(data.Payload)
	case DEBUG:
		log.Debugln(data.Payload)
	}
}

func newLog(logLevel LogLevel, format string, v ...any) Event {
	return Event{
		LogLevel: logLevel,
		Payload:  fmt.Sprintf(format, v...),
	}
}
