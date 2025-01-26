package main

import (
    "database/sql"
    "fmt"

    _ "modernc.org/sqlite"
)

func main() {
    db, err := sql.Open("sqlite", "db_name.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()
} 