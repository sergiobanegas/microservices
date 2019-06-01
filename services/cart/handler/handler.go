package handler

import (
	"microservices/pkg"
)

type CartHandler struct {
	RedisRepository *pkg.RedisRepository
	TokenService    *pkg.TokenService
}
