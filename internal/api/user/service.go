package user

import (
	"context"
	"golang-app/infra/oauth"
	"golang-app/infra/response"
	"golang-app/infra/session"
	"golang-app/internal/config"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/oauth2"
)

type Repository interface {
	GetUserByEmail(ctx context.Context, email string) (model User, err error)
	UpdateUser(ctx context.Context, id primitive.ObjectID, model User) (err error)
	CreateUser(ctx context.Context, model User) (err error)
	FetchGoogleUserInfo(ctx context.Context, token *oauth2.Token) (model OauthUserResponse, err error)
}

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) register(ctx context.Context, req RegisterRequestPayload) (err error) {
	user := NewUserFromRegisterRequest(req)

	if err = user.Validate(); err != nil {
		return
	}

	if err = user.EncryptPassword(int(config.Cfg.Encryption.Salt)); err != nil {
		return
	}

	model, err := s.repo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		if err != response.ErrNotFound {
			return
		}
	}

	if model.IsExists() {
		return response.ErrEmailAlreadyUsed
	}

	return s.repo.CreateUser(ctx, user)

}

func (s service) login(ctx context.Context, req LoginRequestPayload) (token string, err error) {
	user := NewUserFromLoginRequest(req)

	if err = user.ValidateEmail(); err != nil {
		return
	}
	if err = user.ValidatePassword(); err != nil {
		return
	}

	model, err := s.repo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		log.Println("error when try to GetUserByEmail with deail", err.Error())
		return
	}

	if err = user.VerifyPasswordFromPlain(model.Password); err != nil {
		err = response.ErrPasswordNotMatch
		return
	}

	session.Store.Set(ctx, model.PublicID.String(), session.Session{
		Name:   model.Name,
		UserID: model.PublicID,
	})

	token, err = model.GenerateToken(config.Cfg.Encryption.JWTSecret)
	return
}

func (s service) oauth(req OauthRequestPayload) (url string) {

	url = oauth.GetGoogleOauthConfig().AuthCodeURL(req.State)

	return
}

func (s service) oauthCallback(ctx context.Context, req OauthCallbackRequestPayload) (token string, err error) {
	var auth *oauth2.Token

	auth, err = oauth.GetGoogleOauthConfig().Exchange(context.Background(), req.Code)
	if err != nil {
		return
	}

	var ouser OauthUserResponse

	ouser, err = s.repo.FetchGoogleUserInfo(ctx, auth)
	if err != nil {
		return
	}

	model, err := s.repo.GetUserByEmail(ctx, ouser.Email)

	// add state value in user account
	ouser.State = req.State
	user := NewUserFromOauthUserResponse(ouser, auth)

	if err != nil {
		s.repo.CreateUser(ctx, user)
	} else {
		s.repo.UpdateUser(ctx, model.ID, user)
	}

	session.Store.Set(ctx, ouser.Id, session.Session{
		Name:   user.Name,
		UserID: user.PublicID,
	})

	token, err = user.GenerateToken(config.Cfg.Encryption.JWTSecret)

	return
}
