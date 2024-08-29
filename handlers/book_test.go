package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yvann77/bookstore/models"
)

func TestGetBooks(t *testing.T) {
	// Créer un routeur Gin
	router := gin.Default()

	// Enregistrer la route pour GetBooks
	router.GET("/books", GetBooks)

	// Créer une requête de test
	req, err := http.NewRequest(http.MethodGet, "/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Exécuter la requête de test
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Vérifier la réponse
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPostBooks(t *testing.T) {
	// Créer un routeur Gin
	router := gin.Default()

	// Enregistrer la route pour PostBooks
	router.POST("/books", PostBooks)

	// Créer un livre de test
	book := models.Book{
		Title:  "Test Book",
		Author: "Test Author",
		Price:  10.99,
	}

	// Convertir le livre en JSON
	bookJSON, err := json.Marshal(book)
	if err != nil {
		t.Fatal(err)
	}

	// Créer une requête de test
	req, err := http.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(bookJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Exécuter la requête de test
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Vérifier la réponse
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetBookByID(t *testing.T) {
	// Créer un routeur Gin
	router := gin.Default()

	// Enregistrer la route pour GetBookByID
	router.GET("/books/:id", GetBookByID)

	// Créer une requête de test
	req, err := http.NewRequest(http.MethodGet, "/books/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Exécuter la requête de test
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Vérifier la réponse
	assert.Equal(t, http.StatusOK, w.Code)
}
