package models

import (
	"github.com/yvann77/bookstore/database"
)

// Structure représentant un livre
type Book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

// Fonction pour récupérer tous les livres de la base de données
func GetAllBooks() ([]Book, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

// Fonction pour ajouter un nouveau livre à la base de données
func AddBook(book Book) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO books (title, author, price) VALUES (?, ?, ?)", book.Title, book.Author, book.Price)
	if err != nil {
		return err
	}

	return nil
}

// Fonction pour récupérer un livre par son ID de la base de données
func GetBookByID(id string) (Book, error) {
	db, err := database.Connect()
	if err != nil {
		return Book{}, err
	}
	defer db.Close()

	var book Book
	err = db.QueryRow("SELECT * FROM books WHERE id = ?", id).Scan(&book.ID, &book.Title, &book.Author, &book.Price)
	if err != nil {
		return Book{}, err
	}

	return book, nil
}
