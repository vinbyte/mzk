package infras

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/vinbyte/mzk/configs"
)

// PostgresConn wraps a pair of read/write MySQL connections.
type PostgresConn struct {
	Read  *sqlx.DB
	Write *sqlx.DB
}

// ProvidePostgresConn is the provider for PostgresConn.
func ProvidePostgresConn(config *configs.Config) *PostgresConn {
	return &PostgresConn{
		Read:  CreatePostgresReadConn(config),
		Write: CreatePostgresWriteConn(config),
	}
}

// CreatePostgresWriteConn creates a database connection for write access.
func CreatePostgresWriteConn(config *configs.Config) *sqlx.DB {
	return createDBConnection(
		"write",
		config.Database.Write.User,
		config.Database.Write.Password,
		config.Database.Write.Host,
		config.Database.Write.Port,
		config.Database.Write.Name,
		config.Database.Write.SslMode,
		config.Database.Write.MaxIdleConn,
		config.Database.Write.MaxOpenConn,
	)
}

// CreatePostgresReadConn creates a database connection for read access.
func CreatePostgresReadConn(config *configs.Config) *sqlx.DB {
	return createDBConnection(
		"read",
		config.Database.Read.User,
		config.Database.Read.Password,
		config.Database.Read.Host,
		config.Database.Read.Port,
		config.Database.Read.Name,
		config.Database.Read.SslMode,
		config.Database.Read.MaxIdleConn,
		config.Database.Read.MaxOpenConn,
	)
}

// createDBConnection creates a database connection.
func createDBConnection(connType, user, password, host string, port int, dbName, sslMode string, maxIdleConn, maxOpenConn int) *sqlx.DB {
	pqURL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s  sslmode=%s",
		host,
		port,
		user,
		password,
		dbName,
		sslMode,
	)

	db, err := sqlx.Connect("postgres", pqURL)
	if err != nil {
		log.
			Fatal().
			Err(err).
			Str("dbType", connType).
			Str("host", host).
			Int("port", port).
			Str("dbName", dbName).
			Msg("Failed connecting to database")
	}
	db.SetMaxIdleConns(maxIdleConn)
	db.SetMaxOpenConns(maxOpenConn)

	err = db.Ping()
	if err != nil {
		log.
			Fatal().
			Err(err).
			Str("dbType", connType).
			Str("host", host).
			Int("port", port).
			Str("dbName", dbName).
			Msg("Failed ping database")
	}

	log.
		Info().
		Str("dbType", connType).
		Str("host", host).
		Int("port", port).
		Str("dbName", dbName).
		Msg("Connected to database")

	return db
}
