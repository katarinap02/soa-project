package main

import (
	"database-example/handler"
	"database-example/model"
	"database-example/repo"
	"database-example/service"

	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	//connectionStr := "root:root@tcp(database:3306)/soadb?charset=utf8mb4&parseTime=True&loc=Local"
	connectionStr := "root:root@tcp(localhost:3306)/soadb?charset=utf8mb4&parseTime=True&loc=Local"
	var database *gorm.DB
	var err error

	// Pokušaj povezivanja sa bazom (max 20 pokušaja)
	for i := 0; i < 20; i++ {
		database, err = gorm.Open(mysql.Open(connectionStr), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Pokušaj %d/20 - čekam bazu...", i+1)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatal("Ne mogu da se povežem sa bazom:", err)
	}

	log.Println("Uspešno povezan sa bazom")

	// Kreiraj tabelu
	database.AutoMigrate(&model.BlogPost{})

	// Dodaj test podatke
	//database.Exec("INSERT IGNORE INTO BlogPost (id, name, major) VALUES ('test-123', 'Marko Markovic', 'Graficki dizajn')")

	return database
}

func main() {
	// Poveži se sa bazom
	database := initDB()

	// Napravi sve komponente
	blogPostRepo := &repo.BlogPostRepository{DatabaseConnection: database}
	blogPostService := &service.BlogPostService{BlogPostRepo: blogPostRepo}
	blogPostHandler := &handler.BlogHandler{BlogPostService: blogPostService}

	// Napravi rute
	router := mux.NewRouter()
	router.HandleFunc("/blog/create-post", blogPostHandler.CreateBlogPost).Methods("POST")

	// Pokretanje servera
	log.Println("Server pokrenut na portu 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
