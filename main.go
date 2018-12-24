package main

import (
	// "fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qor/admin"
	// "github.com/qor/qor"
	"net/http"
)

// Create a GORM-backend model
type User struct {
	gorm.Model
	Name string
}

// Create another GORM-backend model
type Product struct {
	gorm.Model
	Name        string
	Description string
}

func main() {
	DB, _ := gorm.Open("sqlite3", "demo.db")
	DB.AutoMigrate(&User{}, &Product{})

	// Initalize
	Admin := admin.New(&admin.AdminConfig{DB: DB})

	// Allow to use Admin to manage User, Product
	Admin.AddResource(&User{})
	Admin.AddResource(&Product{})

	// initalize an HTTP request multiplexer
	mux := http.NewServeMux()
	Admin.MountTo("/admin", mux)

	router := gin.Default()
	router.Any("/admin/*resources", gin.WrapH(mux))
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	// Start and run the server
	router.Run(":3000")
}
