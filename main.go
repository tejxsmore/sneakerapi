package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	ConfigRuntime()
	StartGin()
}

// ConfigRuntime sets the number of operating system threads.
func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

func StartGin() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.File("index.html") // Provide the path to your index.html file
	})

	router.GET("/nike", func(c *gin.Context) {
		jsonData, err := os.ReadFile("store/nike.json")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read JSON file"})
			return
		}
		c.Header("Content-Type", "application/json")
		c.Data(http.StatusOK, "application/json", jsonData)
	})

	router.GET("/airmax", func(c *gin.Context) {
		jsonData, err := os.ReadFile("store/airmax.json")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read JSON file"})
			return
		}
		c.Header("Content-Type", "application/json")
		c.Data(http.StatusOK, "application/json", jsonData)
	})

	router.GET("/airforce", func(c *gin.Context) {
		jsonData, err := os.ReadFile("store/airforce.json")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read JSON file"})
			return
		}
		c.Header("Content-Type", "application/json")
		c.Data(http.StatusOK, "application/json", jsonData)
	})

	router.GET("/airjordan", func(c *gin.Context) {
		jsonData, err := os.ReadFile("store/airjordan.json")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read JSON file"})
			return
		}
		c.Header("Content-Type", "application/json")
		c.Data(http.StatusOK, "application/json", jsonData)
	})

	router.GET("/all", func(c *gin.Context) {
		jsonData, err := os.ReadFile("store/products.json")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read JSON file"})
			return
		}
		c.Header("Content-Type", "application/json")
		c.Data(http.StatusOK, "application/json", jsonData)
	})

	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}

func Asset(s string) {
	panic("unimplemented")
}
