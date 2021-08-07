package main

import (
	"database/sql"

	"github.com/Zhangbokai614/Automated-Teller-Machine/pkgs/shibor/controller/update"	"github.com/gin-gonic/gin"
)


func main() {
	update.OnUpdate()

	router := gin.Default()

	dbConn := sql.Open("mysql", "root:123456@(192.168.0.251:6666)/project")
	if err != nil {
		panic(err)
	}

}
