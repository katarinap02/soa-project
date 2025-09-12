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
	//"github.com/rs/cors"
)

func initDB() *gorm.DB {
	connectionStr := "root:root@tcp(database:3306)/soadb?charset=utf8mb4&parseTime=True&loc=Local"
	//connectionStr := "root:root@tcp(localhost:3306)/soadb?charset=utf8mb4&parseTime=True&loc=Local"
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
	database.AutoMigrate(&model.Comment{})
	database.AutoMigrate(&model.BlogLike{})

	// Dodaj test podatke
	//database.Exec("INSERT IGNORE INTO BlogPost (id, name, major) VALUES ('test-123', 'Marko Markovic', 'Graficki dizajn')")

	/*
		database.Exec(`
		  INSERT INTO blog_posts (id, username, title, description, date)
		  VALUES ('11111111-1111-1111-1111-111111111111', 'user1', 'My First Post', 'This is the description', NOW())
		`)
	*/
	return database
}

func main() {
	// Poveži se sa bazom
	database := initDB()

	// Napravi sve komponente
	blogPostRepo := &repo.BlogPostRepository{DatabaseConnection: database}
	blogPostService := &service.BlogPostService{BlogPostRepo: blogPostRepo}
	blogPostHandler := &handler.BlogHandler{BlogPostService: blogPostService}

	commentRepo := &repo.CommentRepository{DatabaseConnection: database}
	commentService := &service.CommentService{CommentRepo: commentRepo}
	commentHandler := &handler.CommentHandler{CommentService: commentService}

	// Napravi rute
	router := mux.NewRouter()
	router.HandleFunc("/blog/create-post", blogPostHandler.CreateBlogPost).Methods("POST", "OPTIONS")
	router.HandleFunc("/blog/create-comment", commentHandler.CreateComment).Methods("POST", "OPTIONS")
	router.HandleFunc("/blog/like-blog", blogPostHandler.CreateBlogLike).Methods("POST")
	router.HandleFunc("/blog/unlike-blog", blogPostHandler.DeleteBlogLike).Methods("POST")
	router.HandleFunc("/blog", blogPostHandler.GetAllBlogPosts).Methods("GET")
	router.HandleFunc("/blog/by-username", blogPostHandler.GetBlogPostsByUsername).Methods("GET")

	// Pokretanje servera

	/*	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})*/
	//handler := c.Handler(router)
	log.Println("Server pokrenut na portu 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
