package store

import (
	"context"
	"crypto/sha1"
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
	idGenerator func(length int) string
	idMaxLength int
}

func NewEnvironment(s Storer) Environment {
	return Environment{storer: s, idGenerator: NewID, idMaxLength: 6}
}

func NewID(length int) string {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)

	h := sha1.New()
	h.Write([]byte(fmt.Sprintf("%s", id)))
	return fmt.Sprintf("%x", h.Sum(nil))[:length]
}

func URLStore(ctx context.Context, env Environment, req Request) (ID string, err error) {
	uid := env.idGenerator(env.idMaxLength)

	if err := env.storer.Store(uid, req.URL); err != nil {
		return "", err
	}
	return uid, nil

}
