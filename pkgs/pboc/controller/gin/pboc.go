package controller

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/Zhangbokai614/Automated-Teller-Machine/pkgs/pboc/model/mysql"
	"github.com/gin-gonic/gin"
)

type pbocController struct {
	db *sql.DB
}

func New(db *sql.DB) *pbocController {
	return &pbocController{
		db: db,
	}
}

func (b *pbocController) RegisterRouter(r gin.IRouter) {
	if r == nil {
		log.Fatal("[InitRouter]: server is nil")
	}

	r.POST("/post/pboc", b.post)
	r.GET("/get/pboc", b.get)
}

func (b *pbocController) post(c *gin.Context) {
	var (
		req struct {
			InfoDate       time.Time `json:"date,omitempty"`
			Period         int32     `json:"period,omitempty"`
			Deal_amount    int32     `json:"dealAmount,omitempty"`
			Rate           float32   `json:"rate,omitempty"`
			Trading_method string    `json:"tradingMethod,omitempty"`
		}
	)

	if err := c.ShouldBind(&req); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	if err := mysql.InsertPboc(
		b.db, req.InfoDate, req.Period, req.Deal_amount, req.Rate, req.Trading_method,
	); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

func (b *pbocController) get(c *gin.Context) {
	pboc, err := mysql.QueryPboc(b.db)

	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "pboc": pboc})
}
