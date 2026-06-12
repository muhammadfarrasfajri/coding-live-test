package main

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadfarrasfajri/coding-live-test/config"
)

type GetItemSeller struct {
	ItemName   string `json:"item_name"`
	SellerName string `json:"seller_name"`
}

func main() {
	config.Connect()
	db := config.Connect()
	defer db.Close()

	r := gin.Default()

	r.GET("/select-join-shop", func(c *gin.Context) {
		model := GetItemSeller{}
		query := `
            SELECT i.name AS item_name, s.name AS seller_name
            FROM items AS i
            JOIN sellers AS s ON i.sellerid = s.id
            WHERE s.rating > 4
            LIMIT 1;
        `
		err := db.QueryRow(query).Scan(&model.ItemName, &model.SellerName)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				c.JSON(http.StatusNotFound, gin.H{"error": "data tidak ditemukan"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, model)
	})

	r.Run(":8080")
}
