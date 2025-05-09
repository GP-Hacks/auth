package logger

import (
	"os"
	"time"

	"github.com/GP-Hacks/auth/internal/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SetupLogger() {
	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	multi := zerolog.MultiLevelWriter(consoleWriter)
	if config.Cfg.Logging.IsProduction {
		httpWriter := NewHTTPWriter(config.Cfg.Logging.VectorURL)
		multi = zerolog.MultiLevelWriter(httpWriter, consoleWriter)
	}

	log.Logger = zerolog.New(multi).
		With().
		Timestamp().
		Caller().
		Str("service", "auth").
		Logger()
}
