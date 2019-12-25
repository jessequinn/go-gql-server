package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jessequinn/go-gql-server/internal/handlers"
	"github.com/jessequinn/go-gql-server/pkg/utils"
)

var HOST, PORT string

func init() {
	HOST = utils.MustGet("GQL_SERVER_HOST")
	PORT = utils.MustGet("GQL_SERVER_PORT")
}

// Run api server
func Run() {
	r := gin.Default()
	// Simple keep-alive/ping handler
	r.GET("/ping", handlers.Ping())
	// Inform the user where the server is listening
	log.Println("Running @ http://" + HOST + ":" + PORT)
	// Print out and exit(1) to the OS if the server cannot run
	log.Fatalln(r.Run(HOST + ":" + PORT))
}
