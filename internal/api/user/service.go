package user

import (
	"context"
	"log"

	"github.com/jfraska/golang-app/infra/oauth"
	"github.com/jfraska/golang-app/infra/response"
	"github.com/jfraska/golang-app/internal/config"

	"github.com/jfraska/golang-app/infra/session"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	session.Store.Set(ctx, model.PublicID.String(), session.SessionStore{
		Name:  model.Name,
		Email: user.Email,
		Image: user.Image,
	})

	token, err = model.GenerateToken()
	return
}

func (s service) logout(ctx context.Context, ID string) {
	session.Store.Del(ctx, ID)
}

func (s service) session(ctx context.Context, ID string) (model User, err error) {
	auth, err := session.Store.Get(ctx, ID)
	if err != nil {
		return
	}

	model = User{
		Name:  auth.Name,
		Email: auth.Email,
		Image: auth.Image,
	}

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

	session.Store.Set(ctx, user.PublicID.String(), session.SessionStore{
		Name:  user.Name,
		Email: user.Email,
		Image: user.Image,
	})

	token, err = user.GenerateToken()

	return
}

func (r repository) createSlugIndex() (err error) {
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"public_id": 1},
		Options: options.Index().SetUnique(true),
	}

	_, err = r.collection.Indexes().CreateOne(context.TODO(), indexModel)
	return
}
