package service

import (
	"echo-mongodb-api/auth"
	"echo-mongodb-api/config"
	"echo-mongodb-api/dao"
	"echo-mongodb-api/model"
	"echo-mongodb-api/util"
	"errors"
)

// PlayerCreate ...
func Register(body model.PlayerCreateBody) (model.PlayerBSON, error) {
	var (
		player = body.ConvertToBSON()
	)

	isEmailExisted, err := dao.CheckEmailExisted(body.Email)
	if err != nil {
		return player, err
	}

	if isEmailExisted {
		err = errors.New("email already existed")
		return player, err
	}

	// hash player's password
	player.Password = util.HashPassword(player.Password)

	//Create user
	doc, err := dao.PlayerCreate(player)
	if err != nil {
		err = errors.New("can not create user")
		return doc, err
	}

	return doc, err
}

// Login ...
func Login(body model.LoginBody) (string, error) {

	player, err := dao.PlayerFindByEmail(body.Email)
	check, _ := util.VerifyPassword(player.Password, body.Password)
	if err != nil || check {
		err := errors.New("can not login")
		return "", err
	}

	envVars := config.GetEnv()
	jwtWrapper := auth.JwtWrapper{
		SecretKey:       envVars.JWT.SecretKey,
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	// geneate token
	signedToken, err := jwtWrapper.GenerateToken(body.Email)
	if err != nil {
		err := errors.New("error signing token")
		return "", err
	}

	return signedToken, nil
}
