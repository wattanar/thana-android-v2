package main

import (
	"github.com/julienschmidt/httprouter"
	c "github.com/wattanar/thana-android/controllers"
	"log"
	"net/http"
)

func main() {
	// Create New Router
	r := httprouter.New()
	// Router
	r.GET("/", c.Landing)
	r.GET("/t/:id", c.Movies)
	r.GET("/theater", c.Theaters)
	// Setup Static Path
	r.NotFound = http.StripPrefix("/static/", http.FileServer(http.Dir("./assets")))
	// Listening on port 8080
	log.Fatal(http.ListenAndServe(":8081", r))
}
