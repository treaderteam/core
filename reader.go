package reader

import "gitlab.com/alexnikita/treader/reader/parsers"

// Book is primary book type
type Book struct {
	Author    string
	Title     string
	Coverpage string
	parsers.SpineStack
}

// Validate checks if file can be readed
func Validate(data []byte) (result bool) {
	return validate(data)
}

// Open opens book to reader
func (b *Book) Open(file []byte) error {
	return open(file, b)
}

// Get returns specified number of html elements
// from book with boundaries
func (b *Book) Get(count, from int) (result string, err error) {
	return
}
