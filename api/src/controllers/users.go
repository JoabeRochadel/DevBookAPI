package controllers

import (
	"DevBookAPI/src/authentication"
	"DevBookAPI/src/data"
	"DevBookAPI/src/models"
	"DevBookAPI/src/repositories"
	"DevBookAPI/src/responses"
	"DevBookAPI/src/security"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Response struct {
	Id      uint64 `json:"id"`
	Message string `json:"message"`
}

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, err := data.Connecting()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	rep := repositories.NewRepositoryUsers(db)

	user, err := rep.FindAllUsers(nameOrNick)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

func FindOneUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["userId"], 10, 64)
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

	rep := repositories.NewRepositoryUsers(db)
	user, err := rep.FindOneUser(userId)

	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.Users
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	err = user.Prepare("registration")
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

	userRepository := repositories.NewRepositoryUsers(db)

	user.Id, err = userRepository.Create(user)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	user.Password = ""

	responses.JSON(w, http.StatusCreated, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	userId, err := strconv.ParseUint(param["userId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	userIdExtracted, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	if userIdExtracted != userId {
		responses.Err(w, http.StatusForbidden, errors.New("user is forbidden to update this ID"))
		return
	}

	bodyReq, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.Users
	err = json.Unmarshal(bodyReq, &user)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	err = user.Prepare("edit")
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

	userRepository := repositories.NewRepositoryUsers(db)
	err = userRepository.UpdateUser(userId, user)

	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	userId, err := strconv.ParseUint(param["userId"], 10, 60)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	userIdExtracted, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}
	if userIdExtracted != userId {
		responses.Err(w, http.StatusForbidden, errors.New("user is forbidden to delete this ID"))
		return
	}

	db, err := data.Connecting()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	rep := repositories.NewRepositoryUsers(db)
	err = rep.DeleteUser(userId)

	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if userId == followerId {
		responses.Err(w, http.StatusBadRequest, errors.New("it's impossible to follow yourself"))
		return
	}

	db, err := data.Connecting()
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}
	defer db.Close()

	rep := repositories.NewRepositoryUsers(db)

	err = rep.Follow(userId, followerId)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}
	if userId == followerId {
		responses.Err(w, http.StatusBadRequest, errors.New("it's impossible to unfollow yourself"))
		return
	}

	db, err := data.Connecting()
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	defer db.Close()

	rep := repositories.NewRepositoryUsers(db)

	err = rep.Unfollow(userId, followerId)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

func FindFollowers(w http.ResponseWriter, r *http.Request) {
	userId, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	db, err := data.Connecting()
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	defer db.Close()

	rep := repositories.NewRepositoryUsers(db)

	users, err := rep.FindFollowers(userId)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)

}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userIdToken, err := authentication.ExtractUserId(r)

	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if userIdToken != userId {
		responses.Err(w, http.StatusForbidden, errors.New("user is forbidden to update this password"))
	}

	bodyReq, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}
	defer r.Body.Close()

	var password models.Password
	err = json.Unmarshal(bodyReq, &password)
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

	rep := repositories.NewRepositoryUsers(db)
	passwordSavedDb, err := rep.FindById(userId)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	err = security.VerifyPassword(passwordSavedDb, password.CurrentPassword)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	newPasswordHash, err := security.Hash(password.NewPassword)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	err = rep.UpdatePassword(userId, newPasswordHash)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, nil)
}
