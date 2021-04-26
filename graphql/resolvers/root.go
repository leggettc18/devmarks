package resolvers

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"leggett.dev/devmarks/api/app"
	"leggett.dev/devmarks/api/auth"
	"leggett.dev/devmarks/api/model"
)

type RootResolver struct {
	App *app.App
}

func NewRoot(app *app.App) (*RootResolver, error) {
	r := &RootResolver{App: app}
	return r, nil
}

func(r RootResolver) Bookmarks(ctx context.Context) (*[]*BookmarkResolver, error){
	user, err := auth.AuthenticateToken(ctx, *r.App)
	if err != nil {
		return nil, err
	}
	
	bookmarks, err := r.App.Database.GetBookmarksByUserID(user.ID)

	if err != nil {
		return nil, err
	}

	var resolvers []*BookmarkResolver
	for _, bookmark := range bookmarks {
		resolvers = append(resolvers, &BookmarkResolver{*bookmark})
	}
	return &resolvers, nil
}

type NewBookmarkArgs struct {
	Name string
	Url string
	Color *string
}

func (r RootResolver) NewBookmark(ctx context.Context, args NewBookmarkArgs) (*BookmarkResolver, error) {
	user, err := auth.AuthenticateToken(ctx, *r.App)
	if err != nil {
		return nil, err
	}

	newBookmark := model.Bookmark {
		Name: args.Name,
		URL: args.Url,
		Color: *args.Color,
		OwnerID: user.ID,
	}

	if err := r.App.Database.CreateBookmark(&newBookmark); err != nil {
		return nil, err
	}

	bookmarkResolver := &BookmarkResolver{newBookmark}

	return bookmarkResolver, nil
}

type DeleteBookmarkArgs struct {
	ID int32
}

func (r RootResolver) DeleteBookmark(ctx context.Context, args DeleteBookmarkArgs) (bool, error) {
	user, err := auth.AuthenticateToken(ctx, *r.App)
	if err != nil {
		return false, err
	}
	bookmark, err := r.App.Database.GetBookmarkByID(uint(args.ID))
	if err != nil {
		return false, err
	}
	if user.ID == bookmark.OwnerID {
		err = r.App.Database.DeleteBookmarkByID(uint(args.ID))
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, DeleteError{Field: "delete", Message: "Permission denied."}
}

type RegisterArgs struct {
	Email string
	Password string
}

func (r *RootResolver) Register(args RegisterArgs) (*AuthResolver, error) {
	passwordHash, errHash := bcrypt.GenerateFromPassword(
		[]byte(args.Password),
		bcrypt.DefaultCost,
	)
	if errHash != nil {
		return nil, errHash
	}

	newUser := model.User{
		Email: args.Email,
		HashedPassword: passwordHash,
	}

	if err := r.App.Database.CreateUser(&newUser); err != nil {
		return nil, err
	}

	token, errToken := GenerateToken(&newUser, r.App.Config.SecretKey)
	if errToken != nil {
		return nil, errToken
	}

	payload := AuthPayload{
		Token: &token,
		User: &newUser,
	}

	return &AuthResolver{payload}, nil
}

type DeleteError struct {
	Field string `json:"field"`
	Message string `json:"message"`
}

func (e DeleteError) Error() string {
	return fmt.Sprintf("error [%s]: %s", e.Field, e.Message)
}

func (e DeleteError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"field": e.Field,
		"message": e.Message,
	}
}

type LoginError struct {
	Field string `json:"field"`
	Message string `json:"message"`
}

func (e LoginError) Error() string {
	return fmt.Sprintf("error [%s]: %s", e.Field, e.Message)
}

func (e LoginError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"field": e.Field,
		"message": e.Message,
	}
}
type LoginArgs struct {
	Email	 string
	Password string
}

func (r *RootResolver) Login(ctx context.Context, args LoginArgs) (*AuthResolver, error) {
	user, errUser := r.App.Database.GetUserByEmail(args.Email)
	if errUser != nil {
		return nil, LoginError{Field: "email", Message: errUser.Error()}
	}

	if correctPassword := model.ComparePasswordHash(user.HashedPassword, []byte(args.Password)); !correctPassword {
		return nil, LoginError{Field: "password", Message: "Incorrect Password"}
	}

	token, errToken := GenerateToken(user, r.App.Config.SecretKey)
	if errToken != nil {
		return nil, errToken
	}
	payload := AuthPayload{
		Token: &token,
		User: user,
	}
	return &AuthResolver{payload}, nil
}