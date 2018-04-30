package reader

import (
	"bytes"

	"gitlab.com/alexnikita/treader/reader/parsers"
)

func parseEpub(data []byte) (result Book, err error) {

	book, entity, err := parsers.GetEPUBInfo(bytes.NewReader(data))
	if err != nil {
		return result, err
	}

	result = Book{
		Author:     book.Author,
		Title:      book.Title,
		Coverpage:  book.Coverpage,
		SpineStack: book.SpineStack,
		Extension:  book.Extension,
		entity:     entity,
	}

	return result, nil
}
