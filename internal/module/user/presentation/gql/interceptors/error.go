package interceptors

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	cerr "project-structure/pkg/errors"
)

func ErrorInterceptor(ctx context.Context, err error) *gqlerror.Error {
	oerr := errors.Unwrap(err)

	if e, ok := oerr.(*cerr.Error); ok { //nolint:errorlint // errors.Is cause more linting issues with false positives
		return &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
			Extensions: map[string]interface{}{
				"code": e.Code(),
			},
		}
	}

	return &gqlerror.Error{
		Message: err.Error(),
		Path:    graphql.GetPath(ctx),
		Extensions: map[string]interface{}{
			"code": cerr.Code(cerr.ErrUnknown),
		},
	}
}
