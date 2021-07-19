package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/shaj13/go-guardian/auth/strategies/token"
	"leggett.dev/devmarks/api/app"
	"leggett.dev/devmarks/api/log"
	"leggett.dev/devmarks/api/model"
)

type AuthService interface {
	AuthMiddleware (http.Handler) http.Handler
}

type authSvc struct {
	App app.App
	ExemptPaths *[]string
}

func NewAuth(exemptPaths *[]string, app app.App) AuthService{
	return &authSvc{
		App: app,
		ExemptPaths: exemptPaths,
	}
}

func contains(s string, array []string) bool {
	for _, b := range array {
		if b == s {
			return true
		}
	}
	return false
}

func (a *authSvc) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := log.GetLogger(ctx)
		if !(contains(r.URL.Path, *a.ExemptPaths)) {
			tokenStrategy := a.App.Authenticator.Strategy(token.CachedStrategyKey)
			userInfo, err := tokenStrategy.Authenticate(r.Context(), r)
			if err != nil {
				if logger != nil {
					logger.WithError(err).Error("unable to get user")
				}
				http.Error(w, "invalid credentials", http.StatusForbidden)
				return
			}
			user, err := a.App.GetUserByEmail(userInfo.UserName())

			if user == nil || err != nil {
				if err != nil {
					if logger != nil {
						logger.WithError(err).Error("unable to get user")
					}
				}
				http.Error(w, "invalid credentials", http.StatusForbidden)
				return
			}

			setUserInCtx(&ctx, user)
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type contextKey struct { key string }
var userKey = &contextKey{"user"}

func setUserInCtx(ctx *context.Context, user *model.User) {
	*ctx = context.WithValue(*ctx, userKey, user)
}

func GetUser(ctx context.Context) *model.User {
	return ctx.Value(userKey).(*model.User)
}

func AuthenticateToken(ctx context.Context, app app.App) (*model.User, error) {
	var token, ok = ctx.Value("token").(string)
	if !ok {
		return nil, errors.New("bookmarks: no authenticated user in context")
	}

	user, err := getUserFromToken(token, app)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func getUserFromToken(tokenString string, app app.App) (*model.User, error) {
	// decode token with the secret if was encoded with
	tokenObj, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return app.Config.SecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	// get user ID from the map we encoded in the token
	userID, ok := tokenObj.Claims.(jwt.MapClaims)["ID"].(float64)
	if !ok {
		return nil, errors.New("GetUserIDFromToken error: type conversion in claims")
	}

	user, err := app.Database.GetUserById(uint(userID))

	if err != nil {
		return nil, errors.New("No user with ID " + fmt.Sprint(userID))
	}

	return user, nil
}