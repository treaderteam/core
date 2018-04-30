package reader

import (
	"errors"

	"gitlab.com/alexnikita/gols/filetype"
	"gitlab.com/alexnikita/gols/filetype/types"
)

func open(data []byte, b *Book) error {
	typ := filetype.Detect(data)

	switch typ {
	case types.EPUB:
		book, err := parseEpub(data)
		if err != nil {
			return errors.New("cannot parse epub")
		}

		*b = book
		return nil
	default:
		return errors.New("unsupported book type")
	}
}
