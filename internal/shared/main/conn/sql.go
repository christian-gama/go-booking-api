package conn

import (
	_sql "database/sql"
	"fmt"

	"github.com/christian-gama/go-booking-api/internal/shared/infra/configger"
)

type sql struct {
	dbConfigger configger.Db

	db         *_sql.DB
	driverName string
	dsn        string
}

func (s *sql) open() error {
	s.setDsn()

	db, err := _sql.Open(s.driverName, s.dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	s.db = db

	return nil
}

func (s *sql) setup() {
	s.db.SetMaxIdleConns(s.dbConfigger.MaxIdleConnections())
	s.db.SetMaxOpenConns(s.dbConfigger.MaxConnections())
	s.db.SetConnMaxLifetime(s.dbConfigger.MaxLifeTimeMin())
}

func (s *sql) setDsn() {
	s.dsn = fmt.Sprintf(`
	host=%s
	port=%d
	dbname=%s
	user=%s
	password=%s
	sslmode=%s`,
		s.dbConfigger.Host(),
		s.dbConfigger.Port(),
		s.dbConfigger.Name(),
		s.dbConfigger.User(),
		s.dbConfigger.Password(),
		s.dbConfigger.SslMode(),
	)
}

func NewSQL(driverName string, dbConfigger configger.Db) (*_sql.DB, error) {
	sql := &sql{
		dbConfigger: dbConfigger,
		driverName:  driverName,
	}

	if err := sql.open(); err != nil {
		return nil, err
	}

	sql.setup()

	return sql.db, nil
}
