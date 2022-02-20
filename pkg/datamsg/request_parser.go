package datamsg

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type RequestParser interface {
	GetKind() MessageKind
	ParseAddItem() (*AddItemMessageRequest, error)
	ParseRemoveItem() (*RemoveItemMessageRequest, error)
	ParseGetItem() (*GetItemMessageRequest, error)
	ParseGetAllItems() (*GetAllItemsMessageRequest, error)
}

type requestParser struct {
	bytes []byte
	kind  MessageKind
}

func RequestParserFromJson(bytes []byte) (RequestParser, error) {
	var baseReq BaseMessageRequest
	err := json.Unmarshal(bytes, &baseReq)
	if err != nil {
		log.Debug().Err(err).Msg("cannot parse JSON")
		return nil, ErrBadJson
	}

	if !baseReq.Kind.IsValid() {
		return nil, ErrBadKind
	}

	return requestParser{bytes: bytes, kind: baseReq.Kind}, nil
}

func (p requestParser) GetKind() MessageKind {
	return p.kind
}

func (p requestParser) ParseAddItem() (*AddItemMessageRequest, error) {
	var request AddItemMessageRequest
	err := json.Unmarshal(p.bytes, &request)
	if err != nil {
		log.Debug().Err(err).Msg("cannot parse JSON")
		return nil, ErrBadJson
	}
	return &request, nil
}

func (p requestParser) ParseRemoveItem() (*RemoveItemMessageRequest, error) {
	var request RemoveItemMessageRequest
	err := json.Unmarshal(p.bytes, &request)
	if err != nil {
		log.Debug().Err(err).Msg("cannot parse JSON")
		return nil, ErrBadJson
	}
	return &request, nil
}

func (p requestParser) ParseGetItem() (*GetItemMessageRequest, error) {
	var request GetItemMessageRequest
	err := json.Unmarshal(p.bytes, &request)
	if err != nil {
		log.Debug().Err(err).Msg("cannot parse JSON")
		return nil, ErrBadJson
	}
	return &request, nil
}

func (p requestParser) ParseGetAllItems() (*GetAllItemsMessageRequest, error) {
	var request GetAllItemsMessageRequest
	err := json.Unmarshal(p.bytes, &request)
	if err != nil {
		log.Debug().Err(err).Msg("cannot parse JSON")
		return nil, ErrBadJson
	}
	return &request, nil
}
