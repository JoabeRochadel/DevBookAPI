package controllers

import (
	"DevBookAPI/src/data"
	"DevBookAPI/src/models"
	"DevBookAPI/src/repositories"
	"DevBookAPI/src/responses"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

func FindOnePost(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	postId, err := strconv.ParseUint(param["postId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := data.Connecting()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	rep := repositories.NewRepositoryPosts(db)
	post, err := rep.FindOnePost(postId)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
	}

	responses.JSON(w, http.StatusOK, post)

}

func FindAllPosts(w http.ResponseWriter, r *http.Request) {

}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	var posts models.Posts

	err = json.Unmarshal(body, &posts)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := data.Connecting()
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	defer db.Close()

	rep := repositories.NewRepositoryPosts(db)
	posts.Id, err = rep.Create(posts)

}

func UpdatePost(w http.ResponseWriter, r *http.Request) {

}

func DeletePost(w http.ResponseWriter, r *http.Request) {

}
