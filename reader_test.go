package reader_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/alexnikita/treader/reader"
)

func TestReader(t *testing.T) {
	filename := "./test_books/1.epub"

	file, err := os.Open(filename)
	if err != nil {
		t.Fatal(err)
	}

	defer file.Close()

	ent, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}

	book := new(reader.Book)

	err = book.Open(ent)
	if err != nil {
		t.Fatal(err)
	}

	first := book.SpineValue()

	result, err := book.Get(first)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, true, result != nil, "result must not be empty")
}
