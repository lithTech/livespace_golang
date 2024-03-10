package config

import (
	"log/slog"
	"os"
	"strings"
	"time"
)

func SetLogger(args []string) {
	var level = slog.LevelInfo
	for i, v := range args {
		if v == "log.level" && len(args) > i + 1 {
			var strLevel string = strings.ToLower(args[i + 1])
			switch strLevel {
			case "debug":
				level = slog.LevelDebug
			case "error":
				level = slog.LevelError
			default:
				level = slog.LevelInfo
			}
		}
	}
	logHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     level,
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Key = "date"
				a.Value = slog.Int64Value(time.Now().Unix())
			}
			return a
		},
	})

	logger := slog.New(logHandler)

	slog.SetDefault(logger)
}
