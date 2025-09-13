package main

import "errors"

var (
	ErrNotFoundPrefectureName = errors.New("not found prefecture name")
	ErrInvalidLeagueKind      = errors.New("invalid league kind")
)
