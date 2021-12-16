package db

import (
	"github.com/uncut-fm/uncut-common/pkg/logger"
	"io"
)

func CloseDB(log logger.Logger, db io.Closer) {
	if err := log.CheckError(db.Close(), CloseDB); err != nil {
		log.Error("err closing db connection", err)
	} else {
		log.Info("db connection gracefully closed")
	}
}
