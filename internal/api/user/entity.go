package user

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jfraska/golang-app/infra/response"
	"github.com/jfraska/golang-app/internal/config"
	pkg "github.com/jfraska/golang-app/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
)

type Role string

const (
	ROLE_Admin Role = "admin"
	ROLE_User  Role = "user"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id" json:"_id"`
	Name          string             `bson:"name" json:"name"`
	Email         string             `bson:"email" json:"email"`
	PublicID      uuid.UUID          `bson:"public_id" json:"public_id"`
	Password      string             `bson:"password" json:"password"`
	Role          Role               `bson:"role" json:"role"`
	EmailVerified time.Time          `bson:"email_verified,omitempty" json:"email_verified,omitempty"`
	Image         string             `bson:"image,omitempty" json:"image,omitempty"`
	CreatedAt     time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt     time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`

	Accounts []Account `bson:"accounts,omitempty" json:"accounts,omitempty"`
}

type Account struct {
	Type              string `bson:"type" json:"type"`
	Provider          string `bson:"provider" json:"provider"`
	ProviderAccountID string `bson:"provider_account_id" json:"provider_account_id"`
	RefreshToken      string `bson:"refresh_token,omitempty" json:"refresh_token,omitempty"`
	AccessToken       string `bson:"access_token,omitempty" json:"access_token,omitempty"`
	ExpiresAt         int64  `bson:"expires_at,omitempty" json:"expires_at,omitempty"`
	TokenType         string `bson:"token_type,omitempty" json:"token_type,omitempty"`
	Scope             string `bson:"scope,omitempty" json:"scope,omitempty"`
	IdToken           string `bson:"id_token,omitempty" json:"id_token,omitempty"`
	SessionState      string `bson:"session_state,omitempty" json:"session_state,omitempty"`
}

func NewUserFromRegisterRequest(req RegisterRequestPayload) User {
	return User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		PublicID: uuid.New(),
		Role:     ROLE_User,
	}
}

func NewUserFromLoginRequest(req LoginRequestPayload) User {
	return User{
		Email:    req.Email,
		Password: req.Password,
	}
}

func NewUserFromOauthUserResponse(res OauthUserResponse, token *oauth2.Token) User {
	user := User{
		Name:          res.Name,
		Email:         res.Email,
		EmailVerified: time.Now(),
		PublicID:      uuid.New(),
		Image:         res.Image,
		Role:          ROLE_User,
	}

	account := Account{
		Type:              "oauth",
		Provider:          "google",
		ProviderAccountID: res.Id,
		AccessToken:       token.AccessToken,
		RefreshToken:      token.RefreshToken,
		ExpiresAt:         token.Expiry.Unix(),
		TokenType:         token.TokenType,
		Scope:             token.Extra("scope").(string),
		IdToken:           token.Extra("id_token").(string),
		SessionState:      res.State,
	}

	user.Accounts = append(user.Accounts, account)

	return user
}

func (a User) Validate() (err error) {
	if err = a.ValidateEmail(); err != nil {
		return
	}
	if err = a.ValidatePassword(); err != nil {
		return
	}
	return
}

func (a User) ValidateEmail() (err error) {
	if a.Email == "" {
		return response.ErrEmailRequired
	}

	emails := strings.Split(a.Email, "@")
	if len(emails) != 2 {
		return response.ErrEmailInvalid
	}
	return
}

func (a User) ValidatePassword() (err error) {
	if a.Password == "" {
		return response.ErrPasswordRequired
	}

	if len(a.Password) < 6 {
		return response.ErrPasswordInvalidLength
	}
	return
}

func (a User) IsExists() bool {
	return !a.ID.IsZero()
}

func (a *User) EncryptPassword(salt int) (err error) {

	encryptedPass, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	a.Password = string(encryptedPass)
	return nil
}

func (a User) VerifyPasswordFromEncrypted(plain string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(plain))
}

func (a User) VerifyPasswordFromPlain(encrypted string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(a.Password))
}

func (a User) GenerateToken() (tokenString string, err error) {
	return pkg.GenerateToken(a.PublicID.String(), string(a.Role), config.Cfg.Encryption.JWTSecret)
}
