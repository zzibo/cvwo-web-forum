package main

import (

    "fmt"
	"log"
	"net/http"

	"github.com/zzibo/cvwo-web-forum/internal/database"
	"github.com/zzibo/cvwo-web-forum/internal/router"
)

func main() {

	db, err := database.GetDB()
	if err != nil{
		log.Fatalf("Failed to connect to the database: %v", err)
	}

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
	 
	r := router.Setup()
	fmt.Print("Listening on port 8000 at http://localhost:8000!")
	log.Fatalln(http.ListenAndServe(":8000", r))
}
