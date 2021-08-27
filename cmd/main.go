package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", "config.toml", "Path to config file")
}

func main() {
	flag.Parse()
	config := loadConfig(configPath)

	router := initRouter(config)

	err := router.Run(config.bindAddr)
	checkErr(err)
}

func loadConfig(path string) *Config {
	config := NewConfig()
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		log.Fatalf("Failed to load config file: %v", err)
	}

	return config
}

func initRouter(config *Config) *gin.Engine {
	router := gin.Default()
	gin.ForceConsoleColor()

	configureRoutes(router, config)
	return router
}

func configureRoutes(router *gin.Engine, config *Config) {

	router.GET("/", func(c *gin.Context) {
		var routes []string
		for _, r := range router.Routes() {
			routes = append(routes, fmt.Sprintf("%-6s %-25s", r.Method, r.Path))
		}

		response := fmt.Sprintf("%s\n\nRoutes:\n%v",
			"Welcome to Go Mock Server",
			strings.Join(routes, "\n"))

		c.String(200, response)
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	for _, route := range config.Routes {
		fullPath := mockPath(config.mocksFolder, route.Filename)
		router.GET(route.Path, jsonHandler(fullPath))
	}
}

func mockPath(mocksFolder string, filename string) string {
	if filename == "" {
		return filename
	}

	return fmt.Sprintf("%s/%s", mocksFolder, filename)
}

func jsonHandler(filename string) func(c *gin.Context) {
	return func(c *gin.Context) {
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			_ = c.AbortWithError(500, err)
		}
		c.Data(http.StatusOK, gin.MIMEJSON, content)
	}
}
