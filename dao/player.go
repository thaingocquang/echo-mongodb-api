package dao

import (
	"context"
	"echo-mongodb-api/model"
	"echo-mongodb-api/module/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PlayerCreate ...
func PlayerCreate(doc model.PlayerBSON) (model.PlayerBSON, error) {
	var (
		userCol = database.PlayerCol()
		ctx     = context.Background()
	)

	// Insert one
	_, err := userCol.InsertOne(ctx, doc)
	return doc, err
}

// PlayerFindByEmail ...
func PlayerFindByEmail(email string) (model.PlayerDetail, error) {
	var (
		userCol       = database.PlayerCol()
		ctx           = context.Background()
		existedPlayer model.PlayerDetail
	)

	// find player
	filter := bson.M{"email": email}
	err := userCol.FindOne(ctx, filter).Decode(&existedPlayer)

	// if err
	if err != nil {
		return existedPlayer, err
	}

	return existedPlayer, nil
}

// PlayerProfileFindByID ...
func PlayerProfileFindByID(ID string) (model.PlayerProfile, error) {
	var (
		userCol = database.PlayerCol()
		ctx     = context.Background()
		profile model.PlayerProfile
	)

	// objectID
	objID, _ := primitive.ObjectIDFromHex(ID)

	// find profile
	filter := bson.M{"_id": objID}
	err := userCol.FindOne(ctx, filter).Decode(&profile)

	// if err
	if err != nil {
		return profile, err
	}

	return profile, nil
}

// CheckEmailExisted ...
func PlayerCheckEmailExisted(email string) bool {
	var (
		userCol = database.PlayerCol()
		ctx     = context.Background()
	)

	count, err := userCol.CountDocuments(ctx, bson.M{"email": email})
	if err != nil || count <= 0 {
		return false
	}

	return true
}
