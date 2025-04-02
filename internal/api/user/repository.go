package user

import (
	"context"
	"encoding/json"

	"github.com/jfraska/golang-app/infra/oauth"
	"github.com/jfraska/golang-app/infra/response"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/oauth2"
)

type repository struct {
	collection *mongo.Collection
}

func newRepository(db *mongo.Database) Repository {
	repo := &repository{
		collection: db.Collection("users"),
	}
	repo.createSlugIndex()
	return repo
}

func (r repository) CreateUser(ctx context.Context, model User) (err error) {

	_, err = r.collection.InsertOne(ctx, model)

	return
}

func (r repository) UpdateUser(ctx context.Context, id primitive.ObjectID, model User) (err error) {

	selector := bson.M{"_id": id}
	changes := bson.M{"$set": model}

	result, err := r.collection.UpdateOne(ctx, selector, changes)
	if err != nil {
		return
	}

	if result.MatchedCount == 0 {
		err = response.ErrNotFound
		return
	}

	return
}

func (r repository) GetUserByEmail(ctx context.Context, email string) (model User, err error) {

	if err = r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&model); err != nil {
		if err == mongo.ErrNoDocuments {
			err = response.ErrNotFound
			return
		}
		return
	}

	return
}

func (r repository) FetchGoogleUserInfo(ctx context.Context, token *oauth2.Token) (model OauthUserResponse, err error) {

	client := oauth.GetGoogleOauthConfig().Client(ctx, token)

	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")

	if err != nil {
		return
	}

	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(&model); err != nil {
		return
	}

	return
}
