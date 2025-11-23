package database

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

func ConnectPostgres() (*sql.DB, error) {
    dsn := "host=localhost user=postgres password=yourPassword dbname=uas_be port=5432 sslmode=disable"
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, err
    }
    return db, db.Ping()
}
