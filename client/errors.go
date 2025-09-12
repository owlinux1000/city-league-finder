package client

import "errors"

var (
	NotFoundPrefectureName = errors.New("not found prefecture name")
	InvalidLeagueType = errors.New("invalid league type")
)
