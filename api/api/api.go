package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/graph-gophers/graphql-go"
	"github.com/shaj13/go-guardian/auth"
	"github.com/shaj13/go-guardian/auth/strategies/token"
	"github.com/shaj13/go-guardian/store"

	"leggett.dev/devmarks/api/app"
	myAuth "leggett.dev/devmarks/api/auth"
	"leggett.dev/devmarks/api/log"
)

var (
	opts = []graphql.SchemaOpt{graphql.UseStringDescriptions()}
)

type statusCodeRecorder struct {
	http.ResponseWriter
	http.Hijacker
	StatusCode int
}

func (r *statusCodeRecorder) WriteHeader(statusCode int) {
	r.StatusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

// API is an object representing our API's configuration, and includes a pointer
// to our App's App object
type API struct {
	App    *app.App
	Config *Config
}

// New returns a new API object from our App's App object
func New(a *app.App) (api *API, err error) {
	api = &API{App: a}
	api.Config, err = InitConfig()
	if err != nil {
		return nil, err
	}
	return api, nil
}

func (a *API) setupGoGuardian() {
	a.App.Authenticator = auth.New()
	a.App.AuthCache = store.NewFIFO(context.Background(), time.Minute*10)

	tokenStrategy := token.New(token.NoOpAuthenticate, a.App.AuthCache)

	a.App.Authenticator.EnableStrategy(token.CachedStrategyKey, tokenStrategy)
}

// used to set any options on the http traffic, i.e. response headers,
// max request size, etc.
func apiMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 100*1024*1024)
		next.ServeHTTP(w, r)
	})
}

// Init Initializes our API (routes, authentication setup, etc.)
func (a *API) Init(r *mux.Router) {
	// authentication
	a.setupGoGuardian()
	logger := log.NewLogger(a.Config.ProxyCount)
	r.Use(logger.LoggerMiddleware)
	authSvc := myAuth.NewAuth(&[]string{"/users/", "/auth/token/"}, *a.App, &logger)
	r.Use(authSvc.AuthMiddleware)
	r.Use(apiMiddleware)
	r.HandleFunc("/auth/token/", a.createToken).Methods("POST")

	// user methods
	r.HandleFunc("/users/", a.CreateUser).Methods("POST")
	r.HandleFunc("/me/", a.GetUser).Methods("GET")

	// bookmark methods
	bookmarksRouter := r.PathPrefix("/bookmarks").Subrouter()
	bookmarksRouter.HandleFunc("/", a.GetBookmarks).Methods("GET")
	bookmarksRouter.HandleFunc("/", a.CreateBookmark).Methods("POST")
	bookmarksRouter.HandleFunc("/{id:[0-9]+}/", a.GetBookmarkByID).Methods("GET")
	bookmarksRouter.HandleFunc("/{id:[0-9]+}/", a.UpdateBookmarkByID).Methods("PATCH")
	bookmarksRouter.HandleFunc("/{id:[0-9]+}/", a.DeleteBookmarkByID).Methods("DELETE")

	foldersRouter := r.PathPrefix("/folders").Subrouter()
	foldersRouter.HandleFunc("/", a.CreateFolder).Methods("POST")
}
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) error {
	response, err := json.Marshal(payload)

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
	return nil
}
