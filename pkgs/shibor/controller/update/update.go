package update

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Zhangbokai614/Automated-Teller-Machine/pkgs/shibor/model/mysql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func OnUpdate() {
	update := gin.Default()
	update.POST("/update", func(c *gin.Context) {
		// c.Bind()
		data, err := ioutil.ReadAll(c.Request.Body)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
		if err != nil {
			fmt.Println(err)
		}
		send(data)
	})
	update.Run(":8080")
}

func send(data []byte) {
	var jsonData mysql.ShiborTable

	if err := json.Unmarshal(data, &jsonData); err != nil {
		fmt.Println("error:", err)
		return
	}

	db, err := sql.Open("mysql", "root:123456@(192.168.0.251:6666)/project")
	if err != nil {
		fmt.Println(err)
	}

	if err := db.Ping(); err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	mysql.Insert(jsonData, db)
}
