package log

import (
	"github.com/rs/zerolog"
	"os"
)

type Log struct {
	infoLogger  *zerolog.Logger
	errorLogger *zerolog.Logger
}

func (log *Log) Info(str string) {
	log.infoLogger.Info().Msg(str)
}

func (log *Log) Error(str string) {
	log.infoLogger.Error().Msg(str)
}

func InitLogger() (*Log, *os.File, *os.File) {
	//UnitFormatter()                  <<
	//todo осознать что это за функция (почему сплитуем по Backend и что за %routers:%d)
	//todo было бы славно узнать как прописывать путь

	loggerInfoFile, err := os.OpenFile("log/info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		panic("Error opening info log file")
	}

	loggerErrorFile, err := os.OpenFile("log/error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		panic("Error opening error log file")
	}

	infoLogger := zerolog.New(loggerInfoFile).With().Timestamp().Caller().Logger()
	errorLogger := zerolog.New(loggerErrorFile).With().Timestamp().Caller().Logger()

	log := &Log{
		infoLogger:  &infoLogger,
		errorLogger: &errorLogger,
	}

	return log, loggerInfoFile, loggerErrorFile
}
