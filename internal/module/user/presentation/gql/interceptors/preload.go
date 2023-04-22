package interceptors

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"project-structure/pkg/common"
)

// PreloadInterceptor extracts subfields from query and add it to the context under preload
//
//	result map[string]{book}
func PreloadInterceptor(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
	newCtx := context.WithValue(ctx, common.PreloadersKey, GetPreloads(ctx))

	return next(newCtx)
}

func GetPreloads(ctx context.Context) []string {
	return GetNestedPreloads(
		graphql.GetOperationContext(ctx),
		graphql.CollectFieldsCtx(ctx, nil),
		"",
	)
}

func GetNestedPreloads(
	ctx *graphql.OperationContext,
	fields []graphql.CollectedField,
	prefix string,
) (preloads []string) {
	for _, column := range fields {
		prefixColumn := GetPreloadString(prefix, column.Name)
		preloads = append(preloads, prefixColumn)
		preloads = append(
			preloads,
			GetNestedPreloads(ctx, graphql.CollectFields(ctx, column.Selections, nil),
				prefixColumn)...,
		)
	}

	return preloads
}

func GetPreloadString(prefix, name string) string {
	if len(prefix) > 0 {
		return prefix + "." + name
	}

	return name
}
