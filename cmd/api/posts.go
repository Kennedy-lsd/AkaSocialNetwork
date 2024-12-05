package api

import (
	"encoding/json"
	"net/http"
	"test/internal/data"
)

func (a *App) createPostHandler(w http.ResponseWriter, r *http.Request) {

	var post data.Post

	err := json.NewDecoder(r.Body).Decode(&post)

	if err != nil {
		http.Error(w, "Unable to parse the post data", http.StatusBadRequest)
		return
	}

	err = a.Data.Post.Create(r.Context(), &post)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, 200, post)
}

func (a *App) getPostsHandler(w http.ResponseWriter, r *http.Request) {

	posts, err := a.Data.Post.GetAll(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, 200, posts)
}

func (a *App) getPostHandler(w http.ResponseWriter, r *http.Request) {

	id, err := idParser(r, 3)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	post, err := a.Data.Post.GetOne(r.Context(), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, 200, post)
}

func (a *App) deletePostHandler(w http.ResponseWriter, r *http.Request) {

	id, err := idParser(r, 3)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.Data.Post.Delete(r.Context(), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	writeJSON(w, 200, map[string]string{"message": "Post successfully deleted"})

}

func (a *App) updatePostHandler(w http.ResponseWriter, r *http.Request) {
	id, err := idParser(r, 3)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var post data.Post

	err = json.NewDecoder(r.Body).Decode(&post)

	if err != nil {
		http.Error(w, "Unable to parse the user data", http.StatusBadRequest)
		return
	}

	err = a.Data.Post.Update(r.Context(), id, &post)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, 200, post)

}
