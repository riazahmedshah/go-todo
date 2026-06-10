package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/riazahmedshah/todo/models"
	"github.com/riazahmedshah/todo/stores"
	"github.com/riazahmedshah/todo/utils"
)

func GetAllTodosHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(models.UserIDKey).(int64)
	todos, err := stores.GetAllTodos(r.Context(), userId)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, map[string]any{"error": err})
		return
	}
	utils.WriteJson(w, http.StatusOK, todos)
}

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(models.UserIDKey).(int64)
	var t models.Todo
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		utils.WriteJson(w, http.StatusBadRequest, map[string]string{"error": "invalid body"})
		return
	}
	t.UserID = userId

	if err := stores.CreateTodo(r.Context(), t); err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJson(w, http.StatusOK, t)
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var body models.UpdateTodoInput
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.WriteJson(w, http.StatusBadRequest, map[string]string{"error": "invalid body"})
		return
	}

	if err := stores.UpdateTodo(r.Context(), id, body); err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{"message": "Todo updated"})
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := stores.DeleteTodo(r.Context(), id); err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{"message": "Todo deleted"})
}

func GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	todo, err := stores.GetTodo(r.Context(), id)

	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, todo)
}
