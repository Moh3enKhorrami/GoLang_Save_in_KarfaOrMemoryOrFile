package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FileStack struct {
	filename string
	file     *os.File
}

func (f *FileStack) Push(item int) error {
	_, err := f.file.WriteString(fmt.Sprintf("%d\n", item))
	return err
}

func (f *FileStack) Pop() (int, error) {
	fileContent, err := os.ReadFile(f.filename)
	if err != nil {
		return 0, err
	}

	lines := strings.Split(strings.TrimSpace(string(fileContent)), "\n")
	if len(lines) == 0 {
		return 0, errors.New("stack is empty")
	}

	lastLine := lines[len(lines)-1]
	lastItem, err := strconv.Atoi(lastLine)
	if err != nil {
		return 0, err
	}

	err = os.WriteFile(f.filename, []byte(strings.Join(lines[:len(lines)-1], "\n")), 0644)
	if err != nil {
		return 0, err
	}

	return lastItem, nil
}

func NewFile(filename string) (*FileStack, error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return &FileStack{filename: filename, file: file}, nil
}
