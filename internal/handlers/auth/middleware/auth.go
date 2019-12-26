package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/jessequinn/go-gql-server/internal/logger"
	"github.com/jessequinn/go-gql-server/internal/orm"
	"github.com/jessequinn/go-gql-server/pkg/utils"

	"github.com/gin-gonic/gin"
)

func authError(c *gin.Context, err error) {
	errKey := "message"
	errMsgHeader := "[Auth] error: "
	e := gin.H{errKey: errMsgHeader + err.Error()}
	c.AbortWithStatusJSON(http.StatusUnauthorized, e)
}

// Middleware wraps the request with auth middleware
func Middleware(path string, cfg *utils.ServerConfig, orm *orm.ORM) gin.HandlerFunc {
	logger.Info("[Auth.Middleware] Applied to path: ", path)
	return gin.HandlerFunc(func(c *gin.Context) {
		if a, err := ParseAPIKey(c, cfg); err == nil {
			user, err := orm.FindUserByAPIKey(a)
			if err != nil {
				authError(c, ErrForbidden)
			}
			c.Request = addToContext(c, utils.ProjectContextKeys.UserCtxKey, user)
			logger.Info("User: ", user.ID)
			c.Next()
		} else {
			if err != ErrEmptyAPIKeyHeader {
				authError(c, err)
			} else {
				t, err := ParseToken(c, cfg)
				if err != nil {
					authError(c, err)
				} else {
					// goth.ContextForClient(c.)
					if claims, ok := t.Claims.(jwt.MapClaims); ok {
						if claims["exp"] != nil {
							issuer := claims["iss"].(string)
							userid := claims["jti"].(string)
							email := claims["email"].(string)
							if claims["aud"] != nil {
								audiences := claims["aud"].(interface{})
								logger.Warnf("\n\naudiences: %s\n\n", audiences)
							}
							if claims["alg"] != nil {
								algo := claims["alg"].(string)
								logger.Warnf("\n\nalgo: %s\n\n", algo)
							}
							// TODO: Verify token with each provider's JWKs
							if user, err := orm.FindUserByJWT(email, issuer, userid); err != nil {
								authError(c, ErrForbidden)
							} else {
								c.Request = addToContext(c, utils.ProjectContextKeys.UserCtxKey, user)
								c.Next()
							}
						} else {
							authError(c, ErrMissingExpField)
						}
					} else {
						authError(c, err)
					}
				}
			}
		}
	})
}
