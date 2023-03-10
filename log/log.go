package log

import (
	"fmt"
	"os"
	"time"

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

func getTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func Println(format string, v ...any) {
	if v == nil {
		fmt.Println("[" + getTimeStr() + "] [Core] " + format)
		return
	}
	fmt.Printf("["+getTimeStr()+"] [Core] "+format+"\n", v)
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

func Throwln(a string, b string, err error) {
	fmt.Println("\n---------------------------")
	Errorln(a+" Error, %s", err)
	Println(b)
	fmt.Println("---------------------------\n")
	os.Exit(1)
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

func newLog(logLevel LogLevel, format string, v ...any) Event {
	return Event{
		LogLevel: logLevel,
		Payload:  fmt.Sprintf(format, v...),
	}
}
