package datamsg

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type ResponseParser interface {
	GetKind() MessageKind
	ParseAddItem() (*AddItemMessageResponse, error)
	ParseRemoveItem() (*RemoveItemMessageResponse, error)
	ParseGetItem() (*GetItemMessageResponse, error)
	ParseGetAllItems() (*GetAllItemsMessageResponse, error)
}

type responseParser struct {
	bytes []byte
	kind  MessageKind
}

func ResponseParserFromJson(bytes []byte) (ResponseParser, error) {
	var baseReq BaseMessageResponse
	err := json.Unmarshal(bytes, &baseReq)
	if err != nil {
		log.Debug().Err(err).Msg("cannot parse JSON")
		return nil, ErrBadJson
	}

	if !baseReq.Kind.IsValid() {
		return nil, ErrBadKind
	}

	return responseParser{bytes: bytes, kind: baseReq.Kind}, nil
}

func (p responseParser) GetKind() MessageKind {
	return p.kind
}

func (p responseParser) ParseAddItem() (*AddItemMessageResponse, error) {
	var response AddItemMessageResponse
	err := json.Unmarshal(p.bytes, &response)
	if err != nil {
		log.Debug().Err(err).Msg("cannot parse JSON")
		return nil, ErrBadJson
	}
	return &response, nil
}

func (p responseParser) ParseRemoveItem() (*RemoveItemMessageResponse, error) {
	var response RemoveItemMessageResponse
	err := json.Unmarshal(p.bytes, &response)
	if err != nil {
		log.Debug().Err(err).Msg("cannot parse JSON")
		return nil, ErrBadJson
	}
	return &response, nil
}

func (p responseParser) ParseGetItem() (*GetItemMessageResponse, error) {
	var response GetItemMessageResponse
	err := json.Unmarshal(p.bytes, &response)
	if err != nil {
		log.Debug().Err(err).Msg("cannot parse JSON")
		return nil, ErrBadJson
	}
	return &response, nil
}

func (p responseParser) ParseGetAllItems() (*GetAllItemsMessageResponse, error) {
	var response GetAllItemsMessageResponse
	err := json.Unmarshal(p.bytes, &response)
	if err != nil {
		log.Debug().Err(err).Msg("cannot parse JSON")
		return nil, ErrBadJson
	}
	return &response, nil
}
