package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yvann77/bookstore/models"
)

// Définir une interface pour les interactions avec la base de données
type BookRepository interface {
	GetAllBooks() ([]models.Book, error)
	AddBook(book models.Book) error
	GetBookByID(id string) (*models.Book, error)
}

// Créer une structure qui implémente l'interface BookRepository
type MockBookRepository struct {
}

// Implémenter les méthodes de l'interface BookRepository
func (r *MockBookRepository) GetAllBooks() ([]models.Book, error) {
	// Retourner des données de test
	return []models.Book{
		{Title: "Test Book 1", Author: "Test Author 1", Price: 10.99},
		{Title: "Test Book 2", Author: "Test Author 2", Price: 19.99},
	}, nil
}

func (r *MockBookRepository) AddBook(book models.Book) error {
	// Simuler l'ajout d'un livre à la base de données
	return nil
}

func (r *MockBookRepository) GetBookByID(id string) (*models.Book, error) {
	// Retourner un livre de test
	return &models.Book{Title: "Test Book 1", Author: "Test Author 1", Price: 10.99}, nil
}

// Modifier les fonctions dans handlers.go pour utiliser l'interface BookRepository
func GetBooks(c *gin.Context) {
	// Récupérer le repository de livres
	repo := c.MustGet("bookRepo").(BookRepository)

	// Récupérer tous les livres depuis le repository
	books, err := repo.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

func PostBooks(c *gin.Context) {
	// Décoder le corps de la requête JSON
	var book models.Book
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Récupérer le repository de livres
	repo := c.MustGet("bookRepo").(BookRepository)

	// Ajouter le livre au repository
	err := repo.AddBook(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book created successfully"})
}

func GetBookByID(c *gin.Context) {
	// Récupérer l'ID du livre depuis l'URL
	id := c.Param("id")

	// Récupérer le repository de livres
	repo := c.MustGet("bookRepo").(BookRepository)

	// Récupérer le livre depuis le repository
	book, err := repo.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

// DatabaseBookRepository implémente l'interface BookRepository en utilisant une base de données
type DatabaseBookRepository struct {
	DB *sql.DB
}

// Implémenter les méthodes de l'interface BookRepository pour DatabaseBookRepository
func (r *DatabaseBookRepository) GetAllBooks() ([]models.Book, error) {
	rows, err := r.DB.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (r *DatabaseBookRepository) AddBook(book models.Book) error {
	_, err := r.DB.Exec("INSERT INTO books (title, author, price) VALUES (?, ?, ?)", book.Title, book.Author, book.Price)
	if err != nil {
		return err
	}

	return nil
}

func (r *DatabaseBookRepository) GetBookByID(id string) (*models.Book, error) {
	var book models.Book
	err := r.DB.QueryRow("SELECT * FROM books WHERE id = ?", id).Scan(&book.ID, &book.Title, &book.Author, &book.Price)
	if err != nil {
		return nil, err
	}

	return &book, nil
}
