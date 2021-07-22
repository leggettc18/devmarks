package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"leggett.dev/devmarks/api/auth"
	"leggett.dev/devmarks/api/model"
)

func (a *API) GetFolders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.GetUser(ctx)
	if user == nil {
		respondWithError(w, http.StatusUnauthorized, "no user signed in")
		return
	}
	folders, err := a.App.Database.GetFoldersByUserID(user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = respondWithJSON(w, http.StatusOK, folders); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (a *API) GetFolderByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.GetUser(ctx)
	id := getIDFromRequest(r)
	if user == nil {
		respondWithError(w, http.StatusInternalServerError, "no user signed in")
		return
	}
	folder, err := a.App.Database.GetFolderByID(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	if user.ID != folder.OwnerID {
		respondWithError(w, http.StatusForbidden, "permission denied")
		return
	}
	if err = respondWithJSON(w, http.StatusOK, folder); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

type CreateFolderInput struct {
	Name string `json:"name"`
	Color string `json:"color"`
	ParentID *uint `json:"parent_id"`
}

func (a *API) CreateFolder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.GetUser(ctx)
	if user == nil {
		respondWithError(w, http.StatusUnauthorized, "no user signed in")
		return
	}
	var input CreateFolderInput
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

	if input.Name == "" {
		respondWithError(w, http.StatusUnprocessableEntity, "name is required")
		return
	}
	folder := &model.Folder{Name: input.Name, Color: input.Color, ParentID:input.ParentID, OwnerID: user.ID}

	if err := a.App.Database.CreateFolder(folder); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	err = respondWithJSON(w, http.StatusCreated, folder)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (a *API) AddBookmarkToFolder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.GetUser(ctx)
	if user == nil {
		respondWithError(w, http.StatusUnauthorized, "no user is signed in")
		return
	}
	folder_id := getIDFromRequest(r)
	bookmark_id := getBIDFromRequest(r)

	if err := a.App.Database.AddBookmarkToFolder(ctx, bookmark_id, folder_id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	folder, err := a.App.Database.GetFolderByID(folder_id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = respondWithJSON(w, http.StatusOK, folder); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}