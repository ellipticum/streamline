package postgres

import (
	"crypto/tls"
	"database/sql"
	"fmt"
)

type Connector struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
	TLS      *tls.Config
}

func (c *Connector) Open() (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", c.User, c.Password, c.Host, c.Port, c.Database, c.TLSMode())

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func (c *Connector) TLSMode() string {
	if c.TLS != nil {
		return "disable"
	}

	return "require"
}
