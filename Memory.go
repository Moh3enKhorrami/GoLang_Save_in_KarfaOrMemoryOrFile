package main

import "errors"

type MemoryStack struct {
	items []int
}

func (m *MemoryStack) Push(item int) error {
	m.items = append(m.items, item)
	return nil
}

func (m *MemoryStack) Pop() (int, error) {
	if len(m.items) == 0 {
		return 0, errors.New("stack is empty")
	}
	lastItem := len(m.items) - 1
	item := m.items[lastItem]
	m.items = m.items[:lastItem]
	return item, nil
}
