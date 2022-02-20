package datamsg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddItemRequest(t *testing.T) {
	const json = `
    {
        "kind": "add_item",
        "key": "a",
        "value": {
          "foo": "test"
        }
    }`

	p, err := RequestParserFromJson(([]byte)(json))
	assert.Nil(t, err)

	m, err := p.ParseAddItem()
	assert.Nil(t, err)
	assert.Equal(t, m.Kind, ADD_ITEM_MESSAGE)
	assert.Equal(t, m.Key, "a")
	assert.Equal(t, m.Value, Data{Foo: "test"})
}

// TODO Implement test for other request types
