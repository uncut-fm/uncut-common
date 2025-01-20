package entgo

import (
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/XSAM/otelsql"
	"github.com/uncut-fm/uncut-common/pkg/config"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	// Required import for Ent GraphQL Postgres connection.
	_ "github.com/lib/pq"
)

func InitPGDriver(dbConfigs config.DBConfigs, tp trace.TracerProvider) (*entsql.Driver, error) {
	connectionStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", dbConfigs.Host, dbConfigs.Port, dbConfigs.DBName, dbConfigs.User, dbConfigs.Password)

	dbDriver, err := otelsql.Open("pgx", connectionStr, otelsql.WithTracerProvider(tp), otelsql.WithAttributes(semconv.DBSystemPostgreSQL), otelsql.WithSQLCommenter(true))
	if err != nil {
		return nil, err
	}

	err = otelsql.RegisterDBStatsMetrics(dbDriver, otelsql.WithAttributes(
		semconv.DBSystemPostgreSQL,
	))

	maxConnLifetime, err := time.ParseDuration(dbConfigs.ConnectionMaxLifetime)
	if err != nil {
		return nil, err
	}

	dbDriver.SetMaxOpenConns(dbConfigs.MaxOpenConnections)
	dbDriver.SetMaxIdleConns(dbConfigs.MaxIdleConnections)
	dbDriver.SetConnMaxLifetime(maxConnLifetime)

	return entsql.OpenDB(dialect.Postgres, dbDriver), nil
}
