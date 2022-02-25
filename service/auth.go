package service

import (
	"echo-mongodb-api/dao"
	"echo-mongodb-api/model"
	"echo-mongodb-api/util"
	"errors"
)

// Register ...
func Register(body model.PlayerCreateBody) (model.PlayerBSON, error) {
	var (
		playerBSON = body.ConvertToBSON()
	)

	// check email exist
	isEmailExisted := dao.PlayerCheckEmailExisted(body.Email)
	if isEmailExisted {
		err := errors.New("email already existed")
		return playerBSON, err
	}

	// hash player's password
	playerBSON.Password = util.HashPassword(playerBSON.Password)

	//Create user
	doc, err := dao.PlayerCreate(playerBSON)
	if err != nil {
		err = errors.New("can not create user")
		return doc, err
	}

	return doc, err
}

// Login ...
func Login(body model.LoginBody) (string, error) {
	// find player in db
	player, err := dao.PlayerFindByEmail(body.Email)

	// verify player password
	check := util.VerifyPassword(player.Password, body.Password)

	// if err
	if err != nil {
		err := errors.New("can not login")
		return "", err
	}

	// wrong password
	if !check {
		err := errors.New("password in correct")
		return "", err
	}

	data := map[string]interface{}{
		"id": player.ID,
	}

	// generate token
	return util.GenerateUserToken(data), nil
}
