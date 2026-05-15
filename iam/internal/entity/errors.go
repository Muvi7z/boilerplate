package entity

import "errors"

var ErrSessionNotFound = errors.New("session not found")
var ErrUserNotFound = errors.New("user not found")
var ErrGetUser = errors.New("error to get user")
var ErrCreateUser = errors.New("error to create user")
var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrRegisterUser = errors.New("error to register user")
