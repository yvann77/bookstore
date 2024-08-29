package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/your-username/bookstore/models"
)

func GetBooks(c *gin.Context) {
	// Récupérer tous les livres depuis la base de données
	books, err := models.GetAllBooks()
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

	// Ajouter le livre à la base de données
	err := models.AddBook(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book created successfully"})
}

func GetBookByID(c *gin.Context) {
	// Récupérer l'ID du livre depuis l'URL
	id := c.Param("id")

	// Récupérer le livre depuis la base de données
	book, err := models.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}
