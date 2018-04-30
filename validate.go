package reader

import (
	"gitlab.com/alexnikita/gols/filetype"
	"gitlab.com/alexnikita/gols/filetype/types"
)

func validate(data []byte) bool {
	typ := filetype.Detect(data)

	switch typ {
	case types.Unknown:
		return false
	case types.EPUB:
		return true
	default:
		return false
	}
}
