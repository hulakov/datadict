package main

import (
	"fmt"

	"github.com/hulakov/datadict/pkg/datamsg"
	"github.com/rs/zerolog/log"
)

func printResponse(bytes []byte) {
	p, err := datamsg.ResponseParserFromJson(bytes)
	if err != nil {
		log.Error().Str("response", string(bytes)).Err(err).Msg("bad response")
		return
	}

	switch p.GetKind() {
	case datamsg.ADD_ITEM_MESSAGE:
		response, err := p.ParseAddItem()
		if err != nil {
			log.Error().
				Str("response", string(bytes)).
				Str("kind", string(datamsg.ADD_ITEM_MESSAGE)).
				Err(err).
				Msg("bad response")
		} else if response.Error != nil {
			log.Error().
				Str("key", response.Key).
				Str("error", *response.Error).
				Msg("cannot add item")
		} else {
			log.Info().
				Str("key", response.Key).
				Msg("item was added")
		}
	case datamsg.REMOVE_ITEM_MESSAGE:
		response, err := p.ParseRemoveItem()
		if err != nil {
			log.Error().
				Str("response", string(bytes)).
				Str("kind", string(datamsg.REMOVE_ITEM_MESSAGE)).
				Err(err).
				Msg("bad response")
		} else if response.Error != nil {
			log.Error().
				Str("key", response.Key).
				Str("error", *response.Error).
				Msg("cannot remove item")
		} else {
			log.Info().
				Str("key", response.Key).
				Msg("item was removed")
		}
	case datamsg.GET_ITEM_MESSAGE:
		response, err := p.ParseGetItem()
		if err != nil {
			log.Error().
				Str("response", string(bytes)).
				Str("kind", string(datamsg.GET_ITEM_MESSAGE)).
				Err(err).
				Msg("bad response")
		} else if response.Error != nil {
			log.Error().
				Str("key", response.Key).
				Str("error", *response.Error).
				Msg("cannot retrieve item")
		} else {
			log.Info().
				Str("key", response.Key).
				Str("value", response.Value.String()).
				Msg("item was retrieved")
		}
	case datamsg.GET_ALL_ITEMS_MESSAGE:
		response, err := p.ParseGetAllItems()
		if err != nil {
			log.Error().
				Str("response", string(bytes)).
				Str("kind", string(datamsg.GET_ALL_ITEMS_MESSAGE)).
				Err(err).
				Msg("bad response")
		} else {
			log.Info().
				Str("items", fmt.Sprintf("%v", response.Items)).
				Msg("items were retrieved")
		}

	}
	_, _ = p, err

	log.Debug().Str("response", string(bytes)).Msg("response")

}
