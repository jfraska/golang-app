package invitation

import (
	"context"

	"github.com/jfraska/golang-app/infra/response"
	"github.com/jfraska/golang-app/pkg/utils"
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

func (r repository) GetAllInvitations(ctx context.Context, UserID string, pagination utils.Pagination) ([]Invitation, utils.Pagination, error) {
	var invitations []Invitation

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{{Key: "user_ids", Value: UserID}}}},

		{{
			Key: "$facet", Value: bson.M{
				"metadata": bson.A{
					bson.M{
						"$group": bson.M{
							"_id":        nil,
							"total_rows": bson.M{"$sum": 1},
							"page":       bson.M{"$first": pagination.Page},
							"limit":      bson.M{"$first": pagination.Limit},
						},
					},
					bson.M{
						"$addFields": bson.M{
							"total_pages": bson.M{
								"$ceil": bson.M{
									"$divide": bson.A{"$total_rows", pagination.Limit},
								},
							},
						},
					},
				},
				"data": bson.A{
					bson.M{"$skip": (pagination.Page - 1) * pagination.Limit},
					bson.M{"$limit": pagination.Limit},
				},
			},
		}},

		{{Key: "$unwind", Value: "$metadata"}},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, pagination, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {

		var result struct {
			Metadata bson.M       `bson:"metadata"`
			Data     []Invitation `bson:"data"`
		}

		err = cursor.Decode(&result)
		if err != nil {
			return nil, pagination, err
		}

		pagination = utils.NewPaginationFromModel(result.Metadata)

		invitations = append(invitations, result.Data...)
	}

	return invitations, pagination, nil
}

func (r repository) createSubdomainIndex() (err error) {
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"subdomain": 1},
		Options: options.Index().SetUnique(true),
	}

	_, err = r.collection.Indexes().CreateOne(context.TODO(), indexModel)
	return
}
