package applogger

import (
	"io"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

func New(w io.Writer) log.Logger {
	sw := log.NewSyncWriter(w)
	logger := log.NewJSONLogger(sw)
	logger = log.With(logger,
		"timestamp", log.DefaultTimestampUTC,
	)
	logger = level.NewFilter(logger, level.AllowInfo())

	return logger
}
