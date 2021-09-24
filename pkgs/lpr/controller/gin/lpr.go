package controller

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/Zhangbokai614/Automated-Teller-Machine/pkgs/lpr/model/mysql"
	"github.com/gin-gonic/gin"
)

type lprController struct {
	db *sql.DB
}

func New(db *sql.DB) *lprController {
	return &lprController{
		db: db,
	}
}

func (b *lprController) RegisterRouter(r gin.IRouter) {
	if r == nil {
		log.Fatal("[InitRouter]: server is nil")
	}

	r.POST("/post/lpr", b.post)
	r.GET("/get/lpr", b.get)
}

func (b *lprController) post(c *gin.Context) {
	var (
		req struct {
			InfoDate time.Time `json:"date,omitempty"`
			OneYear  float32   `json:"1Y,omitempty"`
			FiveYear float32   `json:"5Y,omitempty"`
		}
	)

	if err := c.ShouldBind(&req); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	if err := mysql.InsertLpr(
		b.db, req.InfoDate, req.FiveYear, req.OneYear,
	); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

func (b *lprController) get(c *gin.Context) {
	lpr, err := mysql.QueryLpr(b.db)

	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "lpr": lpr})
}
