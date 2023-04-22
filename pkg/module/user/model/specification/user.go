package specification

import (
	specification "github.com/behrouz-rfa/mongo-specification/pkg/infrastructure/database/specefication"
	"project-structure/pkg/module/user/model/filters"
	"project-structure/pkg/module/user/model/sort"
)

type UserSpecification interface {
	specification.Set
	By(filter filters.UserBy) UserSpecification
	Filter(filter *filters.UserFilter) UserSpecification
	SortBy(sort *sort.UserSort) UserSpecification
}
