package handler

import "microservices/pkg"

type ProductHandler struct {
	DB *pkg.MysqlRepository
}
