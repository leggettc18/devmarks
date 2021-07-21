package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"leggett.dev/devmarks/api/auth"
	"leggett.dev/devmarks/api/model"
)

type CreateFolderInput struct {
	Name string `json:"name"`
	Color string `json:"color"`
	ParentID *uint `json:"parent_id"`
}

type CreateFolderResponse struct {
	ID uint `json:"id"`
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
		respondWithError(w, http.StatusBadRequest, "name is required")
		return
	}
	folder := &model.Folder{Name: input.Name, Color: input.Color, ParentID:input.ParentID, OwnerID: user.ID}

	if err := a.App.Database.CreateFolder(folder); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	err = respondWithJSON(w, http.StatusOK, &CreateFolderResponse{ID: folder.ID})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}