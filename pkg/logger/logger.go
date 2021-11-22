package logger

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"reflect"
	"runtime"
)

func NewLogger() Logger {
	logger := logrus.New()

	return log{logger}
}

type log struct {
	*logrus.Logger
}

type Logger interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})
	CheckError(err error, i interface{}) error
	CheckErrPanic(err error)
	BeforeQuery(context.Context, *pg.QueryEvent) (context.Context, error)
	AfterQuery(context.Context, *pg.QueryEvent) error
}

func (l log) CheckError(err error, i interface{}) error {
	if err != nil {
		l.Warn("Function name: "+getFunctionName(i)+" | Error: ", err)
	}

	return err
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func (l log) CheckErrPanic(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func (l log) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (l log) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
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
