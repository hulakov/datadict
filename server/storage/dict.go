package storage

import (
	"errors"

	"github.com/hulakov/datadict/server/models"
)

var (
	ErrItemAlreadyExist = errors.New("item with such key already exist")
	ErrItemDoesNotExist = errors.New("item with such key doesn't exist")
)

type Dict interface {
	Add(key string, value models.Data) error
	Remove(key string) error
	Get(key string) (*models.Data, error)
	GetAll() []models.Item
}
