package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"path/filepath"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

// Response structure to send back JSON responses
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Request structure to parse incoming JSON data
type Request struct {
	Message string `json:"message"`
}

// HandlePost handles POST requests
func HandlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if req.Message == "" {
		response := Response{
			Status:  "fail",
			Message: "Invalid JSON message",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	fmt.Println("Received message:", req.Message)

	response := Response{
		Status:  "success",
		Message: "Data successfully received",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// HandleGet handles GET requests
func HandleGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	response := Response{
		Status:  "success",
		Message: "GET request received",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ServeStaticHTML(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("static", "handlepost.html")
	http.ServeFile(w, r, filePath)
}

func main() {

	// Connect database
	dsn := "user=postgres password=root dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})

	// Create
	db.Create(&User{Name: "Ismail", Email: "a@example.com"})

	// Read
	var user User
	db.First(&user, 1)
	log.Println(user)

	// Update
	db.Model(&user).Update("Name", "Ayupov")

	// Delete
	db.Delete(&user, 1)

	// Run server
	http.HandleFunc("/post", HandlePost)
	http.HandleFunc("/get", HandleGet)

	// API routes
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	http.HandleFunc("/postmessage", ServeStaticHTML)

	port := ":8080"
	fmt.Println("Server is running on port" + port)

	// Start server on port specified above
	log.Fatal(http.ListenAndServe(port, nil))

}
