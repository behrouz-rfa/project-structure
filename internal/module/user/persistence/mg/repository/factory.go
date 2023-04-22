package repository

import (
	"github.com/behrouz-rfa/mongo-specification/pkg/infrastructure/database"
	"project-structure/pkg/module/user/persistence/repository"
)

type MongoRepoFactory struct {
	withTestID bool
}

func (m MongoRepoFactory) NewUser(getter database.DataContextGetter) repository.User {
	return NewUser(getter)
}

func NewMongoRepoFactory() MongoRepoFactory {
	return MongoRepoFactory{}
}
