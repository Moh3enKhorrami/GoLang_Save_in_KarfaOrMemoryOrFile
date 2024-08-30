package main

type StackStorage interface {
	Push(item int) error
	Pop() (int, error)
}
