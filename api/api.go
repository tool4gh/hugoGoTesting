package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tool4gh/hugoGoTesting/external"
)

func StartServer() {
	fmt.Println("Starting server")

	router := gin.Default()
	router.GET("/all", getAll)
	router.GET("/ping", ping)

	router.Run("localhost:8080")
}

func getAll(g *gin.Context) {
	res := external.Login()
	g.IndentedJSON(http.StatusOK, res)
}

func ping(g *gin.Context) { g.IndentedJSON(200, gin.H{"pong": "pong"}) }
