package repository

import (
	monggDb "github.com/behrouz-rfa/mongo-specification/pkg/database/mongo"
	data "github.com/behrouz-rfa/mongo-specification/pkg/infrastructure/database"
	"github.com/google/uuid"
	"project-structure/pkg/module/user/model/entity"
	"project-structure/pkg/module/user/persistence/repository"
)

func NewUser(getter data.DataContextGetter) repository.User {
	r := monggDb.NewRepo(getter, data.NewDefaultMapper[User, entity.User]())
	return &r
}

type User struct {
	data.DocumentBase `bson:"inline"`
	Email             string
	FirstName         string
	LastName          string
	MiddleName        *string
	ProfilePicture    *string
	PhoneNumber       *string
	EmailVerified     bool
	Active            bool
	Role              string
}

func (d User) CollectionName() string {
	return "users"
}
func (u User) GetID() string {

	return u.ID
}

func (u User) SetID(id string) {
	u.ID = id
}

func (u User) GenerateID() {
	u.ID = uuid.New().String()
}
