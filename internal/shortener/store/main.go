package store

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

type Storer interface {
	Store(ID, URL string) error
}

type Request struct {
	URL string
}

func NewRequest(URL string) Request {
	return Request{URL: URL}
}

type Environment struct {
	storer      Storer
	idGenerator func() string
}

func NewEnvironment(s Storer) Environment {
	return Environment{storer: s, idGenerator: NewULID}
}

func NewULID() string {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return fmt.Sprintf("%v", id)
}

func URLStore(ctx context.Context, env Environment, req Request) (ID string, err error) {
	uid := env.idGenerator()

	if err := env.storer.Store(uid, req.URL); err != nil {
		return "", err
	}
	return uid, nil

}
