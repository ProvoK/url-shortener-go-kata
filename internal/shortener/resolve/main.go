package resolve

import (
	"context"
	"errors"
	"fmt"
)

type Resolver interface {
	Resolve(string) (string, bool)
}

type Request struct {
	ID string
}

func NewRequest(ID string) Request {
	return Request{ID: ID}
}

type Environment struct {
	resolver Resolver
}

func NewEnvironment(r Resolver) Environment {
	return Environment{resolver: r}
}

func URLResolve(ctx context.Context, env Environment, req Request) (URL string, err error) {
	URL, ok := env.resolver.Resolve(req.ID)
	if !ok {
		err = errors.New(fmt.Sprintf("UrlResolve: failed to fetch %v", req.ID))
		return "", err
	}
	return URL, nil
}
