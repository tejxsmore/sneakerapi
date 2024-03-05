package main

import (
	"encoding/json"
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
	// router.LoadHTMLGlob("resources/*.templ.html")

	router.GET("/featured", func(c *gin.Context) {
		jsonFile, err := os.ReadFile("store/featured.json")
		if err != nil {
			c.String(500, "Failed to read JSON file")
			return
		}

		var jsonData map[string]interface{}
		err = json.Unmarshal(jsonFile, &jsonData)
		if err != nil {
			c.String(500, "Failed to parse JSON data")
			return
		}

		c.JSON(200, jsonData)
	})

	router.GET("/airmax", func(c *gin.Context) {
		jsonFile, err := os.ReadFile("store/airmax.json")
		if err != nil {
			c.String(500, "Failed to read JSON file")
			return
		}

		var jsonData map[string]interface{}
		err = json.Unmarshal(jsonFile, &jsonData)
		if err != nil {
			c.String(500, "Failed to parse JSON data")
			return
		}

		c.JSON(200, jsonData)
	})

	router.GET("/airforce", func(c *gin.Context) {
		jsonFile, err := os.ReadFile("store/airforce.json")
		if err != nil {
			c.String(500, "Failed to read JSON file")
			return
		}

		var jsonData map[string]interface{}
		err = json.Unmarshal(jsonFile, &jsonData)
		if err != nil {
			c.String(500, "Failed to parse JSON data")
			return
		}

		c.JSON(200, jsonData)
	})

	router.GET("/airjordan", func(c *gin.Context) {
		jsonFile, err := os.ReadFile("store/airjordan.json")
		if err != nil {
			c.String(500, "Failed to read JSON file")
			return
		}

		var jsonData map[string]interface{}
		err = json.Unmarshal(jsonFile, &jsonData)
		if err != nil {
			c.String(500, "Failed to parse JSON data")
			return
		}

		c.JSON(200, jsonData)
	})

	router.GET("/all", func(c *gin.Context) {
		jsonFile, err := os.ReadFile("store/nike.json")
		if err != nil {
			c.String(500, "Failed to read JSON file")
			return
		}

		var jsonData map[string]interface{}
		err = json.Unmarshal(jsonFile, &jsonData)
		if err != nil {
			c.String(500, "Failed to parse JSON data")
			return
		}

		c.JSON(200, jsonData)
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
