package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jessequinn/go-gql-server/internal/handlers/auth"
	"github.com/jessequinn/go-gql-server/internal/orm"
	"github.com/jessequinn/go-gql-server/pkg/utils"
)

// Auth routes
func Auth(cfg *utils.ServerConfig, r *gin.Engine, orm *orm.ORM) error {
	// OAuth handlers
	g := r.Group(cfg.VersionedEndpoint("/auth"))
	g.GET("/:provider", auth.Begin())
	g.GET("/:provider/callback", auth.Callback(cfg, orm))
	// g.GET(:provider/refresh", auth.Refresh(cfg, orm))
	return nil
}
