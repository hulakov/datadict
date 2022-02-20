package datamsg

import "fmt"

type Data struct {
	Foo string `json:"foo"`
}

func (d *Data) String() string {
	return fmt.Sprintf("foo=%s", d.Foo)
}

type BaseMessageRequest struct {
	Kind MessageKind `json:"kind"`
}

type AddItemMessageRequest struct {
	BaseMessageRequest
	Key   string `json:"key"`
	Value Data   `json:"value"`
}

type RemoveItemMessageRequest struct {
	BaseMessageRequest
	Key string `json:"key"`
}

type GetItemMessageRequest struct {
	BaseMessageRequest
	Key string `json:"key"`
}

type GetAllItemsMessageRequest struct {
	BaseMessageRequest
}

type BaseMessageResponse struct {
	Kind  MessageKind `json:"kind"`
	Error *string     `json:"error,omitempty"`
}

type AddItemMessageResponse struct {
	BaseMessageResponse
	Key string `json:"key"`
}

type RemoveItemMessageResponse struct {
	BaseMessageResponse
	Key string `json:"key"`
}

type GetItemMessageResponse struct {
	BaseMessageResponse
	Key   string `json:"key"`
	Value Data   `json:"value"`
}

type DataItem struct {
	Key   string `json:"key"`
	Value Data   `json:"value"`
}

type GetAllItemsMessageResponse struct {
	BaseMessageResponse
	Items []DataItem `json:"items"`
}
