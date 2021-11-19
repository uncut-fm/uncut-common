package entgo

import (
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/uncut-fm/uncut-common/pkg/config"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	// Required import for Ent GraphQL Postgres connection.
	_ "github.com/lib/pq"
)

func InitPGDriver(dbConfigs config.DBConfigs) (*sql.Driver, error) {
	connectionStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", dbConfigs.Host, dbConfigs.Port, dbConfigs.DBName, dbConfigs.User, dbConfigs.Password)

	dbDriver, err := sql.Open("pgx", connectionStr)
	if err != nil {
		return nil, err
	}

	maxConnLifetime, err := time.ParseDuration(dbConfigs.ConnectionMaxLifetime)
	if err != nil {
		return nil, err
	}

	dbDriver.DB().SetMaxOpenConns(dbConfigs.MaxOpenConnections)
	dbDriver.DB().SetMaxIdleConns(dbConfigs.MaxIdleConnections)
	dbDriver.DB().SetConnMaxLifetime(maxConnLifetime)

	return dbDriver, nil
}
