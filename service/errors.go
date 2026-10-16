package service

import "errors"

const noExcludeId uint = 0

var (
	ErrLoginTaken   = errors.New("login is already exists")
	ErrEmailTaken   = errors.New("email is already exists")
	ErrUserNotFound = errors.New("user not found")

	ErrInvalidLogin    = errors.New("invalid login")
	ErrInvalidEmail    = errors.New("invalid email")
	ErrInvalidName     = errors.New("invalid name")
	ErrInvalidAge      = errors.New("invalid age")
	ErrInvalidPassword = errors.New("invalid age")
)

var (
	ErrClipNotFound = errors.New("clip not found")
	ErrClipExists   = errors.New("clip is already exists")

	ErrInvalidTitle       = errors.New("invalid title")
	ErrInvalidArtist      = errors.New("invalid artist")
	ErrInvalidDescription = errors.New("invalid description")
	ErrInvalidLink        = errors.New("invalid link")
	ErrInvalidTags        = errors.New("invalid tags")
	ErrInvalidMoods       = errors.New("invalid moods")
)
