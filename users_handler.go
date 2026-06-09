package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "invalid body"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 5)
	if err != nil {
		writeJson(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		return
	}

	u.Password = string(hash)

	if err := createUser(r.Context(), u); err != nil {
		writeJson(w, http.StatusInternalServerError, err)
		return
	}

	writeJson(w, http.StatusOK, map[string]string{"message": "registered successfully"})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var input Login
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "invalid body"})
		return
	}

	user, err := findUserByEmail(r.Context(), input.Email)
	if err != nil {
		writeJson(w, http.StatusNotFound, map[string]string{"error": "user not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "email/password does not match"})
		return
	}

	token, err := generateToken(user.ID)

	if err != nil {
		writeJson(w, http.StatusInternalServerError, map[string]string{"error": "something went wrong"})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  token,
		Value: token,
		Path:  "/",
	})

	writeJson(w, http.StatusOK, map[string]string{"message": "Logged in sucesssfully"})
}
