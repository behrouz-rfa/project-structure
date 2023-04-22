package factory

import (
	"errors"
	mgRepo "project-structure/internal/module/user/persistence/mg/repository"
	"project-structure/pkg/module/user/persistence/repository"
)

const (
	Test = "test"
	Pg   = "gorm-pg"
	Data = "data"
)

var (
	ErrNotImplemented   = errors.New("not implemented")
	ErrMissingParameter = errors.New("missing parameter")
)

func NewRepoFactory() repository.RepoFactory {

	return mgRepo.NewMongoRepoFactory()

}
