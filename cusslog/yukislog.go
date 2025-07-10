package cusslog

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

var (
	lvFatal = slog.Level(12)
)

func InitCussLog() {
	if strings.ToLower(os.Getenv("ENV")) != "local" {
		color.NoColor = true
	}

	setSlogDefaultLogger()
}

type (
	YukiHandlerOptions struct {
		SlogOpts slog.HandlerOptions
	}

	YukiHandler struct {
		slog.Handler
		l *log.Logger
	}
)

func (h *YukiHandler) Handle(ctx context.Context, r slog.Record) error {
	lvStr := buildLevelString(r.Level)

	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()

		return true
	})

	b, err := json.MarshalIndent(fields, "", "  ")
	if err != nil {
		return err
	}

	timeStr := "[" + time.Now().Format(time.RFC3339Nano) + "]"
	msg := color.WhiteString(r.Message)

	h.l.Println(timeStr, lvStr, msg, string(b))

	return nil
}

func NewYukiJSONHandler(w io.Writer, opts YukiHandlerOptions) *YukiHandler {
	return &YukiHandler{
		Handler: slog.NewJSONHandler(w, &opts.SlogOpts),
		l:       log.New(w, "", 0),
	}
}

func NewYukiTextHandler(w io.Writer, opts YukiHandlerOptions) *YukiHandler {
	return &YukiHandler{
		Handler: slog.NewTextHandler(w, &opts.SlogOpts),
		l:       log.New(w, "", 0),
	}
}

func setSlogDefaultLogger() {
	opts := YukiHandlerOptions{
		SlogOpts: slog.HandlerOptions{
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
		},
	}

	var handler slog.Handler
	if os.Getenv("LOG_FORMAT") == "json" {
		handler = NewYukiJSONHandler(os.Stdout, opts)
	} else {
		handler = NewYukiTextHandler(os.Stdout, opts)
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)
}

func setLogLevel(levelStr string) slog.Level {
	lv := slog.LevelInfo

	switch levelStr {
	case "debug":
		lv = slog.LevelDebug
	case "info":
		lv = slog.LevelInfo
	case "warn":
		lv = slog.LevelWarn
	case "error":
		lv = slog.LevelError
	case "fatal":
		lv = lvFatal
	}

	return lv
}

func buildLevelString(level slog.Level) string {
	f := "[" + level.String() + "]"
	switch level {
	case slog.LevelDebug:
		return color.CyanString(f)
	case slog.LevelInfo:
		return color.GreenString(f)
	case slog.LevelWarn:
		return color.YellowString(f)
	case slog.LevelError:
		return color.RedString(f)
	case lvFatal:
		return color.HiRedString(f)
	default:
		return color.HiBlackString(f)
	}
}

func Debug(msg string, args ...any) {
	slog.Debug(msg, args...)
}

func Info(msg string, args ...any) {
	slog.Info(msg, args...)
}

func Error(msg string, args ...any) {
	slog.Error(msg, args...)
}

func Warn(msg string, args ...any) {
	slog.Warn(msg, args...)
}
