package setup

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var loggers = make(map[string]zerolog.Logger)

func Zerolog(names ...string) {
	_, f, _, _ := runtime.Caller(0)
	split := strings.Split(f, "conf")
	for _, name := range names {
		logName := name + ".log"
		openFile, err := os.OpenFile(filepath.Join(split[0], "logs", logName), os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend|os.ModePerm)
		if err != nil {
			log.Err(err).Send()
		}
		loggers[name] = zerolog.New(openFile).With().Timestamp().Logger()
	}
}

func Logger(name string) *zerolog.Logger {
	instance := loggers[name]
	return &instance
}
