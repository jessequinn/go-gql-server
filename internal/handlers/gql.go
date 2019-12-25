package handlers

import (
	//"github.com/99designs/gqlgen/handler"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/jessequinn/go-gql-server/internal/gql"
	"github.com/jessequinn/go-gql-server/internal/gql/resolvers"
)

// GraphqlHandler defines the GQLGen GraphQL server handler
func GraphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	c := gql.Config{
		Resolvers: &resolvers.Resolver{},
	}

	h := handler.New(gql.NewExecutableSchema(c))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// PlaygroundHandler defines a handler to expose the Playground
func PlaygroundHandler(path string) gin.HandlerFunc {
	h := playground.Handler("Go GraphQL Server", path)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
