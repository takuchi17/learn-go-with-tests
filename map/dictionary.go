package main

type Dictionary map[string]string

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("word already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	difinition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return difinition, nil
}

func (d Dictionary) Add(word string, definition string) error {
	if _, err := d.Search(word); err == nil {
		return ErrWordExists
	}
	d[word] = definition
	return nil
}
