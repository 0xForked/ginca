package domain

import "errors"

// error text
var (
	//ServiceUnavailable = errors.New("service currently unavailable")
	//MySQLUnavailable = errors.New("mysql currently unavailable")
	RedisUnavailable =  errors.New("redis currently unavailable")
	RouteNotFound = errors.New("route not found")
)

// success text
var (
	ServiceAvailable = "service is running well"
	MySQLAvaiable = "mysql is running well"
	RedisAvailable = "redis is running well"
)