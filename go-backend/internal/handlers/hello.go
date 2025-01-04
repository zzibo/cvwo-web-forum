package handlers

import(

	"net/http"
	"encoding/json"
	
)


type Response struct {
	Message string `json:"message"`
}


func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// Handle CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// Create a response object
	response := Response{
		Message: "Connection successful! Hello from the backend.",
	}

	// Set the Content-Type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	json.NewEncoder(w).Encode(response)
}
