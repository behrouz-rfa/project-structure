package service

import (
	"context"
	"github.com/behrouz-rfa/mongo-specification/pkg/infrastructure/database"
	"github.com/behrouz-rfa/mongo-specification/pkg/mspecification"
	"github.com/behrouz-rfa/nilmapper"
	"gopkg.in/jeevatkm/go-model.v1"
	"project-structure/pkg/module/user/application/dto"
	pkgService "project-structure/pkg/module/user/application/service"
	"project-structure/pkg/module/user/model/entity"
	"project-structure/pkg/module/user/persistence/repository"
)

type User struct {
	repository.RepoFactory
	transactionFactory database.TransactionFactoryGetter
}

var _ pkgService.User = (*User)(nil)

func NewUser(repoFactory repository.RepoFactory, transactionFactory database.TransactionFactoryGetter) *User {
	return &User{RepoFactory: repoFactory, transactionFactory: transactionFactory}
}

func (u User) Create(ctx context.Context, input dto.UserCreateInput) (entity.User, error) {
	var userEntity entity.User
	model.Copy(&userEntity, input)
	factory, _ := u.transactionFactory.GetTransactionFactory()
	tx := factory.New()
	tx.Begin(ctx)
	userRepo := u.RepoFactory.NewUser(tx)
	id, err := userRepo.Create(ctx, userEntity)
	if err != nil {
		return entity.User{}, err
	}
	tx.Commit(ctx)

	spec := mspecification.NewBaseSpecification()
	spec.WithContext(ctx)
	spec.FilterByID(id)
	item, err := userRepo.FindOneBy(ctx, spec)

	return item, nil
}

func (u User) Update(ctx context.Context, id string, input dto.UserUpdateInput) (entity.User, error) {
	var userEntity entity.User
	nilmapper.MapStruct(input, &userEntity)

	factory, _ := u.transactionFactory.GetTransactionFactory()
	tx := factory.New()
	tx.Begin(ctx)
	userRepo := u.RepoFactory.NewUser(tx)
	user, err := userRepo.Update(ctx, id, userEntity)
	if err != nil {
		return entity.User{}, err
	}
	tx.Commit(ctx)

	return user, nil
}
