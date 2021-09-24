package main

import (
	"database/sql"
	"log"

	lpr "github.com/Zhangbokai614/Automated-Teller-Machine/pkgs/lpr/controller/gin"
	shibor "github.com/Zhangbokai614/Automated-Teller-Machine/pkgs/shibor/controller/gin"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()

	dbConn, err := sql.Open("mysql", "root:614525683@(192.168.0.22:3308)/bank")
	if err != nil {
		panic(err)
	}

	shiborCon := shibor.New(dbConn)
	shiborCon.RegisterRouter(router.Group("api/v1"))

	lprCon := lpr.New(dbConn)
	lprCon.RegisterRouter(router.Group("api/v1"))

	log.Fatal(router.Run(":8081"))
}
