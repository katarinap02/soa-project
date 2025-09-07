package main

import (
	"database-example/handler"
	"database-example/model"
	"database-example/repo"
	"database-example/service"

	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	// // Kreiraj tabelu
	database.AutoMigrate(&model.Student{})
	database.AutoMigrate(&model.User{})
	database.AutoMigrate(&model.UserInfo{})

	// Dodaj test podatke
	database.Exec("INSERT IGNORE INTO students (id, name, major) VALUES ('test-123', 'Marko Markovic', 'Graficki dizajn')")
	database.Exec(`
	  INSERT IGNORE INTO users (id, username, password, email, role, account_status)
	  VALUES (
		'11111111-1111-1111-1111-111111111111',
		'admin',
		'$2a$10$7g9kOKEjN5cYqK0zTGfn/OwWlOE0uR7m2Fz5oRtBoB4gI0KDuNYLa',
		'admin@example.com',
		'Admin',
		'Activated'
	  )
	`) //password je admin123

	database.Exec(`
	  INSERT IGNORE INTO users (id, username, password, email, role, account_status)
	  VALUES (
		'22222222-2222-2222-2222-222222222222',
		'tourist1',
		'$2a$10$wv7vUoW9tc/tXhVo3jkaEuAvMRrKO7CkR9iAnj6UYIh35exg5jXbS',
		'tourist@example.com',
		'Tourist',
		'Activated'
	  )
	`) //sifra je test123

	return database
}

func debugMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("🌐 %s %s - Origin: %s", r.Method, r.URL.Path, r.Header.Get("Origin"))

		// Dodajte CORS headers eksplicitno
		if origin := r.Header.Get("Origin"); origin == "http://localhost:4200" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		// Handle preflight
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// Poveži se sa bazom
	database := initDB()

	// Napravi sve komponente
	studentRepo := &repo.StudentRepository{DatabaseConnection: database}
	studentService := &service.StudentService{StudentRepo: studentRepo}
	studentHandler := &handler.StudentHandler{StudentService: studentService}

	//user
	userRepo := &repo.UserRepository{DatabaseConnection: database}
	userService := &service.UserService{UserRepo: userRepo}
	userHandler := &handler.UserHandler{UserService: userService}

	//userinfo
	userInfoRepo := &repo.UserInfoRepository{DatabaseConnection: database}
	userInfoService := &service.UserInfoService{UserInfoRepo: userInfoRepo}
	userInfoHandler := &handler.UserInfoHandler{UserInfoService: userInfoService}

	// Napravi rute
	router := mux.NewRouter()
	router.HandleFunc("/students/{id}", studentHandler.Get).Methods("GET")
	router.HandleFunc("/students", studentHandler.Create).Methods("POST")
	router.HandleFunc("/users/register", userHandler.Register).Methods("POST")
	router.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
	router.HandleFunc("/profile/{id}", userInfoHandler.GetProfile).Methods("GET")
	router.HandleFunc("/profile/{id}", userInfoHandler.UpdateProfile).Methods("PUT")
	router.HandleFunc("/users/blockuser", userHandler.BlockUser).Methods("POST")

	wrappedRouter := debugMiddleware(router)

	headersOk := handlers.AllowedHeaders([]string{
		"X-Requested-With",
		"Content-Type",
		"Authorization",
		"Accept",
		"Origin",
	})

	originsOk := handlers.AllowedOrigins([]string{
		"http://localhost:4200",
		"http://127.0.0.1:4200",
	})

	methodsOk := handlers.AllowedMethods([]string{
		"GET", "POST", "PUT", "DELETE", "OPTIONS",
	})

	// Dodajte credentials support
	credentialsOk := handlers.AllowCredentials()

	// ISPRAVKA: log.Println IDE PRE log.Fatal
	log.Println("🚀 Server pokrenut na portu 8080")
	log.Println("🔗 CORS omogućen za: http://localhost:4200")

	// Pokretanje servera sa CORS middleware-om
	log.Fatal(http.ListenAndServe(":8080",
		handlers.CORS(originsOk, headersOk, methodsOk, credentialsOk)(wrappedRouter)))

}
