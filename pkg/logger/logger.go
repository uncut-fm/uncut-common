package logger

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

func (l Logger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (l Logger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	str, err := q.FormattedQuery()
	if err != nil {
		return err
	}
	if q.Err != nil {
		l.Warn(string(str))
		return nil
	}

	l.Info(string(str))
	return nil
}

func NewLogger() Logger {
	logger := logrus.New()
	return Logger{logger}
}
