package main

import (
	"fmt"
	"log"
	"simba/album-store-api/album"
	"simba/album-store-api/handler"


	"simba/album-store-api/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// var albums = []album{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

func main() {
	loadEnv()
	loadDatabase()

	albumRepository := album.NewAlbumRepository(database.Database)
	albumService := album.NewAlbumService(albumRepository)
	albumController := handler.NewAlbumController(albumService)

	serveApplication(albumController)
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&album.Album{})
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func serveApplication(albumController *handler.AlbumController) {
	router := gin.Default()

	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.GET("/albums", albumController.GetAllAlbums)
	router.GET("/albums/:id", albumController.GetAlbumById)
	router.POST("/albums", albumController.CreateAlbum)
	router.DELETE("/albums/:id", albumController.DeleteAlbum)
	router.PUT("/albums/:id", albumController.UpdateAlbum)

	router.Run("localhost:8080")
	fmt.Println("Server running on port 8080")
}