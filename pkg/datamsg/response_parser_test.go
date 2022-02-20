package datamsg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllItemsResponse(t *testing.T) {
	const json = `
    {
        "kind": "get_all_items",
        "items": [
            {
                "key": "a",
                "value": {
                  "foo": "test1"
                }
            },
            {
                "key": "b",
                "value": {
                    "foo": "test2"
                }
            }
        ]
    }`

	p, err := ResponseParserFromJson(([]byte)(json))
	assert.Nil(t, err)

	m, err := p.ParseGetAllItems()
	assert.Nil(t, err)
	assert.Equal(t, m.Kind, GET_ALL_ITEMS_MESSAGE)
	assert.Nil(t, m.Error)
	assert.Equal(t, m.Items, []DataItem{
		{Key: "a", Value: Data{Foo: "test1"}},
		{Key: "b", Value: Data{Foo: "test2"}},
	})
}

// TODO Implement test for other reponse types
