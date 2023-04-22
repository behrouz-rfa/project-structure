package gql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"project-structure/internal/module/user/presentation/gql/generated"
	"project-structure/internal/module/user/presentation/gql/interceptors"
	"project-structure/internal/module/user/presentation/gql/resolvers"
	"project-structure/pkg/common"
	"time"
)

const pingInterval = 30 * time.Second

type NewServerParams struct {
	*common.Services
}

func NewServer(params *NewServerParams) *handler.Server {
	config := generated.Config{
		Resolvers: resolvers.New(&resolvers.NewResolverParams{
			User: params.User,
		}),
	}

	schema := generated.NewExecutableSchema(config)

	srv := handler.NewDefaultServer(schema)
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: pingInterval,
	})

	srv.Use(extension.Introspection{})
	srv.AroundFields(interceptors.PreloadInterceptor)
	srv.SetErrorPresenter(interceptors.ErrorInterceptor)

	return srv
}
