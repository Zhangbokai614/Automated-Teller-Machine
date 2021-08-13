package controller

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/Zhangbokai614/Automated-Teller-Machine/pkgs/shibor/model/mysql"
	"github.com/gin-gonic/gin"
)

type ShiborController struct {
	db *sql.DB
}

func New(db *sql.DB) *ShiborController {
	return &ShiborController{
		db: db,
	}
}

func (b *ShiborController) RegisterRouter(r gin.IRouter) {
	if r == nil {
		log.Fatal("[InitRouter]: server is nil")
	}

	r.POST("/update/shibor", b.update)
}

func (b *ShiborController) update(c *gin.Context) {
	var (
		req struct {
			InfoDate   time.Time `json:"date,omitempty"`
			OverNight  float32   `json:"O/N,omitempty"`
			OneWeek    float32   `json:"1W,omitempty"`
			TwoWeek    float32   `json:"2W,omitempty"`
			OneMonth   float32   `json:"1M,omitempty"`
			ThreeMonth float32   `json:"3M,omitempty"`
			SixMonth   float32   `json:"6M,omitempty"`
			NineMonth  float32   `json:"9M,omitempty"`
			OneYear    float32   `json:"1Y,omitempty"`
		}
	)

	if err := c.ShouldBind(&req); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	if err := mysql.InsertShibor(
		b.db, req.InfoDate, req.OverNight, req.OneWeek, req.TwoWeek,
		req.OneMonth, req.ThreeMonth, req.SixMonth, req.NineMonth, req.OneYear,
	); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
