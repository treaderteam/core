package reader

import (
	"errors"

	"gitlab.com/alexnikita/treader/reader/parsers"

	"gitlab.com/alexnikita/gols/epub"
	"gitlab.com/alexnikita/treader/reader/getters"
)

// Book is primary book type
type Book struct {
	Author    string
	Title     string
	Coverpage string
	Extension string
	Hash      string
	spine     parsers.SpineStack
	entity    interface{}
}

// SpineNext alias to spine stack next
func (b *Book) SpineNext() bool {
	return b.spine.Next()
}

// SpineValue alias to spine value
func (b *Book) SpineValue() string {
	return b.spine.Value()
}

// Validate checks if file can be readed
func Validate(data []byte) (result bool) {
	return validate(data)
}

// Open opens book to reader
func (b *Book) Open(file []byte) error {
	return open(file, b)
}

// Get returns specified file from book
func (b *Book) Get(href string) (result []byte, err error) {
	switch b.Extension {
	case "epub":
		book := b.entity.(*epub.Book)
		return getters.GetByHref(book, href)
	default:
		return nil, errors.New("unsupported type of book to get")
	}
}
