package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"
	"project-structure/pkg/module/user/application/dto"
	"project-structure/pkg/module/user/model/entity"
	"project-structure/pkg/module/user/model/filters"
	"project-structure/pkg/module/user/model/sort"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input dto.UserCreateInput) (*entity.User, error) {
	create, err := r.User.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	return &create, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, by filters.UserBy, input dto.UserUpdateInput) (*entity.User, error) {
	update, err := r.User.Update(ctx, *by.ID, input)
	if err != nil {
		return nil, err
	}
	return &update, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteUser - deleteUser"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, by filters.UserBy) (*entity.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, filter *filters.UserFilter, sort *sort.UserSort) ([]*entity.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}