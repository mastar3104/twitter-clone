package adapter

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type MySQLHandler struct {
	username string
	password string
	host     string
	port     string
	dBName   string
	db       *sql.DB
}

func (m *MySQLHandler) Connect() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", m.username, m.password, m.host, m.port, m.dBName))
	if err != nil {
		panic(err.Error())
	}
	m.db = db
}

func (m *MySQLHandler) Exec(query string, args ...any) {
	_, err := m.db.Exec(query, args...)
	if err != nil {
		panic(err.Error())
	}
}

func (m *MySQLHandler) Query(query string, args ...any) *sql.Rows {
	rows, err := m.db.Query(query, args...)
	if err != nil {
		panic(err.Error())
	}
	return rows
}

func (m *MySQLHandler) Close() {
	if m.db != nil {
		err := m.db.Close()
		if err != nil {
			panic(err.Error())
		}
	}
}

func GetDatabaseHandler() *MySQLHandler {
	mySQLHandler := &MySQLHandler{
		username: os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		dBName:   os.Getenv("DB_DATABASE"),
	}
	mySQLHandler.Connect()
	return mySQLHandler
}
