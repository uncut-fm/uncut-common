package entgo

import (
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/uncut-fm/uncut-common/pkg/config"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	// Required import for Ent GraphQL Postgres connection.
	_ "github.com/lib/pq"
)

func InitPGDriver(dbConfigs config.DBConfigs) (*entsql.Driver, error) {
	connectionStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", dbConfigs.Host, dbConfigs.Port, dbConfigs.DBName, dbConfigs.User, dbConfigs.Password)

	dbDriver, err := sql.Open("pgx", connectionStr)
	if err != nil {
		return nil, err
	}

	maxConnLifetime, err := time.ParseDuration(dbConfigs.ConnectionMaxLifetime)
	if err != nil {
		return nil, err
	}

	dbDriver.SetMaxOpenConns(dbConfigs.MaxOpenConnections)
	dbDriver.SetMaxIdleConns(dbConfigs.MaxIdleConnections)
	dbDriver.SetConnMaxLifetime(maxConnLifetime)

	return entsql.OpenDB(dialect.Postgres, dbDriver), nil
}
