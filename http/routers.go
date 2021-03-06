package http

import (
	"github.com/gin-gonic/gin"
	"github.com/killtw/lemonade/lemonade"
	"net/http"
)

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to lemonade server",
	})
}

func heartbeatHandler(c *gin.Context) {
	c.AbortWithStatus(http.StatusOK)
}

func replaceHandler(c *gin.Context) {
	message := c.PostForm("message")
	filtered, matches := lemonade.Replace(message)

	c.JSON(http.StatusOK, gin.H{
		"original": message,
		"message":  filtered,
		"matches":  matches,
	})
}

func addWordHandler(c *gin.Context) {
	word := c.PostForm("word")
	lemonade.Add(word)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"word": word,
	})
}

func routers() *gin.Engine {
	router := gin.Default()

	router.POST("/replace", replaceHandler)
	router.POST("/addWord", addWordHandler)
	router.GET("/healthz", heartbeatHandler)
	router.GET("/", rootHandler)

	return router
}

