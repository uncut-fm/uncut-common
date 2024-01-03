package logger

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"reflect"
	"runtime"
	"strings"
)

func New() Logger {
	baseLogger := logrus.New()

	logger := log{baseLogger}

	logger.Formatter = &logrus.JSONFormatter{}

	return logger
}

type log struct {
	*logrus.Logger
}

type Logger interface {
	Infof(format string, args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})
	CheckError(err error, i interface{}) error
	CheckInfoError(err error, i interface{}) error
	BeforeQuery(context.Context, *pg.QueryEvent) (context.Context, error)
	AfterQuery(context.Context, *pg.QueryEvent) error
}

func (l log) CheckError(err error, i interface{}) error {
	if err != nil && !strings.Contains(err.Error(), "not found") {
		l.Warn("Function name: "+getFunctionName(i)+" | Error: ", err)
	}

	return err
}

func (l log) CheckInfoError(err error, i interface{}) error {
	if err != nil {
		l.Info("Function name: "+getFunctionName(i)+" | Error: ", err)
	}

	return err
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
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
