package main

import "errors"

var (
	NotFoundPrefectureName = errors.New("not found prefecture name")
	InvalidLeagueKind      = errors.New("invalid league kind")
)
