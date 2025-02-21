package main

import (
	"net/http"
	//   "gorm.io/driver/postgres"
	//   "gorm.io/gorm"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Creative-genius001/Connekt/cmd/api/routes"
	"github.com/Creative-genius001/Connekt/config"
	limit "github.com/aviddiviner/gin-limit"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	router := gin.Default()
	router.Use(limit.MaxAllowed(200))

	//initialise DB
	config.ConnectDB()

	//init routes
	routes.InitializeRoutes(router)

	// Configure CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AddAllowHeaders("Authorization")
	corsConfig.AllowOrigins = []string{"*"}
	router.Use(cors.New(corsConfig))

	//startup server
	PORT := os.Getenv("PORT")
	fmt.Println("runnning on port:", PORT)
	s := &http.Server{
		Addr:           ":" + PORT,
		Handler:        router,
		ReadTimeout:    18000 * time.Second,
		WriteTimeout:   18000 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if s.ListenAndServe(); err != nil {
		log.Fatal("Failed to start server: %v", err)
	}
	fmt.Sprintf("Server is running on port: %v", PORT)
}
