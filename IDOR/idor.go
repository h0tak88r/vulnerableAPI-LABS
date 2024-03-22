package idor

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Score int    `json:"score"`
}

var Users = []User{
	{
		ID:    0,
		Name:  "Mosaad",
		Score: 88,
	},
	{
		ID:    1,
		Name:  "Sallam",
		Score: 100,
	},
	{
		ID:    2,
		Name:  "Ahmed",
		Score: 50,
	},
}

func main() {
	// Define HTTP endpoints and handlers
	http.HandleFunc("/api/user", userEndpoint)
	http.HandleFunc("/", index)

	fmt.Println("Starting api server")
	// Start HTTP server on port 8080
	http.ListenAndServe(":8080", nil)
}

// Handle requests to root endpoint
func index(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Handling Request from user to the index")
	// Set response content type
	w.Header().Set("Content-Type", "application/json")
	// Send welcome message and endpoints documentation
	fmt.Fprintf(w, `{"message": "Welcome to Sallam's IDOR-Vulnerable Api", "endpoints": {"/api/user": {"methods": ["GET", "POST"], "description": "Endpoint to fetch or add users. GET request accepts optional 'id' parameter. Example: GET /api/user?id=1. POST request accepts a JSON object representing the user. Example: POST /api/user. Body: {"id":88,"name":"Ahmed","score":10}"}}}`)
}

// Handle requests to /api/user endpoint
func userEndpoint(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		get(w, req) // Handle GET requests
	case "POST":
		post(w, req) // Handle POST requests
	default:
		// Return Method Not Allowed error for unsupported HTTP methods
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

// Handle GET requests to /api/user endpoint
func get(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Handling Get Request /api/user")
	// Extract query parameters
	query := req.URL.Query()
	id := query.Get("id")

	var response []byte
	var err error

	if id == "" {
		// If no ID provided, return all users
		response, err = json.Marshal(Users)
	} else {
		// If ID provided, return user with that ID
		intID, err := strconv.Atoi(id)
		if err == nil {
			for _, user := range Users {
				if user.ID == intID {
					response, err = json.Marshal(user)
					break
				}
			}
			if err != nil {
				err = fmt.Errorf("user not found")
			}
		}
	}

	// error handling
	handleError(w, err)
	if err == nil {
		// If no error, set response content type and send JSON response
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", response)
	}
}

// Handle POST requests to /api/user endpoint
func post(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Handling Post Request to /api/user")

	var u User             // declare u as type of user object
	defer req.Body.Close() // make sure to close the request body

	// Decode the json request body
	if err := json.NewDecoder(req.Body).Decode(&u); err != nil {
		handleError(w, err)
		return
	}

	// Add the new user to the Users slice
	Users = append(Users, u)
	w.WriteHeader(http.StatusCreated)

	// Return the new user as response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

// Handle errors by sending appropriate HTTP error response
func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}
}
