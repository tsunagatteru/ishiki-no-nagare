package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Post struct {
	ID       int    `json:"id"`
	Message  string `json:"message"`
	FileName string `json:"filename"`
	Edited   string `json:"edited"`
	Created  string `json:"created"`
}

func Open(dataPath string) *sql.DB {
	dbConn, err := sql.Open("sqlite3", dataPath+"data.db")
	if err != nil {
		log.Fatalln(err)
	}
	return dbConn
}

func CreateTable(dbConn *sql.DB) {
	query := `CREATE TABLE if not exists Posts(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    message TEXT,
    filename TEXT,
    updated TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    created TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP);`
	if _, err := dbConn.Exec(query); err != nil {
		log.Fatalln(err)
	}
}

func AddPost(dbConn *sql.DB, message string, filename string) {
	query := `INSERT INTO Posts (message, filename) VALUES ($1, $2);`
	if _, err := dbConn.Exec(query, message, filename); err != nil {
		log.Println(err)
	}
}

func RetrievePage(dbConn *sql.DB, pageNumber int, pageLength int) []Post {
	query := `SELECT id, message, filename, updated, created
    FROM Posts
    ORDER BY id DESC
    LIMIT $1 OFFSET $2;`
	offset := (pageNumber - 1) * pageLength
	rows, err := dbConn.Query(query, pageLength, offset)
	if err != nil {
		log.Println(err)
	}
	result := []Post{}
	for rows.Next() {
		row := Post{}
		if err := rows.Scan(&(row.ID), &(row.Message), &(row.FileName), &(row.Edited), &(row.Created)); err != nil {
			log.Println(err)
		} else {
			result = append(result, row)
		}
	}
	return result
}

func RetrievePost(dbConn *sql.DB, id int) Post {
	query := `SELECT id, message, filename, updated, created
    FROM Posts
    WHERE id=$1`
	row, err := dbConn.Query(query, id)
	if err != nil {
		log.Println(err)
	}
	result := Post{}
	row.Next()
	if err := row.Scan(&(result.ID), &(result.Message), &(result.FileName), &(result.Edited), &(result.Created)); err != nil {
		log.Println(err)
	}
	return result
}
