package main

import "errors"

type Dictionary map[string]string

var ErrKeyNotFound = errors.New("no such key")

func (d *Dictionary) Search(key string) (string, error) {
	result, ok := (*d)[key]
	if !ok {
		return "", ErrKeyNotFound
	}
	return result, nil
}

func (d *Dictionary) Add(key, value string) {
	(*d)[key] = value
}
