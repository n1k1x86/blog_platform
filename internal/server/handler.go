package server

import (
	"blog-api/internal/database"
	"net/http"

	"github.com/gorilla/mux"
)

func GetArticles(repo *database.Repo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func GetArticleByID(repo *database.Repo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = mux.Vars(r)
	}
}

func DeleteArticleByID(repo *database.Repo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = mux.Vars(r)
	}
}

func CreateNewArticle(repo *database.Repo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func UpdateArticleByID(repo *database.Repo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = mux.Vars(r)
	}
}

func BuildHandlers(r *mux.Router, repo *database.Repo) {
	r.HandleFunc("/articles", GetArticles(repo)).Methods("GET")
	r.HandleFunc("/articles/{id}", GetArticleByID(repo)).Methods("POST")
	r.HandleFunc("/articles/{id}", DeleteArticleByID(repo)).Methods("DELETE")
	r.HandleFunc("/articles", CreateNewArticle(repo)).Methods("POST")
	r.HandleFunc("/articles/{id}", UpdateArticleByID(repo)).Methods("PUT")
}
