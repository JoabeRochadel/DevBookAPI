package controllers

import (
	"DevBookAPI/src/authentication"
	"DevBookAPI/src/data"
	"DevBookAPI/src/models"
	"DevBookAPI/src/repositories"
	"DevBookAPI/src/responses"
	"DevBookAPI/src/security"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	db, err := data.Connecting()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()
	repo := repositories.NewRepositoryUsers(db)
	userDb, err := repo.FindByEmail(user.Email)

	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	err = security.VerifyPassword(userDb.Password, user.Password)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	token, err := authentication.CreateToken(userDb.Id)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	fmt.Println(token)

	_, err = w.Write([]byte(token))
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
}
