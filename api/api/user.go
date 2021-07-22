package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/shaj13/go-guardian/auth"
	"leggett.dev/devmarks/api/app"
	myAuth "leggett.dev/devmarks/api/auth"
	"leggett.dev/devmarks/api/model"
)

// UserInput represents the input to the CreateUser function
type UserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserResponse represents the response written to the HTTP response header upon
// CreateUser's completion
type UserResponse struct {
	ID uint `json:"id"`
}

// CreateUser creates a new user based on the json data provided in the HTTP Request
func (a *API) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input UserInput

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

	user := &model.User{Email: input.Email}

	if err := a.App.CreateUser(user, input.Password); err != nil {
		if err, ok := err.(*app.ValidationError); ok {
			respondWithError(w, http.StatusUnprocessableEntity, err.Error())
		} else {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	err = respondWithJSON(w, http.StatusInternalServerError, &UserResponse{ID: user.ID})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

// GetUser Retrieves the authenticated user from the database
func (a *API) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := myAuth.GetUser(ctx)
	if user == nil {
		respondWithError(w, http.StatusUnauthorized, "no user signed in")
		return
	}

	err := respondWithJSON(w, http.StatusOK, user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (a *API) validateLogin(r *http.Request, userName, password string) (auth.Info, error) {
	user, err := a.App.GetUserByEmail(userName)

	if user == nil || err != nil {
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	if ok := user.CheckPassword(password); !ok {
		return nil, errors.Wrap(err, "invalid credentials")
	}
	return auth.NewDefaultUser(user.Email, strconv.Itoa(int(user.ID)), nil, nil), nil
}

type TokenResponse struct {
	Token string `json:"token"`
}

func (a *API) createToken(w http.ResponseWriter, r *http.Request) {
	var input UserInput

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

	user, err := a.validateLogin(r, input.Email, input.Password)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	bearerToken := uuid.New().String()
	a.App.AuthCache.Store(bearerToken, user, r)
	err = respondWithJSON(w, http.StatusOK, &TokenResponse{Token: bearerToken})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
