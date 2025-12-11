package db
import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)
var DB *sql.DB

func NewPostgresConn(host, user, password, dbname, port string) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	DB = db
	return db, nil
}