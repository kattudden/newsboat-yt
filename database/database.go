package database

import (
	"database/sql"
	"fmt"
	"kattudden/newsboat-yt/config"

	"github.com/google/uuid"
    _"github.com/mattn/go-sqlite3"
)

var conf, _ = config.New()

func getDatabase() (*sql.DB, error){
    db, err := sql.Open("sqlite3", conf.DatabasePath)
    if err != nil {
        panic(err)
    }
    return db, nil
}


func createTable(db *sql.DB) {
    sqlStatement := "CREATE TABLE IF NOT EXISTS queue (id TEXT not null primary key, link TEXT not null, downloaded BOOLEAN)"
    _, err := db.Exec(sqlStatement)
    if err != nil {
        panic(err)
    }
}


func InsertUrl(url string) error {
    db, err := getDatabase()
    if err != nil {
        return err
    }

    createTable(db) 

    uuid := uuid.NewString()
    sqlStatement := fmt.Sprintf("INSERT INTO queue (id, link, downloaded) VALUES (?, ?, ?)")

    tx, err := db.Begin()
    if err != nil {
        return err
    }

    stmt, err := tx.Prepare(sqlStatement)
    if err != nil {
            return err
    }

    defer stmt.Close()

    _, err = stmt.Exec(uuid, url, false)
    if err != nil {
        return err
    }

    tx.Commit()
    db.Close()

    return nil
}


type Url struct {
    ID string
    Link string
}


func GetUrls() ([]Url){
    db, _ := getDatabase()

    queryStatement := "SELECT id, link from queue WHERE downloaded = 0"
    rows, _ := db.Query(queryStatement)

    var urls []Url

    for rows.Next() {
        var id string
        var link string

        rows.Scan(&id, &link)

        entry := Url{
            ID: id,
            Link: link,
        }

        urls = append(urls, entry)
    }

    db.Close()

    return urls
}


func MarkUrlDownloaded(id string) error {
    db, err := getDatabase()
    if err != nil {
        return err
    }

    sqlStatement := `UPDATE queue SET downloaded = true WHERE id = ?`

    tx, err := db.Begin()
    if err != nil {
        return err
    }

    stmt, err := tx.Prepare(sqlStatement)
    if err != nil {
            return err
    }

    defer stmt.Close()

    _, err = stmt.Exec(id)
    if err != nil {
        return err
    }

    tx.Commit()
    db.Close()

    return nil
}
