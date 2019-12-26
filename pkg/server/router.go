package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jessequinn/go-gql-server/internal/orm"
	"github.com/jessequinn/go-gql-server/pkg/server/routes"
	"github.com/jessequinn/go-gql-server/pkg/utils"
)

// RegisterRoutes register the routes for the server
func RegisterRoutes(cfg *utils.ServerConfig, r *gin.Engine, orm *orm.ORM) (err error) {
	// Auth routes
	if err = routes.Auth(cfg, r, orm); err != nil {
		return err
	}
	// GraphQL server routes
	if err = routes.GraphQL(cfg, r, orm); err != nil {
		return err
	}
	// Miscellaneous routes
	if err = routes.Misc(cfg, r, orm); err != nil {
		return err
	}
	return err
}
