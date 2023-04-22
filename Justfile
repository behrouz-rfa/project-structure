gen what="all":
    @just _gen_{{what}}

@_gen_gql:
    go get github.com/99designs/gqlgen@v0.17.24
    go run github.com/99designs/gqlgen generate