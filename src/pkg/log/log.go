package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var Logger = logrus.New()

// InitializeLogger sets up the logger with preferred formatting and options
func InitializeLogger(level logrus.Level) {
	Logger.SetOutput(os.Stdout)
	Logger.SetLevel(level)

	Logger.SetFormatter(&coloredFormatter{
		DisableColors: false,
	})
}

// Info logs a message at level Info on the Logger
func Info(message string) {
	Logger.Info(message)
}

// Warn logs a message at level Warn on the Logger
func Warn(message string) {
	Logger.Warn(message)
}

// Fatal logs a message at level Fatal on the Logger
func Fatal(message error) {
	Logger.Fatal(message)
}

// Debug logs a message at level Debug on the Logger
func Debug(message string) {
	Logger.Debug(message)
}

func ExecutionTime(fn func() error) error {
	start := time.Now()
	err := fn()
	elapsed := time.Since(start)

	level := logrus.InfoLevel
	messagePrefix := "Done"
	if err != nil {
		level = logrus.ErrorLevel
		messagePrefix = "Errored"
	}

	if elapsed < time.Second {
		Logger.Logf(level, "%s in %d ms", messagePrefix, elapsed.Milliseconds())
	} else {
		Logger.Logf(level, "%s in %.2f s", messagePrefix, elapsed.Seconds())
	}

	return err
}

type coloredFormatter struct {
	DisableColors bool
}

func (f *coloredFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var color string

	// Determine the color based on the log level
	switch entry.Level {
	case logrus.DebugLevel:
		color = gray
	case logrus.InfoLevel:
		color = green
	case logrus.WarnLevel:
		color = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		color = red
	default:
		color = ""
	}

	// Create the formatted log message
	message := entry.Message
	if !f.DisableColors && color != "" {
		message = fmt.Sprintf("%s%s%s", color, message, reset)
	}

	formatted := fmt.Sprintf("%s\n", message)
	return []byte(formatted), nil
}

const (
	gray   = "\033[90m"
	green  = "\033[32m"
	yellow = "\033[33m"
	red    = "\033[31m"
	reset  = "\033[0m"
)
