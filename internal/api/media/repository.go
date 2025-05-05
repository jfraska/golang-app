package media

import (
	"bytes"
	"context"

	"github.com/jfraska/golang-app/pkg/utils"
	"github.com/minio/minio-go/v7"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	collection *mongo.Collection
	storage    *minio.Client
}

func newRepository(db *mongo.Database, storage *minio.Client) Repository {
	return &repository{
		collection: db.Collection("media"),
		storage:    storage,
	}

}

func (r repository) CreateStorage(ctx context.Context, file []byte, model Media) (err error) {

	data := bytes.NewReader(file)

	_, err = r.storage.PutObject(ctx, model.Collection, model.InvitationID.Hex()+"/"+model.Filename, data, model.FileSize, minio.PutObjectOptions{
		ContentType: model.FileType,
	})

	if err != nil {
		return
	}

	return
}

func (r repository) CreateMedia(ctx context.Context, model Media) (err error) {

	_, err = r.collection.InsertOne(ctx, model)
	if err != nil {
		return
	}

	return
}

func (r repository) GetAllMediaByObjectName(ctx context.Context, model Media, pagination utils.Pagination) ([]Media, utils.Pagination, error) {
	var media []Media

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{{Key: "invitation_id", Value: model.InvitationID}}}},

		{{
			Key: "$facet", Value: bson.M{
				"metadata": bson.A{
					bson.M{
						"$group": bson.M{
							"_id":        nil,
							"total_size": bson.M{"$sum": "$file_size"},
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
				"data": func() bson.A {
					matchStage := bson.D{}
					if model.FileType != "" {
						matchStage = bson.D{{Key: "file_type", Value: model.FileType}}
					}
					return bson.A{
						bson.M{"$match": matchStage},
						bson.M{"$skip": (pagination.Page - 1) * pagination.Limit},
						bson.M{"$limit": pagination.Limit},
					}
				}(),
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
			Metadata bson.M  `bson:"metadata"`
			Data     []Media `bson:"data"`
		}

		err = cursor.Decode(&result)
		if err != nil {
			return nil, pagination, err
		}

		pagination = utils.NewPaginationFromModel(result.Metadata)

		media = append(media, result.Data...)
	}

	return media, pagination, nil
}

func (r repository) DeleteStorage(ctx context.Context, model Media) (err error) {

	return r.storage.RemoveObject(ctx, model.Collection, model.InvitationID.Hex()+"/"+model.Filename, minio.RemoveObjectOptions{})
}

func (r repository) DeleteMedia(ctx context.Context, model Media) (err error) {

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": model.ID})
	if err != nil {
		return
	}

	return
}

func (r repository) GetMediaByID(ctx context.Context, ID primitive.ObjectID) (model Media, err error) {

	if err = r.collection.FindOne(ctx, bson.M{"_id": ID}).Decode(&model); err != nil {
		return
	}

	return
}
