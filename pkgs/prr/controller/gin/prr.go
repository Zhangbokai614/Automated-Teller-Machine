package controller

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/Zhangbokai614/Automated-Teller-Machine/pkgs/prr/model/mysql"
	"github.com/gin-gonic/gin"
)

type prrController struct {
	db *sql.DB
}

func New(db *sql.DB) *prrController {
	return &prrController{
		db: db,
	}
}

func (b *prrController) RegisterRouter(r gin.IRouter) {
	if r == nil {
		log.Fatal("[InitRouter]: server is nil")
	}

	r.POST("/post/prr", b.post)
	r.GET("/get/prr", b.get)
}

func (b *prrController) post(c *gin.Context) {
	var (
		req struct {
			InfoDate      time.Time `json:"date,omitempty"`
			One_day       float32   `json:"DR001"`
			Seven_day     float32   `json:"DR007"`
			Fourteen_day  float32   `json:"DR014"`
			Twentyone_day float32   `json:"DR021"`
			One_month     float32   `json:"DR1M"`
			Two_month     float32   `json:"DR2M"`
		}
	)

	if err := c.ShouldBind(&req); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	if err := mysql.InsertPrr(
		b.db, req.InfoDate, req.One_day, req.Seven_day, req.Fourteen_day,
		req.Twentyone_day, req.One_month, req.Two_month,
	); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

func (b *prrController) get(c *gin.Context) {
	prr, err := mysql.QueryPrr(b.db)

	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "prr": prr})
}
