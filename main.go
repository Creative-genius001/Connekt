package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Creative-genius001/Connekt/cmd/api/routes"
	"github.com/Creative-genius001/Connekt/cmd/middleware"
	"github.com/Creative-genius001/Connekt/config"
	"github.com/Creative-genius001/Connekt/utils"
	limit "github.com/aviddiviner/gin-limit"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	expectedHost := "localhost:" + PORT
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(limit.MaxAllowed(200))
	router.Use(func(c *gin.Context) {
		if c.Request.Host != expectedHost {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid host header"})
			return
		}
		c.Header("X-Frame-Options", "DENY")
		c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Header("Referrer-Policy", "strict-origin")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
		c.Next()
	})

	//initialise DB
	config.ConnectDB()

	//loggers initialization
	router.Use(middleware.LoggerMiddleware())
	utils.InitLogger()

	//init routes
	routes.InitializeRoutes(router)

	// Configure CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AddAllowHeaders("Authorization")
	corsConfig.AllowOrigins = []string{"*"}
	router.Use(cors.New(corsConfig))

	//startup server

	s := &http.Server{
		Addr:           ":" + PORT,
		Handler:        router,
		ReadTimeout:    18000 * time.Second,
		WriteTimeout:   18000 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	utils.Info("Server is starting and running.......", logrus.Fields{"port": PORT})
	if s.ListenAndServe(); err != nil {
		utils.Error("Failed to start server ", err, nil)
	}

}
