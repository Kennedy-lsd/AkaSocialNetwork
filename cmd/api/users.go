package api

import (
	"encoding/json"
	"net/http"
	"test/internal/data"
)

func (a *App) createUserHandler(w http.ResponseWriter, r *http.Request) {

	var user data.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Unable to parse the user data", http.StatusBadRequest)
		return
	}

	err = a.Data.User.Create(r.Context(), &user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, 200, user)
}

func (a *App) getUsersHandler(w http.ResponseWriter, r *http.Request) {

	users, err := a.Data.User.GetAll(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	writeJSON(w, 200, users)
}

func (a *App) getUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := idParser(r, 3)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := a.Data.User.GetOne(r.Context(), id)

	if err != nil {
		http.Error(w, "Not found", http.StatusInternalServerError)
		return
	}

	writeJSON(w, 200, user)
}

func (a *App) deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := idParser(r, 3) // because /api/users/{id}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.Data.User.Delete(r.Context(), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	writeJSON(w, 200, map[string]string{"message": "User successfully deleted"})
}
