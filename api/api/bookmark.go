package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"leggett.dev/devmarks/api/auth"
	"leggett.dev/devmarks/api/model"
)

// GetBookmarks returns the bookmarks corresponding to the currently authenticated user in json form
func (a *API) GetBookmarks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.GetUser(ctx)
	if user == nil {
		respondWithError(w, http.StatusUnauthorized, "no user signed in")
		return
	}
	bookmarks, err := a.App.Database.GetBookmarksByUserID(user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = respondWithJSON(w, http.StatusOK, bookmarks)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

// CreateBookmarkInput represents the input to the CreateBookmark function
type CreateBookmarkInput struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Color string `json:"color"`
}

// CreateBookmarkResponse represents the response that will be sent upon completion of
// the CreateBookmark function
type CreateBookmarkResponse struct {
	ID uint `json:"id"`
}

// CreateBookmark creates a new bookmark owned by the currently authenticated user based
// on json from the HTTP Request
func (a *API) CreateBookmark(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.GetUser(ctx)
	if user == nil {
		respondWithError(w, http.StatusUnauthorized, "no user signed in")
		return
	}
	var input CreateBookmarkInput

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := json.Unmarshal(body, &input); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if input.URL == "" {
		respondWithError(w, http.StatusBadRequest, "url is required")
		return
	}
	if input.Name == "" {
		respondWithError(w, http.StatusBadRequest, "name is required")
		return
	}
	bookmark := &model.Bookmark{Name: input.Name, URL: input.URL, Color: &input.Color, OwnerID: user.ID}

	if err := a.App.Database.CreateBookmark(bookmark); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = respondWithJSON(w, http.StatusOK, &CreateBookmarkResponse{ID: bookmark.ID})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

// GetBookmarkByID writes the json representation of a bookmark to the HTTP Response Header,
// if the currently authenticated user has access to it.
func (a *API) GetBookmarkByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.GetUser(ctx)
	if user == nil {
		respondWithError(w, http.StatusUnauthorized, "no user signed in")
		return
	}
	id := getIDFromRequest(r)
	bookmark, err := a.App.Database.GetBookmarkByID(id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if user.ID != bookmark.OwnerID {
		respondWithError(w, http.StatusForbidden, "permission denied")
		return
	}

	err = respondWithJSON(w, http.StatusOK, bookmark)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

// UpdateBookmarkInput represents the input to the UpdateBookmark function
type UpdateBookmarkInput struct {
	Name  *string `json:"name"`
	URL   *string `json:"url"`
	Color *string `json:"color"`
}

// UpdateBookmarkByID updates the bookmark whose ID is specified in the HTTP request if it is owned
// by the currently authenticated user.
func (a *API) UpdateBookmarkByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.GetUser(ctx)
	if user != nil {
		respondWithError(w, http.StatusUnauthorized, "no user signed in")
		return
	}
	id := getIDFromRequest(r)

	var input UpdateBookmarkInput

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := json.Unmarshal(body, &input); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	existingBookmark, err := a.App.Database.GetBookmarkByID(id)
	if err != nil || existingBookmark == nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if user.ID != existingBookmark.OwnerID {
		respondWithError(w, http.StatusForbidden, "permission denied")
		return
	}

	if input.Name != nil {
		existingBookmark.Name = *input.Name
	}
	if input.URL != nil {
		existingBookmark.URL = *input.URL
	}
	if input.Color != nil {
		existingBookmark.Color = input.Color
	}

	err = a.App.Database.UpdateBookmark(existingBookmark)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = respondWithJSON(w, http.StatusOK, existingBookmark)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

// DeleteBookmarkByID deletes the bookmark whose ID is specified in the HTTP request if it is
// owned by the currently authenticated user
func (a *API) DeleteBookmarkByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.GetUser(ctx)
	if user == nil {
		respondWithError(w, http.StatusUnauthorized, "no user signed in")
		return
	}
	id := getIDFromRequest(r)
	err := a.App.Database.DeleteBookmarkByID(id)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = respondWithJSON(w, http.StatusNoContent, "")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func getIDFromRequest(r *http.Request) uint {
	vars := mux.Vars(r)
	id := vars["id"]

	intID, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return 0
	}

	return uint(intID)
}

func getBIDFromRequest(r *http.Request) uint {
	vars := mux.Vars(r)
	id := vars["bid"]

	intID, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return 0
	}

	return uint(intID)
}
