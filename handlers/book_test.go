package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yvann77/bookstore/models"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	// Force Gin en mode release
	gin.SetMode(gin.ReleaseMode)

	// Créer une seule instance de Gin
	router = gin.Default()

	// Enregistrer les routes
	router.GET("/books", GetBooks)
	router.POST("/books", PostBooks)
	router.GET("/books/:id", GetBookByID)

	// Exécuter les tests
	code := m.Run()

	// Arrêter le serveur Gin
	// (Pas nécessaire dans ce cas, mais utile si tu utilises un serveur réel)
	// router.Engine.Close()

	os.Exit(code)
}

func TestGetBooks(t *testing.T) {
	// Créer un mock du repository de livres
	mockRepo := &MockBookRepository{}

	// Ajouter le mock au contexte de Gin *dans le handler de test*
	router.Use(func(c *gin.Context) {
		c.Set("bookRepo", mockRepo)
		c.Next()
	})

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
	// Créer un mock du repository de livres
	mockRepo := &MockBookRepository{}

	// Ajouter le mock au contexte de Gin *dans le handler de test*
	router.Use(func(c *gin.Context) {
		c.Set("bookRepo", mockRepo)
		c.Next()
	})

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
	// Créer un mock du repository de livres
	mockRepo := &MockBookRepository{}

	// Ajouter le mock au contexte de Gin *dans le handler de test*
	router.Use(func(c *gin.Context) {
		c.Set("bookRepo", mockRepo)
		c.Next()
	})

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
