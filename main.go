package main

import (
	"net/http"
	//   "gorm.io/driver/postgres"
	//   "gorm.io/gorm"
	"os"
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/gin-contrib/cors"
	_ "github.com/lib/pq"
	limit "github.com/aviddiviner/gin-limit"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	router := gin.Default()
	router.Use(limit.MaxAllowed(200))

	config.ConnectDB()


    // Configure CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AddAllowHeaders("Authorization")
	corsConfig.AllowOrigins = []string{"*"} 
	router.Use(cors.New(corsConfig))

	PORT := os.Getenv("PORT")
	fmt.Println("runnning on port:", PORT)
	s := &http.Server{
			Addr:           ":" + PORT,
			Handler:        router,
			ReadTimeout:    18000 * time.Second,
			WriteTimeout:   18000 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
	
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		})
	})
	if s.ListenAndServe()
	err != nil {
		panic(err)
		fmt.Printf("Failed to start server: %v", err)
		}
}