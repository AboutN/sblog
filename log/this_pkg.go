package log

import (
	"fmt"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rs/zerolog"
)

var log zerolog.Logger

func Init() {
	lf, err := rotatelogs.New(
		"log/access_%Y%m%d.log",
		rotatelogs.WithLinkName("log/access.log"),
		rotatelogs.WithRotationCount(31),
	)
	if err != nil {
		fmt.Println(err)
		panic("日志工具初始化失败")
	}
	log = zerolog.New(lf).
		Level(zerolog.InfoLevel).
		With().Logger()
}

func Info() *zerolog.Event {
	return log.Info()
}

func Debug() *zerolog.Event {
	return log.Debug()
}

func Error() *zerolog.Event {
	return log.Error()
}

func Warn() *zerolog.Event {
	return log.Warn()
}
