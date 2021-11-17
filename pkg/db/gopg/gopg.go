package gopg

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/uncut-fm/uncut-common/pkg/config"
	"github.com/uncut-fm/uncut-common/pkg/logger"
)

func InitDBClient(dbConfigs config.DBConfigs, l logger.Logger) (*pg.DB, error) {
	connectionStr := fmt.Sprintf("postgresql://%v:%v/%v?sslmode=disable", dbConfigs.Host, dbConfigs.Port, dbConfigs.DBName)

	l.Info(connectionStr)
	opt, err := pg.ParseURL(connectionStr)
	opt.User = dbConfigs.User
	opt.Password = dbConfigs.Password
	if err != nil {
		return nil, err
	}

	db := pg.Connect(opt)
	db.AddQueryHook(l)

	_, err = db.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}

	return db, nil
}
