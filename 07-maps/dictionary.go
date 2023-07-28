package main

type Dictionary map[string]string

type DictionaryErr string

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

// Search a definition in the dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add a definition to the dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		// add word if it doesn't exist
		d[word] = definition
	case nil:
		// return error if it already exists
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) {
	d[word] = definition
}
