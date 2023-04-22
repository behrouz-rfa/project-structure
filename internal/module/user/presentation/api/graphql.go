package api

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"project-structure/internal/module/user/presentation/gql"
)

func (s *apiServer) RegisterGraphql() {
	playgroundHandler := playground.Handler("GraphQL", "/gql")

	graphql := s.gin.Group("/gql", GinToGqlMiddleware())
	{
		graphql.GET("", func(c *gin.Context) {
			playgroundHandler.ServeHTTP(c.Writer, c.Request)
		})

		graphql.POST("", func(c *gin.Context) {
			gql.
				NewServer(&gql.NewServerParams{Services: s.services}).
				ServeHTTP(c.Writer, c.Request)
		})
	}
}
