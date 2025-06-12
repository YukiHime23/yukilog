package cusslog

import (
	"log/slog"
	"os"

	"github.com/fatih/color"
)

var (
	lvFatal = slog.Level(12)
)

func InitCussLog() {
	setSlogDefaultLogger()
}

func setSlogDefaultLogger() {
	var logger *slog.Logger

	opts := slog.HandlerOptions{
		Level: setLogLevel(os.Getenv("LOG_LEVEL")),
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				lv := a.Value.Any().(slog.Level)
				if lv == lvFatal {
					a.Value = slog.StringValue("FATAL")
				}
			}
			return a
		},
	}

	if os.Getenv("LOG_FORMAT") == "json" {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &opts))
	} else {
		logger = slog.New(slog.NewTextHandler(os.Stdout, &opts))
	}

	// levelStr := os.Getenv("LOG_LEVEL")
	// lv := setLogLevel(levelStr)

	slog.SetDefault(logger)
}

func setLogLevel(levelStr string) slog.Level {
	lv := slog.LevelInfo

	switch levelStr {
	case "debug":
		color.CyanString("Debug")
		lv = slog.LevelDebug
	case "info":
		color.GreenString("Info")
		lv = slog.LevelInfo
	case "warn":
		color.YellowString("Warn")
		lv = slog.LevelWarn
	case "error":
		color.RedString("Error")
		lv = slog.LevelError
	case "fatal":
		color.HiRedString("Fatal")
		lv = slog.Level(12)
	}

	return slog.Level(lv)
}

func Info(msg string, args ...any) {
	slog.Info(msg, args...)
}

func Debug(msg string, args ...any) {
	slog.Debug(msg, args...)
}

func Error(msg string, args ...any) {
	slog.Error(msg, args...)
}

func Warn(msg string, args ...any) {
	slog.Warn(msg, args...)
}
