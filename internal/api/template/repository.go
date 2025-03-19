package template

import (
	"context"
	"golang-app/infra/response"
	pkg "golang-app/pkg/utils"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repository struct {
	collection *mongo.Collection
}

func newRepository(db *mongo.Database) Repository {
	repo := &repository{
		collection: db.Collection("templates"),
	}
	repo.createSlugIndex()
	return repo
}

func (r repository) CreateTemplate(ctx context.Context, model Template) (err error) {

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

func (r repository) GetAllTemplates(ctx context.Context, model pkg.Pagination) ([]Template, pkg.Pagination, error) {
	var templates []Template

	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var row Template
		err := cursor.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		templates = append(templates, row)
	}

	return templates, model, nil
}

func (r repository) GetTemplateBySlug(ctx context.Context, slug string) (model Template, err error) {

	if err = r.collection.FindOne(ctx, bson.M{"slug": slug}).Decode(&model); err != nil {
		if err == mongo.ErrNoDocuments {
			err = response.ErrNotFound
			return
		}
		return
	}

	return
}

func (r repository) createSlugIndex() (err error) {
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"slug": 1},
		Options: options.Index().SetUnique(true),
	}

	_, err = r.collection.Indexes().CreateOne(context.TODO(), indexModel)
	return
}
