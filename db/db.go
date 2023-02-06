package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tsunagatteru/ishiki-no-nagare/model"
)

func Open()(*sql.DB){
	dbConn, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatalln(err)
	}
	return dbConn
}

func CreateTable(dbConn *sql.DB){
	query := `CREATE TABLE Posts(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    message TEXT,
    updated TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    created TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP);`
	if _, err := dbConn.Exec(query); err != nil {
		log.Fatalln(err)
	}
}

func AddPost(dbConn *sql.DB, message string){
	query := `INSERT INTO Posts (message) VALUES ($1);`
	if _, err := dbConn.Exec(query, message); err != nil {
		log.Println(err)
	}
}

func RetrievePage(dbConn *sql.DB, pageNumber int)([]model.Post){
	query := `SELECT id, message, updated, created
    FROM Posts
    ORDER BY id ASC
    LIMIT 20 OFFSET $1;`
	offset := pageNumber * 20
	rows, err := dbConn.Query(query, offset)
	if  err != nil {
		log.Println(err)
	}
	result := []model.Post{}
	for rows.Next(){
		row := model.Post{}
		if err := rows.Scan(&(row.ID), &(row.Message), &(row.Edited), &(row.Created)); err != nil {
			log.Println(err)
		} else {
			result = append(result, row)
		}
	}
	return result
}
