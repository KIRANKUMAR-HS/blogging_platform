package logger

import (
	syslog "log"
	"os"

	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func Init() {

	switch strings.ToLower(viper.GetString("logger.level")) {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if viper.GetString("logger.file") != "" {
		f, err := os.OpenFile(viper.GetString("logger.file"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			syslog.Fatalf("Error opening log file %s: %s", viper.GetString("logger.file"), err)
		}
		log.Logger = zerolog.New(f).With().Timestamp().Caller().Logger()
		return
	}
	// for json logging:
	// log.Logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()

	// for user friendly logging:
	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false}).With().Timestamp().Caller().Logger()
}
