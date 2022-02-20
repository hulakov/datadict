package rabbitmq

import (
	"encoding/json"
	"errors"

	"github.com/hulakov/datadict/pkg/datamsg"
	"github.com/hulakov/datadict/pkg/rabbitrpc"
	"github.com/hulakov/datadict/server/models"
	"github.com/hulakov/datadict/server/services"
	"github.com/hulakov/datadict/server/storage"
	"github.com/rs/zerolog/log"
)

type dataDictPort struct {
	dict services.Dict
}

func convertItems(items []models.Item) []datamsg.DataItem {
	result := make([]datamsg.DataItem, 0, len(items))
	for _, item := range items {
		result = append(result, datamsg.DataItem{
			Key:   item.Key,
			Value: datamsg.Data(item.Value),
		})
	}

	return result
}

func errToString(err error) *string {
	if err == nil {
		return nil
	}

	s := err.Error()
	return &s
}

func (p *dataDictPort) processErr(request []byte) ([]byte, error) {
	parser, err := datamsg.RequestParserFromJson(request)
	if err != nil {
		return json.Marshal(datamsg.BaseMessageResponse{
			Kind:  datamsg.BAD_MESSAGE,
			Error: errToString(err),
		})
	}

	switch parser.GetKind() {
	case datamsg.ADD_ITEM_MESSAGE:
		request, err := parser.ParseAddItem()
		if err == nil {
			err = p.dict.Add(request.Key, models.Data(request.Value))
		}

		return json.Marshal(datamsg.AddItemMessageResponse{
			BaseMessageResponse: datamsg.BaseMessageResponse{
				Kind:  parser.GetKind(),
				Error: errToString(err),
			},
			Key: request.Key,
		})

	case datamsg.REMOVE_ITEM_MESSAGE:
		request, err := parser.ParseRemoveItem()
		if err == nil {
			err = p.dict.Remove(request.Key)
		}

		return json.Marshal(datamsg.RemoveItemMessageResponse{
			BaseMessageResponse: datamsg.BaseMessageResponse{
				Kind:  parser.GetKind(),
				Error: errToString(err),
			},
			Key: request.Key,
		})

	case datamsg.GET_ITEM_MESSAGE:
		request, err := parser.ParseGetItem()
		var data models.Data
		if err == nil {
			var dataPtr *models.Data
			dataPtr, err = p.dict.Get(request.Key)
			if dataPtr != nil {
				data = *dataPtr
			}
		}

		return json.Marshal(datamsg.GetItemMessageResponse{
			BaseMessageResponse: datamsg.BaseMessageResponse{
				Kind:  parser.GetKind(),
				Error: errToString(err),
			},
			Key:   request.Key,
			Value: datamsg.Data(data),
		})

	case datamsg.GET_ALL_ITEMS_MESSAGE:
		_, err := parser.ParseGetAllItems()

		return json.Marshal(datamsg.GetAllItemsMessageResponse{
			BaseMessageResponse: datamsg.BaseMessageResponse{
				Kind:  parser.GetKind(),
				Error: errToString(err),
			},
			Items: convertItems(p.dict.GetAll()),
		})

	default:
		return json.Marshal(datamsg.BaseMessageResponse{
			Kind:  datamsg.BAD_MESSAGE,
			Error: errToString(errors.New("bad message kind")),
		})
	}
}

func (p *dataDictPort) process(request []byte) []byte {
	b, err := p.processErr(request)
	if err != nil {
		log.Error().Err(err).Msg("error serializing to JSON")
		return nil
	}

	return b
}

func Run(rabbitMQHost string) error {
	s := storage.New()
	p := &dataDictPort{
		dict: services.New(s),
	}
	return rabbitrpc.RunServer(rabbitMQHost, p.process)
}
