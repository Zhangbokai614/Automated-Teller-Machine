package main

import (
	"database/sql"
	"log"

	shibor "github.com/Zhangbokai614/Automated-Teller-Machine/pkgs/shibor/controller/gin"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()

	dbConn, err := sql.Open("mysql", "root:123456@(192.168.0.251:6666)/project")
	if err != nil {
		panic(err)
	}

	shiborCon := shibor.New(dbConn)
	shiborCon.RegisterRouter(router.Group("api/v1"))

	log.Fatal(router.Run(":8080"))
}
