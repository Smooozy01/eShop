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

// Response structure to send back JSON responses
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type User struct {
	gorm.Model
	Name   string `json:"name"`
	Email  string `json:"email"`
	Method string `json:"method"`
	ID     int    `json:"id"`
}

func main() {

	// Run server
	http.HandleFunc("/post", HandlePost)
	http.HandleFunc("/get", HandleGet)

	// API routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/about.html")
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	http.HandleFunc("/message", ServeStaticHTML)

	port := ":8080"
	fmt.Println("Server is running on port" + port)

	// Start server on port specified above
	log.Fatal(http.ListenAndServe(port, nil))

}

func createDB(name, email string, db *gorm.DB) bool {
	result := db.Create(&User{Name: name, Email: email})
	return result.Error == nil
}

func readDB(id int, db *gorm.DB) (bool, []User) {
	var users []User
	result := db.Where("id = ?", id).First(&users)
	if len(users) != 0 {
		fmt.Println(users[0])
	}
	return result.Error == nil, users
}

func updateDB(id int, name string, db *gorm.DB) bool {
	result := db.Model(&User{}).Where("id = ?", id).Updates(User{Name: name})
	if result.RowsAffected == 0 {
		return false
	}
	return result.Error == nil
}

func deleteDB(id int, db *gorm.DB) bool {
	result := db.Where("id = ?", id).Delete(&User{})
	if result.RowsAffected == 0 {
		return false
	}
	return result.Error == nil
}

func readAllDB(db *gorm.DB) (bool, []User) {
	var users []User
	result := db.Find(&users)
	return result.Error == nil, users
}

// HandlePost handles POST requests
func HandlePost(w http.ResponseWriter, r *http.Request) {

	dsn := "user=postgres password=root dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req User
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	if (req.Email == "" || req.Name == "") && req.Method == "create" {
		response := Response{
			Status:  "fail",
			Message: "Invalid JSON message",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	fmt.Println("Received message:", req.Name, req.Email, req.Method)

	if req.Method == "create" {
		success := createDB(req.Name, req.Email, db)
		w.Header().Set("Content-Type", "application/json")
		if success {
			json.NewEncoder(w).Encode("Successfully created user")
		} else {
			json.NewEncoder(w).Encode("Failed to create user")
		}

	} else if req.Method == "read" {
		success, matches := readDB(int(req.ID), db)
		w.Header().Set("Content-Type", "application/json")
		if success && matches != nil {
			json.NewEncoder(w).Encode(matches)
		} else if success && matches == nil {
			json.NewEncoder(w).Encode("No user found")
		} else {
			json.NewEncoder(w).Encode("Failed to read user")
		}

	} else if req.Method == "update" {
		success := updateDB(int(req.ID), req.Name, db)
		w.Header().Set("Content-Type", "application/json")
		if success {
			json.NewEncoder(w).Encode("Successfully updated user")
		} else {
			json.NewEncoder(w).Encode("Failed to update user")
		}

	} else if req.Method == "delete" {
		success := deleteDB(int(req.ID), db)
		w.Header().Set("Content-Type", "application/json")
		if success {
			json.NewEncoder(w).Encode("Successfully deleted user")
		} else {
			json.NewEncoder(w).Encode("Failed to delete user")
		}

	} else if req.Method == "getRecords" {
		success, matches := readAllDB(db)
		w.Header().Set("Content-Type", "application/json")
		if success && matches != nil {
			json.NewEncoder(w).Encode(matches)
		} else if success && matches == nil {
			json.NewEncoder(w).Encode("No user found")
		} else {
			json.NewEncoder(w).Encode("Failed to read users")
		}
	}

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

func ServeStaticCSS(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("static", "style.css")
	http.ServeFile(w, r, filePath)
}
