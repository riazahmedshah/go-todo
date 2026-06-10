package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/riazahmedshah/todo/models"
	"github.com/riazahmedshah/todo/stores"
	"github.com/riazahmedshah/todo/utils"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		utils.WriteJson(w, http.StatusBadRequest, map[string]string{"error": "invalid body"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 5)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		return
	}

	u.Password = string(hash)

	if err := stores.CreateUser(r.Context(), u); err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{"message": "registered successfully"})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var input models.Login
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteJson(w, http.StatusBadRequest, map[string]string{"error": "invalid body"})
		return
	}

	user, err := stores.FindUserByEmail(r.Context(), input.Email)
	if err != nil {
		utils.WriteJson(w, http.StatusNotFound, map[string]string{"error": "user not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		utils.WriteJson(w, http.StatusBadRequest, map[string]string{"error": "email/password does not match"})
		return
	}

	token, err := utils.GenerateToken(user.ID)

	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, map[string]string{"error": "something went wrong"})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: token,
		Path:  "/",
	})

	utils.WriteJson(w, http.StatusOK, map[string]string{"message": "Logged in sucesssfully"})
}
