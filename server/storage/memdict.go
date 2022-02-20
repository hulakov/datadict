package storage

import (
	"sync"

	"github.com/hulakov/datadict/server/models"
)

type item struct {
	key   string
	value models.Data
	prev  *item
	next  *item
}

type memdict struct {
	head  *item
	tail  *item
	index map[string]*item
	m     sync.RWMutex
}

func New() Dict {
	fake := &item{}

	return &memdict{
		head:  fake,
		tail:  fake,
		index: make(map[string]*item),
	}
}

func (m *memdict) Add(key string, value models.Data) error {
	m.m.Lock()
	defer m.m.Unlock()

	_, ok := m.index[key]
	if ok {
		return ErrItemAlreadyExist
	}

	newItem := &item{
		key:   key,
		value: value,
		prev:  m.tail,
		next:  nil,
	}
	m.tail.next = newItem
	m.tail = newItem
	m.index[key] = newItem

	return nil
}

func (m *memdict) Remove(key string) error {
	m.m.Lock()
	defer m.m.Unlock()

	item, ok := m.index[key]
	if !ok {
		return ErrItemDoesNotExist
	}

	item.prev.next = item.next
	if item.next != nil {
		item.next.prev = item.prev
	}
	delete(m.index, key)

	return nil
}

func (m *memdict) Get(key string) (*models.Data, error) {
	m.m.RLock()
	defer m.m.RUnlock()

	item, ok := m.index[key]
	if !ok {
		return nil, ErrItemDoesNotExist
	}

	return &item.value, nil
}

func (m *memdict) GetAll() []models.Item {
	m.m.RLock()
	defer m.m.RUnlock()

	var result []models.Item
	current := m.head.next
	for current != nil {
		result = append(result, models.Item{
			Key:   current.key,
			Value: current.value,
		})
		current = current.next
	}

	return result
}
