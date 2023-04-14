package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"github.com/tsunagatteru/ishiki-no-nagare/model"
)

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

func RetrievePage(dbConn *sql.DB, pageNumber int) []model.Post {
	pageLength := viper.Get("pagelength").(int)
	query := `SELECT id, message, filename, updated, created
    FROM Posts
    ORDER BY id DESC
    LIMIT $1 OFFSET $2;`
	offset := (pageNumber - 1) * pageLength
	rows, err := dbConn.Query(query, pageLength, offset)
	if err != nil {
		log.Println(err)
	}
	result := []model.Post{}
	for rows.Next() {
		row := model.Post{}
		if err := rows.Scan(&(row.ID), &(row.Message), &(row.FileName), &(row.Edited), &(row.Created)); err != nil {
			log.Println(err)
		} else {
			result = append(result, row)
		}
	}
	return result
}

func RetrievePost(dbConn *sql.DB, id int) model.Post {
	query := `SELECT id, message, filename, updated, created
    FROM Posts
    WHERE id=$1`
	row, err := dbConn.Query(query, id)
	if err != nil {
		log.Println(err)
	}
	result := model.Post{}
	row.Next()
	if err := row.Scan(&(result.ID), &(result.Message), &(result.FileName), &(result.Edited), &(result.Created)); err != nil {
		log.Println(err)
	}
	return result
}
