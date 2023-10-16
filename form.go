package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	database "github.com/pukarlamichhane/project.git/database"
	helper "github.com/pukarlamichhane/project.git/helper"
	"github.com/pukarlamichhane/project.git/model"
)

var jwtKey = []byte("your-secret-key")

func log(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	fname := r.FormValue("email")
	lname := r.FormValue("password")

	var password, role string
	a := database.GetConnection()
	stmt := "SELECT password, role FROM user WHERE email = $1"
	defer a.Close()
	row := a.QueryRow(stmt, fname)
	err := row.Scan(&password, &role)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	if lname == password {
		expirationTime := time.Now().Add(5 * time.Minute)
		claims := &jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Subject:   fname,
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
		switch role {
		case "admin":
			http.Redirect(w, r, "/admindash", http.StatusFound)
		case "user":
			http.Redirect(w, r, "/cusdash", http.StatusFound)
		}
	}

}

func sup(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
		return
	}
	name := r.FormValue("name")
	fname := r.FormValue("email")
	lname := r.FormValue("password")
	name2 := r.FormValue("confirm-password")

	if name2 == lname {
		a := database.GetConnection()
		defer a.Close()
		err := a.Ping()
		if err != nil {
			panic(err.Error())
		}
		stmt, err := a.Prepare("INSERT INTO user (username, email, password, role) VALUES ($1, $2, $3, 'user')")
		if err != nil {
			panic(err.Error())
		}
		defer stmt.Close()

		_, err = stmt.Exec(name, fname, lname)
		if err != nil {
			panic(err.Error())
		}
	}

}

func us(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/contact", http.StatusSeeOther)
		return
	}
	name := r.FormValue("name")
	fname := r.FormValue("email")
	lname := r.FormValue("message")

	helper.Contactmail(lname, name, fname)
	temp.ExecuteTemplate(w, "Contactus.html", nil)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	file, _, err := r.FormFile("photo")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	cloudinaryURL := "cloudinary://146428434594893:kbZyO9xImXjNFO1Qnbv1LUXmw0c@dzi4hcdch"
	cld, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	uploadParams := uploader.UploadParams{
		PublicID: "project",
	}
	uploadResult, err := cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	name := r.FormValue("name")
	description := r.FormValue("description")
	price := r.FormValue("price")

	a := database.GetConnection()
	defer a.Close()
	err = a.Ping()
	if err != nil {
		panic(err.Error())
	}
	stmt, err := a.Prepare("INSERT INTO items (itemname, price, description, imageurl) VALUES ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, price, description, uploadResult.URL)
	if err != nil {
		panic(err.Error())
	}
}

func getdata(w http.ResponseWriter, r *http.Request) {
	db := database.GetConnection()
	defer db.Close()
	ss := []model.Model{}
	s := model.Model{}
	rows, err := db.Query("SELECT itemname, price, description, imageurl FROM items")
	if err != nil {
		fmt.Fprint(w, err)
	} else {
		for rows.Next() {
			rows.Scan(&s.Name, &s.Price, &s.Description, &s.ImgURL)
			ss = append(ss, s)
		}
		json.NewEncoder(w).Encode(ss)
	}

}

func last(w http.ResponseWriter, r *http.Request) {
	s := model.Order{}
	json.NewDecoder(r.Body).Decode(&s)
	helper.Ordermail(s.Name, s.Address, s.Item, s.Quantity, s.Phone)
}

func getall(w http.ResponseWriter, r *http.Request) {
	db := database.GetConnection()
	defer db.Close()
	ss := []model.Model{}
	s := model.Model{}
	rows, err := db.Query("select * from items")
	if err != nil {
		fmt.Fprint(w, err)
	} else {
		for rows.Next() {
			rows.Scan(&s.ID, &s.Name, &s.Price, &s.Description, &s.ImgURL)
			ss = append(ss, s)
		}
		json.NewEncoder(w).Encode(ss)
	}
}

func delete(w http.ResponseWriter, r *http.Request) {
	db := database.GetConnection()
	defer db.Close()
	parms := mux.Vars(r)
	id, _ := strconv.Atoi(parms["id"])
	result, err := db.Exec("delete from items where id=$1", id)
	if err != nil {
		fmt.Fprint(w, err)
	} else {
		_, err := result.RowsAffected()
		if err != nil {
			fmt.Print(w, err)
		}
	}
}

func update(w http.ResponseWriter, r *http.Request) {
	db := database.GetConnection()
	defer db.Close()

	s := model.Model{}
	json.NewDecoder(r.Body).Decode(&s)

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	fmt.Print(s.Name, s.ImgURL, s.Description, s.Price, id)
	_, err = db.Exec("UPDATE items SET itemname=?, imageurl=?, description=?, price=? WHERE id=?", s.Name, s.ImgURL, s.Description, s.Price, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
