package client

import "errors"

var (
	ErrNotFoundPrefectureName = errors.New("not found prefecture name")
	ErrInvalidLeagueType      = errors.New("invalid league type")
)
