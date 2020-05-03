package store

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type StoreMock struct {
	mock.Mock
}

func (m *StoreMock) Store(ID, URL string) error {
	args := m.Called(ID, URL)
	return args.Error(0)
}

func TestExample(t *testing.T) {
	assert.Equal(t, 123, 123, "ga")
}

func TestURLStoreOK(t *testing.T) {
	assert := assert.New(t)
	storer := new(StoreMock)
	ID, URL := "dummy-id", "http://test.org"

	storer.On("Store", ID, URL).Return(nil)

	shortID, err := URLStore(
		context.TODO(),
		Environment{storer: storer, idGenerator: func(a int) string { return ID }},
		Request{URL: URL},
	)

	assert.Nil(err)
	assert.Equal(ID, shortID)
}

func TestURLStoreKO(t *testing.T) {
	assert := assert.New(t)
	storer := new(StoreMock)
	err := errors.New("dummy-err")

	storer.On("Store", mock.Anything, mock.Anything).Return(err)

	_, rErr := URLStore(
		context.TODO(),
		Environment{storer: storer, idGenerator: func(a int) string { return "123" }},
		Request{URL: "dummy"},
	)

	assert.Equal(err, rErr)
	assert.Error(err)
}

func TestDefaultGeneratorIsID(t *testing.T) {
	s := StoreMock{}
	assert.IsType(t, NewID, NewEnvironment(&s).idGenerator)
}

func TestNewRequestConstructor(t *testing.T) {
	req := NewRequest("url")
	assert.Equal(t, "url", req.URL)
}

func TestNewIDReturnsString(t *testing.T) {
	for i := 1; i <= 26; i++ {
		ID := NewID(i)
		assert.IsType(t, "", ID)
		assert.Equal(t, i, len(ID))
	}
}

func TestNewIDReturnsChangesWithInvokation(t *testing.T) {
	assert.NotEqual(t, NewID(6), NewID(6))
}
