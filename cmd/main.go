package main

import (
	"log"
	"net/http"

	"github.com/TomPo62/bakiverse-backend/internal/pkg/cors"
	"github.com/TomPo62/bakiverse-backend/internal/pkg/database"

	"github.com/joho/godotenv" // Pour charger les variables d'environnement
)

func main() {
	// Charger les variables d'environnement
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erreur lors du chargement des variables d'environnement :", err)
	}

	// 1. Connexion à la base de données MariaDB avec des variables d'env
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Erreur de connexion à la base de données :", err)
	}
	// Appliquer le middleware CORS à toutes les routes
	handlerWithCORS := infrastructure.CORSMiddleware(finalMux)

	// 7. Démarrer le serveur HTTP
	log.Println("Serveur démarré sur le port 8080")
	log.Fatal(http.ListenAndServe(":8080", handlerWithCORS))
}
