package datamsg

type MessageKind string

const (
	BAD_MESSAGE           MessageKind = "<bad>"
	ADD_ITEM_MESSAGE      MessageKind = "add_item"
	REMOVE_ITEM_MESSAGE   MessageKind = "remove_item"
	GET_ITEM_MESSAGE      MessageKind = "get_item"
	GET_ALL_ITEMS_MESSAGE MessageKind = "get_all_items"
)

var kindNames = map[MessageKind]interface{}{
	ADD_ITEM_MESSAGE:      nil,
	REMOVE_ITEM_MESSAGE:   nil,
	GET_ITEM_MESSAGE:      nil,
	GET_ALL_ITEMS_MESSAGE: nil,
}

func (kind MessageKind) IsValid() bool {
	_, ok := kindNames[kind]
	return ok
}
