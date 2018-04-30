package parsers_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/alexnikita/gols/epub"
	"gitlab.com/alexnikita/treader/reader/parsers"
)

func TestParseEPUB(t *testing.T) {
	var (
		err      error
		filename = "../test_books/1.epub"
	)

	_, err = epub.Open(filename)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetEPUBInfo(t *testing.T) {
	var (
		err      error
		filename = "../test_books/1.epub"
	)

	file, err := os.Open(filename)
	if err != nil {
		t.Fatal(err)
	}

	info, err := parsers.GetEPUBInfo(file)
	if err != nil {
		t.Fatal(err)
	}

	file.Close()

	assert.Equal(t, true, info.Coverpage != "")
	assert.Equal(t, true, info.SpineStack.Value() != "", "spine stack must not be empty")
}
