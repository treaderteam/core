package reader

import (
	"bytes"
	"io"
	"log"
	"os"

	"gitlab.com/alexnikita/gols/epub"
	"gitlab.com/alexnikita/gols/filetype"
	"gitlab.com/alexnikita/gols/filetype/types"
)

const (
	TYPE_EPUB = "epub"
)

// Validate checks if file can be readed
func Validate(file []byte) (result bool) {
	typ := filetype.Detect(file)

	switch typ {
	case types.Unknown:
		return false
	case types.EPUB:
		return true
	default:
		return false
	}
}

type Book struct {
	Title string
}

// Open opens book to reader
func (b *Book) Open(file []byte) {
	typ := filetype.Detect(file)

	switch typ {
	case types.EPUB:
		book, err := openEpub(file)
		if err != nil {
			log.Println(err)
		}

		(*b).Title = book.Opf.Metadata.Title[0]
		break
	default:
		break
	}
}

// Test func
func Test() string {
	return "hello"
}

func openEpub(data []byte) (result epub.Book, err error) {

	tmpfilename := os.TempDir() + "/tmpepub"

	file, err := os.OpenFile(tmpfilename, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return result, err
	}

	defer file.Close()

	_, err = io.Copy(file, bytes.NewReader(data))
	if err != nil {
		return result, err
	}

	book, err := epub.Open(tmpfilename)
	if err != nil {
		return result, err
	}

	return *book, nil
}
