package loggerconfig

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func ConfigureBaseLogger() {
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	log.Logger = log.With().Str("application-name", "TarefasCrud").Logger()
}
