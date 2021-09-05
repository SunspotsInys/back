package logs

import (
	"io"
	"os"

	"github.com/SunspotsInys/thedoor/configs"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger zerolog.Logger

func init() {
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05"
	ljWrite := &lumberjack.Logger{
		Filename:   configs.Conf.LogFilename,
		MaxSize:    configs.Conf.LogMaxAge,
		MaxBackups: configs.Conf.LogMaxBackups,
		MaxAge:     configs.Conf.LogMaxAge,
		Compress:   false,
	}
	var writer io.Writer
	if configs.Conf.LogConsole {
		writer = zerolog.MultiLevelWriter(
			ljWrite, zerolog.ConsoleWriter{Out: os.Stdout},
		)
	} else {
		writer = ljWrite
	}

	logger = zerolog.New(writer).With().
		Timestamp().Logger().Hook(new(h1))
	switch configs.Conf.LogLevel {
	case "debug":
		SetLogLevel(zerolog.DebugLevel)
	case "release":
		SetLogLevel(zerolog.InfoLevel)
	}
}

func SetLogLevel(l zerolog.Level) {
	logger.Level(l)
}

func Error(msg string) {
	logger.Error().Msg(msg)
}

func Errorf(format string, v ...interface{}) {
	logger.Error().Msgf(format, v...)
}

func Info(msg string) {
	logger.Info().Msg(msg)
}

func Infof(format string, v ...interface{}) {
	logger.Info().Msgf(format, v...)
}

func Debug(msg string) {
	logger.Debug().Msg(msg)
}

func Debugf(format string, v ...interface{}) {
	logger.Debug().Msgf(format, v...)
}
