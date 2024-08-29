package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yvann77/bookstore/database"
	"github.com/yvann77/bookstore/routes"
)

func main() {
	// Initialiser le routeur Gin
	router := gin.Default()

	// Connecter à la base de données
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Enregistrer les routes
	routes.SetupBookRoutes(router)

	// Démarrer le serveur
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server started on port %s\n", port)
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatal(err)
	}
}
