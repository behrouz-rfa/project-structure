package service

import (
	"context"
	"project-structure/pkg/module/user/application/dto"
	"project-structure/pkg/module/user/model/entity"
)

type User interface {
	Create(context.Context, dto.UserCreateInput) (entity.User, error)
	Update(context.Context, string, dto.UserUpdateInput) (entity.User, error)
}
