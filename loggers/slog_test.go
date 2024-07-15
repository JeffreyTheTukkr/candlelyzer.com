package loggers

import (
	"log/slog"
	"os"
	"testing"
)

func Test_getLogLevel(t *testing.T) {
	tests := []struct {
		name string
		env  string
		want slog.Level
	}{
		{name: "error", env: "error", want: slog.LevelError},
		{name: "warn", env: "warn", want: slog.LevelWarn},
		{name: "info", env: "info", want: slog.LevelInfo},
		{name: "debug", env: "debug", want: slog.LevelDebug},
		{name: "default", env: "default", want: slog.LevelWarn},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = os.Setenv("LOG_LEVEL", tt.env)

			if got := getLogLevel(); got != tt.want {
				t.Errorf("getLogLevel() returned  %v, wants %v", got, tt.want)
			}
		})
	}
}
