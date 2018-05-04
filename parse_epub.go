package reader

import (
	"bytes"

	"gitlab.com/alexnikita/gols/util"
	"gitlab.com/alexnikita/treader/reader/parsers"
)

func parseEpub(data []byte) (result Book, err error) {

	hash, err := util.Hash(data)
	if err != nil {
		return Book{}, err
	}

	book, entity, err := parsers.GetEPUBInfo(bytes.NewReader(data))
	if err != nil {
		return result, err
	}

	result = Book{
		Author:    book.Author,
		Title:     book.Title,
		Coverpage: book.Coverpage,
		Extension: book.Extension,
		entity:    entity,
		spine:     book.SpineStack,
		Hash:      hash,
	}

	return result, nil
}
