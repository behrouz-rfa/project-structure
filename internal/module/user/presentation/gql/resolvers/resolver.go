package resolvers

import "project-structure/internal/module/user/application/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	User *service.User
}

type NewResolverParams struct {
	User *service.User
}

func New(params *NewResolverParams) *Resolver {
	return &Resolver{
		User: params.User,
	}
}
