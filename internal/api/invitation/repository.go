package invitation

import (
	"context"

	"github.com/jfraska/golang-app/infra/response"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repository struct {
	collection *mongo.Collection
}

func newRepository(db *mongo.Database) Repository {
	repo := &repository{
		collection: db.Collection("invitations"),
	}
	repo.createSubdomainIndex()
	return repo
}

func (r repository) CreateInvitation(ctx context.Context, model Invitation) (err error) {

	_, err = r.collection.InsertOne(ctx, model)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			err = response.ErrorBadRequest
			return
		}
		return
	}

	return
}

func (r repository) createSubdomainIndex() (err error) {
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"subdomain": 1},
		Options: options.Index().SetUnique(true),
	}

	_, err = r.collection.Indexes().CreateOne(context.TODO(), indexModel)
	return
}
