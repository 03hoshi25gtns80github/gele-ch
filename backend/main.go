package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Memo struct {
    Date string `json:"date"`
    Memo string `json:"memo"`
}

func main() {
    router := gin.Default()

    db, err := sql.Open("mysql", "user:password@tcp(db:3306)/gelech")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    router.GET("/api/memos", func(c *gin.Context) {
        date := c.Query("date")
        var memo string
        err := db.QueryRow("SELECT memo FROM memos WHERE date = ?", date).Scan(&memo)
        if err != nil {
            if err == sql.ErrNoRows {
                c.JSON(http.StatusOK, gin.H{"memo": ""})
            } else {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            }
            return
        }
        c.JSON(http.StatusOK, gin.H{"memo": memo})
    })

    router.POST("/api/memos", func(c *gin.Context) {
        var memo Memo
        if err := c.ShouldBindJSON(&memo); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        _, err := db.Exec("INSERT INTO memos (date, memo) VALUES (?, ?) ON DUPLICATE KEY UPDATE memo = ?", memo.Date, memo.Memo, memo.Memo)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"status": "success"})
    })

    router.Run(":8080")
}