package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type Database interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Close() error
}

type PostgresDB struct {
	db *sql.DB
}

func NewPostgresDB(config DBConfig) (*PostgresDB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &PostgresDB{db: db}, nil
}

func (db *PostgresDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	fmt.Println(query, "\n", args)
	return db.db.Query(query, args...)
}

func (db *PostgresDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.db.Exec(query, args...)
}

func (db *PostgresDB) Close() error {
	return db.db.Close()
}
