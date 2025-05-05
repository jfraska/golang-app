package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

func ConvertObjectID(ID string) (newID primitive.ObjectID, err error) {
	newID, err = primitive.ObjectIDFromHex(ID)

	if err != nil {
		return
	}
	return
}
