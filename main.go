package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	middleware "github.com/pukarlamichhane/project.git/middleware"
)

func main() {
	r := mux.NewRouter()
	r.Use(middleware.CorsMiddleware)
	staticFileDirectory := http.Dir("./static/")
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")
	r.HandleFunc("/", homepage)
	r.HandleFunc("/login", login)
	r.HandleFunc("/signup", signup)
	r.HandleFunc("/contact", contactus)
	r.HandleFunc("/log", log).Methods("POST")
	r.HandleFunc("/sup", sup).Methods("POST")
	r.HandleFunc("/us", us).Methods("POST")
	r.HandleFunc("/submit", middleware.Auth(http.HandlerFunc(weed)))
	r.HandleFunc("/cusdash", middleware.Auth(http.HandlerFunc(custmer)))
	r.HandleFunc("/admindash", middleware.Auth(http.HandlerFunc(admin)))
	r.HandleFunc("/orders", middleware.Auth(http.HandlerFunc(last)))
	r.HandleFunc("/upload", middleware.Auth(http.HandlerFunc(Upload))).Methods("POST")
	r.HandleFunc("/get", getdata).Methods("GET")
	r.HandleFunc("/getall", middleware.Auth(http.HandlerFunc(getall))).Methods("GET")
	r.HandleFunc("/items/{id}", middleware.Auth(http.HandlerFunc(delete))).Methods("DELETE")
	r.HandleFunc("/deteil", deteil)
	r.HandleFunc("/update/{id}", update).Methods("PUT")

	// Read port from environment variables, default to 0.0.0.0:9999
	port := os.Getenv("PORT")
	if port == "" {
		port = "0.0.0.0"
	}
	http.ListenAndServe(port, r)
}
