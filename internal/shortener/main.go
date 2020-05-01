package shortener

import (
	"context"
	"gopinionated/internal/shortener/resolve"
	"gopinionated/internal/shortener/store"
)

type Backend interface {
	resolve.Resolver
	store.Storer
}

type Controller struct {
	backend Backend
}

func NewController(b Backend) Controller {
	return Controller{backend: b}
}

func (c *Controller) ResolveID(ctx context.Context, ID string) (string, error) {
	url, err := resolve.URLResolve(
		ctx,
		resolve.NewEnvironment(c.backend),
		resolve.NewRequest(ID),
	)

	if err != nil {
		return "", err
	}
	return url, nil
}

func (c *Controller) Shorten(ctx context.Context, URL string) (string, error) {
	ID, err := store.URLStore(
		ctx,
		store.NewEnvironment(c.backend),
		store.NewRequest(URL),
	)
	if err != nil {
		return "", err
	}
	return ID, nil
}
