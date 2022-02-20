package services

import (
	"github.com/hulakov/datadict/server/models"
	"github.com/hulakov/datadict/server/storage"
)

type Dict interface {
	Add(key string, value models.Data) error
	Remove(key string) error
	Get(key string) (*models.Data, error)
	GetAll() []models.Item
}

var (
	ErrItemAlreadyExist = storage.ErrItemAlreadyExist
	ErrItemDoesNotExist = storage.ErrItemDoesNotExist
)

type dict struct {
	s storage.Dict
}

func New(s storage.Dict) Dict {
	return &dict{s: s}
}

func (d *dict) Add(key string, value models.Data) error {
	return d.s.Add(key, value)
}

func (d *dict) Remove(key string) error {
	return d.s.Remove(key)
}

func (d *dict) Get(key string) (*models.Data, error) {
	return d.s.Get(key)
}

func (d *dict) GetAll() []models.Item {
	return d.s.GetAll()
}
