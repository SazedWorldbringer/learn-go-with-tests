package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

// Search a definition in the dictionary 
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add a definition to the dictionary
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
