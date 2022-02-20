package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/hulakov/datadict/pkg/datamsg"
	"github.com/rs/zerolog/log"
)

func commandToJson(str string) ([]byte, error) {
	cc := strings.Split(str, ":")
	if len(cc) != 2 {
		return nil, errors.New("expected command format <cmd>:<arg1>,...,<argN>")
	}

	kind := datamsg.MessageKind(cc[0])
	args := strings.Split(cc[1], ",")

	switch kind {
	case datamsg.ADD_ITEM_MESSAGE:
		if len(args) != 2 {
			return nil, fmt.Errorf("expected command format %s:<key>,<value>", datamsg.ADD_ITEM_MESSAGE)
		}
		key := args[0]
		value := args[1]
		log.Info().
			Str("key", key).
			Str("value", value).
			Msg("add item")
		return json.Marshal(datamsg.AddItemMessageRequest{
			BaseMessageRequest: datamsg.BaseMessageRequest{Kind: datamsg.ADD_ITEM_MESSAGE},
			Key:                key,
			Value:              datamsg.Data{Foo: value},
		})

	case datamsg.REMOVE_ITEM_MESSAGE:
		if len(args) != 1 {
			return nil, fmt.Errorf("expected command format %s:<key>,<value>", datamsg.REMOVE_ITEM_MESSAGE)
		}
		key := args[0]
		log.Info().
			Str("key", key).
			Msg("remove item")
		return json.Marshal(datamsg.RemoveItemMessageRequest{
			BaseMessageRequest: datamsg.BaseMessageRequest{Kind: datamsg.REMOVE_ITEM_MESSAGE},
			Key:                key,
		})

	case datamsg.GET_ITEM_MESSAGE:
		if len(args) != 1 {
			return nil, fmt.Errorf("expected command format %s:<key>,<value>", datamsg.GET_ITEM_MESSAGE)
		}
		key := args[0]
		log.Info().
			Str("key", key).
			Msg("get item")
		return json.Marshal(datamsg.GetItemMessageRequest{
			BaseMessageRequest: datamsg.BaseMessageRequest{Kind: datamsg.GET_ITEM_MESSAGE},
			Key:                key,
		})

	case datamsg.GET_ALL_ITEMS_MESSAGE:
		if len(args) != 0 {
			return nil, fmt.Errorf("expected command format %s:<key>,<value>", datamsg.GET_ITEM_MESSAGE)
		}
		log.Info().
			Msg("get all items")
		return json.Marshal(datamsg.GetAllItemsMessageRequest{
			BaseMessageRequest: datamsg.BaseMessageRequest{Kind: datamsg.GET_ITEM_MESSAGE},
		})

	default:
		return nil, fmt.Errorf("bad kind of message: %s", kind)
	}
}
