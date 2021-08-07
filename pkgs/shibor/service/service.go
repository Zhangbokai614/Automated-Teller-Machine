package service

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

type UpdateController struct {
	DB        *sql.DB
	tableName string
}

func New(DB *sql.DB, tableName string) *UpdateController {
	return &UpdateController{
		DB:        DB,
		tableName: tableName,
	}
}

func (b *UpdateController) RegisterRouter(r gin.IRouter) {
	if r == nil {
		log.Fatal("[InitRouter]: server is nil")
	}
}
