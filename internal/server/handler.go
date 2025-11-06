package server

import (
	"blog-api/internal/database"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type GetArticlesRequest struct {
	Tag string `json:"tag"`
}

type GetArticlesResponse struct {
	Data []*database.ArticleJSON `json:"data"`
}

type ErrorResponse struct {
	Detail string `json:"detail"`
}

func BuildErrorResponse(err error) ([]byte, error) {
	resp := ErrorResponse{
		Detail: err.Error(),
	}
	data, err := json.Marshal(&resp)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func HandleError(w http.ResponseWriter, err error) {
	details, err := BuildErrorResponse(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = w.Write(details)
	if err != nil {
		log.Println(err)
		return
	}
}

func GetArticles(repo *database.Repo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			HandleError(w, err)
			return
		}
		var reqBody GetArticlesRequest
		err = json.Unmarshal(data, &reqBody)
		if err != nil {
			HandleError(w, err)
			return
		}
		articles, err := repo.GetArticlesList(reqBody.Tag)
		if err != nil {
			HandleError(w, err)
			return
		}
		data, err = json.Marshal(&GetArticlesResponse{Data: articles})
		if err != nil {
			HandleError(w, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}

func GetArticleByID(repo *database.Repo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		article, err := repo.GetArticleByID(params["id"])
		if err != nil {
			HandleError(w, err)
			return
		}
		data, err := json.Marshal(&article)
		if err != nil {
			HandleError(w, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}

func DeleteArticleByID(repo *database.Repo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		err := repo.DeleteArticleByID(params["id"])
		if err != nil {
			HandleError(w, err)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

type CreateNewArticleResponse struct {
	ID string `json:"id"`
}

func CreateNewArticle(repo *database.Repo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			HandleError(w, err)
			return
		}
		var reqBody database.NewArticle
		err = json.Unmarshal(data, &reqBody)
		if err != nil {
			HandleError(w, err)
			return
		}
		id, err := repo.CreateNewArticle(&reqBody)
		if err != nil {
			HandleError(w, err)
			return
		}
		resp := CreateNewArticleResponse{ID: id}
		data, err = json.Marshal(resp)
		if err != nil {
			HandleError(w, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}

func UpdateArticleByID(repo *database.Repo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		data, err := io.ReadAll(r.Body)
		if err != nil {
			HandleError(w, err)
			return
		}
		var reqBody database.ArticleJSON
		err = json.Unmarshal(data, &reqBody)
		if err != nil {
			HandleError(w, err)
			return
		}
		err = repo.UpdateArticleByID(params["id"], &reqBody)
		if err != nil {
			HandleError(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func BuildHandlers(r *mux.Router, repo *database.Repo) {
	r.HandleFunc("/articles", GetArticles(repo)).Methods("GET")
	r.HandleFunc("/articles/{id}", GetArticleByID(repo)).Methods("GET")
	r.HandleFunc("/articles/{id}", DeleteArticleByID(repo)).Methods("DELETE")
	r.HandleFunc("/articles", CreateNewArticle(repo)).Methods("POST")
	r.HandleFunc("/articles/{id}", UpdateArticleByID(repo)).Methods("PUT")
}
