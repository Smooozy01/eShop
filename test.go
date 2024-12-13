package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Request structure for sending JSON data
type request struct {
	Message string `json:"message"`
}

func main() {
	// Test POST request
	testPost()

	// Test GET request
	testGet()
}

func testPost() {
	url := "http://localhost:8080/post"

	// Create JSON data
	data := request{
		Message: "Lalalalal.",
	}

	// Serialize the data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Make the POST request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	defer resp.Body.Close()

	// Read and display the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("POST Response:")
	fmt.Println(string(body))
}

func testGet() {
	url := "http://localhost:8080/get"

	// Make the GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer resp.Body.Close()

	// Read and display the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("GET Response:")
	fmt.Println(string(body))
}
