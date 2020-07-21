package domain

import "errors"

// error text
var (
	RedisUnavailable =  errors.New("redis currently unavailable")
)

// success text
var (
	RedisAvailable = "redis is running well"
)