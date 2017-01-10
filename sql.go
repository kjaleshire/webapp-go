package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/unrolled/render.v1"
)

type Book struct {
	Title  string
	Author string
}

func main() {
	db := NewDB()
	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", ShowBooks(db))
}

func ShowBooks(db *sql.DB) http.Handler {
	renderer := render.New(render.Options{})
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var books []*Book

		rows, err := db.Query("SELECT title, author FROM books")
		if err != nil {
			panic(err)
		}

		for rows.Next() {
			book := new(Book)
			if err := rows.Scan(&book.Title, &book.Author); err != nil {
				panic(err)
			}

			books = append(books, book)
		}

		renderer.HTML(rw, http.StatusOK, "books", books)
	})
}

func NewDB() *sql.DB {
	db, err := sql.Open("sqlite3", "example.sqlite")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS books(title text, author text)")
	if err != nil {
		panic(err)
	}

	return db
}
