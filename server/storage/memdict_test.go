package storage

import (
	"testing"

	"github.com/hulakov/datadict/server/models"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	m := New()
	get := func(key string) models.Data {
		v, err := m.Get(key)
		assert.Nil(t, err)
		return *v
	}

	getErr := func(key string) error {
		_, err := m.Get(key)
		return err
	}

	assert.Nil(t, m.Add("a", models.Data{Foo: "a"}))
	assert.Equal(t, models.Data{Foo: "a"}, get("a"))
	assert.Equal(t, ErrItemDoesNotExist, getErr("b"))
	assert.Equal(t, []models.Item{
		{Key: "a", Value: models.Data{Foo: "a"}},
	}, m.GetAll())

	assert.Nil(t, m.Add("b", models.Data{Foo: "b"}))
	assert.Equal(t, models.Data{Foo: "a"}, get("a"))
	assert.Equal(t, models.Data{Foo: "b"}, get("b"))
	assert.Equal(t, []models.Item{
		{Key: "a", Value: models.Data{Foo: "a"}},
		{Key: "b", Value: models.Data{Foo: "b"}},
	}, m.GetAll())

	assert.Nil(t, m.Add("c", models.Data{Foo: "c"}))
	assert.Equal(t, models.Data{Foo: "a"}, get("a"))
	assert.Equal(t, models.Data{Foo: "b"}, get("b"))
	assert.Equal(t, models.Data{Foo: "c"}, get("c"))
	assert.Equal(t, []models.Item{
		{Key: "a", Value: models.Data{Foo: "a"}},
		{Key: "b", Value: models.Data{Foo: "b"}},
		{Key: "c", Value: models.Data{Foo: "c"}},
	}, m.GetAll())

	assert.Equal(t, ErrItemAlreadyExist, m.Add("c", models.Data{Foo: "xx"}))
	assert.Equal(t, models.Data{Foo: "a"}, get("a"))
	assert.Equal(t, models.Data{Foo: "b"}, get("b"))
	assert.Equal(t, models.Data{Foo: "c"}, get("c"))
	assert.Equal(t, []models.Item{
		{Key: "a", Value: models.Data{Foo: "a"}},
		{Key: "b", Value: models.Data{Foo: "b"}},
		{Key: "c", Value: models.Data{Foo: "c"}},
	}, m.GetAll())

	assert.Nil(t, m.Remove("b"))
	assert.Equal(t, models.Data{Foo: "a"}, get("a"))
	assert.Equal(t, ErrItemDoesNotExist, getErr("b"))
	assert.Equal(t, models.Data{Foo: "c"}, get("c"))
	assert.Equal(t, []models.Item{
		{Key: "a", Value: models.Data{Foo: "a"}},
		{Key: "c", Value: models.Data{Foo: "c"}},
	}, m.GetAll())

	assert.Equal(t, ErrItemDoesNotExist, m.Remove("b"))
	assert.Equal(t, models.Data{Foo: "a"}, get("a"))
	assert.Equal(t, ErrItemDoesNotExist, getErr("b"))
	assert.Equal(t, models.Data{Foo: "c"}, get("c"))
	assert.Equal(t, []models.Item{
		{Key: "a", Value: models.Data{Foo: "a"}},
		{Key: "c", Value: models.Data{Foo: "c"}},
	}, m.GetAll())

}

func TestRemoveSingleItem(t *testing.T) {
	m := New()
	get := func(key string) models.Data {
		v, err := m.Get(key)
		assert.Nil(t, err)
		return *v
	}
	getErr := func(key string) error {
		_, err := m.Get(key)
		return err
	}

	assert.Nil(t, m.Add("a", models.Data{Foo: "a"}))
	assert.Equal(t, models.Data{Foo: "a"}, get("a"))

	assert.Nil(t, m.Remove("a"))
	assert.Equal(t, ErrItemDoesNotExist, getErr("a"))

	assert.Nil(t, m.Add("a", models.Data{Foo: "a"}))
	assert.Equal(t, models.Data{Foo: "a"}, get("a"))
}
