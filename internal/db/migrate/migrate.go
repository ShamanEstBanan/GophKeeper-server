package migrate

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
)

type options struct {
	driverName string
	path       string
}

type OptionsFunc func(opts *options)

func WithPath(path string) OptionsFunc {
	return func(opts *options) {
		opts.path = path
	}
}

func Run(dsn string, opts ...OptionsFunc) error {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}
	curDriverName := "pgx"
	if o.driverName != "" {
		curDriverName = o.driverName
	}
	sqlDB, err := sql.Open(curDriverName, dsn)
	if err != nil {
		return fmt.Errorf("open database connection error: %w ", err)
	}
	defer func() { _ = sqlDB.Close() }()

	var curPath string
	if o.path != "" {
		curPath = o.path
	}

	if err = goose.Up(sqlDB, curPath); err != nil {
		return fmt.Errorf("up migrations: %w", err)
	}
	return nil
}
