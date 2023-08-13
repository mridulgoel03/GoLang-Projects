package main

import (
	"fmt"
	"net/http"
	"math/rand"
	"time"

	"github.com/gorilla/mux"
)

var (
	shortToLong = make(map[string]string)
	longToShort = make(map[string]string)
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func generateShortURL() string {
	rand.Seed(time.Now().UnixNano())
	shortURL := ""
	for i := 0; i < 6; i++ {
		shortURL += string(letters[rand.Intn(len(letters))])
	}
	return shortURL
}

func createShortURL(w http.ResponseWriter, r *http.Request) {
	longURL := r.FormValue("url")
	if longURL == "" {
		http.Error(w, "Missing URL", http.StatusBadRequest)
		return
	}

	if _, ok := longToShort[longURL]; !ok {
		shortURL := generateShortURL()
		shortToLong[shortURL] = longURL
		longToShort[longURL] = shortURL
		fmt.Fprintf(w, "Short URL: http://localhost:8080/%s", shortURL)
	} else {
		fmt.Fprintf(w, "Short URL already exists: http://localhost:8080/%s", longToShort[longURL])
	}
}

func redirectToLongURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["shortURL"]
	longURL, ok := shortToLong[shortURL]
	if !ok {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, longURL, http.StatusSeeOther)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", createShortURL).Methods("POST")
	r.HandleFunc("/{shortURL}", redirectToLongURL).Methods("GET")
	http.Handle("/", r)
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
