package main

import (
	"html/template"
	"net/http"
)

var temp *template.Template
var temp1 *template.Template
var temp2 *template.Template

func init() {
	temp = template.Must(template.ParseGlob("templete/homepage/*.html"))
	temp1 = template.Must(template.ParseGlob("templete/custmer/*.html"))
	temp2 = template.Must(template.ParseGlob("templete/admin/*.html"))
}

func homepage(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "home.html", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "login.html", nil)
}

func signup(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "signup.html", nil)
}

func contactus(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Contactus.html", nil)
}

func weed(w http.ResponseWriter, r *http.Request) {
	temp2.ExecuteTemplate(w, "upload.html", nil)
}

func custmer(w http.ResponseWriter, r *http.Request) {
	temp1.ExecuteTemplate(w, "cusdash.html", nil)
}

func admin(w http.ResponseWriter, r *http.Request) {
	temp2.ExecuteTemplate(w, "admin.html", nil)
}

func deteil(w http.ResponseWriter, r *http.Request) {
	temp2.ExecuteTemplate(w, "deteil.html", nil)
}
