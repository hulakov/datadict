package services

import (
	"fmt"

	"github.com/hulakov/datadict/server/models"
	"github.com/hulakov/datadict/server/storage"
	"github.com/rs/zerolog/log"
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
	err := d.s.Add(key, value)
	if err != nil {
		log.Error().
			Str("key", key).
			Str("value", fmt.Sprintf("%v", value)).
			Err(err).
			Msg("add item")
	} else {
		log.Info().
			Str("key", key).
			Str("value", fmt.Sprintf("%v", value)).
			Err(err).
			Msg("add item")
	}
	return err
}

func (d *dict) Remove(key string) error {
	err := d.s.Remove(key)
	if err != nil {
		log.Info().
			Str("key", key).
			Err(err).
			Msg("remove item")
	} else {
		log.Error().
			Str("key", key).
			Err(err).
			Msg("remove item")
	}
	return err
}

func (d *dict) Get(key string) (*models.Data, error) {
	data, err := d.s.Get(key)
	if err != nil {
		log.Error().
			Str("key", key).
			Str("value", fmt.Sprintf("%v", data)).
			Err(err).
			Msg("get item")
	} else {
		log.Info().
			Str("key", key).
			Str("value", fmt.Sprintf("%v", data)).
			Err(err).
			Msg("get item")
	}
	return data, err
}

func (d *dict) GetAll() []models.Item {
	items := d.s.GetAll()
	log.Info().
		Str("items", fmt.Sprintf("%v", items)).
		Msg("get all items")
	return items
}
