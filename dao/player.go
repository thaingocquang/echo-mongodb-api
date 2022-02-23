package dao

import (
	"context"
	"echo-mongodb-api/model"
	"echo-mongodb-api/module/database"

	"go.mongodb.org/mongo-driver/bson"
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

func PlayerFindByEmail(email string) (model.PlayerDetail, error) {
	var (
		userCol = database.PlayerCol()
		ctx     = context.Background()
	)

	var existedPlayer model.PlayerDetail

	// find player
	filter := bson.M{"email": email}
	err := userCol.FindOne(ctx, filter).Decode(&existedPlayer)

	if err != nil {
		return existedPlayer, err
	}
	return existedPlayer, nil
}

func PlayerProfileByEmail(email string) (model.PlayerProfile, error) {
	var (
		userCol = database.PlayerCol()
		ctx     = context.Background()
	)

	var profile model.PlayerProfile

	// find profile
	filter := bson.M{"email": email}
	err := userCol.FindOne(ctx, filter).Decode(&profile)

	if err != nil {
		return profile, err
	}
	return profile, nil
}
