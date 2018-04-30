package parsers

import (
	"io"

	"gitlab.com/alexnikita/gols/epub"
)

// BasicBookInfo book info
type BasicBookInfo struct {
	Author    string
	Title     string
	Coverpage string
	Extension string
	Words     map[string]string
	SpineStack
}

// GetEPUBInfo return basic book info from epub
func GetEPUBInfo(rdr io.Reader) (result BasicBookInfo, book *epub.Book, err error) {
	return getEPUBInfo(rdr)
}
