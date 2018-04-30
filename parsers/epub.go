package parsers

import (
	"bytes"
	"encoding/base64"
	"io"
	"log"
	"os"
	"strings"

	"gitlab.com/alexnikita/gols/epub"
	"gitlab.com/alexnikita/gols/util"
)

// GetEPUBInfo implementation
func getEPUBInfo(rdr io.Reader) (result BasicBookInfo, err error) {
	book, err := parseEPUBFromReader(rdr)
	if err != nil {
		return
	}

	author := ""
	title := ""
	coverpage := ""

	if len(book.Opf.Metadata.Creator) > 0 {
		author = book.Opf.Metadata.Creator[0].Data
	}

	if len(book.Opf.Metadata.Title) > 0 {
		title = book.Opf.Metadata.Title[0]
	}

	if len(book.Opf.Metadata.Meta) > 0 {
		coverid := ""
		covername := ""
		for _, v := range book.Opf.Metadata.Meta {
			if v.Name == "cover" {
				coverid = v.Content
				break
			}
		}

		if coverid != "" && len(book.Opf.Manifest) > 0 {
			for _, v := range book.Opf.Manifest {
				if v.ID == coverid {
					covername = v.Href
					break
				}
			}
		}

		if covername != "" {
			coverdata, err := book.GetFile(strings.Replace(covername, "OEBPS/", "", 1))
			if err == nil {
				coverpage = base64.StdEncoding.EncodeToString(coverdata)
			} else {
				log.Println(err)
			}
		}
	}

	words, err := book.GetWords()
	if err != nil {
		log.Println(err)
		words = make(map[string]string)
	}

	result = BasicBookInfo{
		Author:    author,
		Title:     title,
		Coverpage: coverpage,
		Words:     words,
		Extension: "epub",
	}

	return
}

// parseEPUBFromReader read and parse epub from reader
func parseEPUBFromReader(rdr io.Reader) (book *epub.Book, err error) {
	var (
		tmpFilename = util.RandStringBytesRmndr(24)
		file        *os.File
	)

	if file, err = os.Create(tmpFilename); err != nil {
		return
	}

	if _, err = io.Copy(file, rdr); err != nil {
		return
	}

	if book, err = epub.Open(tmpFilename); err != nil {
		return
	}

	file.Close()
	os.Remove(tmpFilename)

	return
}

// ParseEPUB read and parse epub
func ParseEPUB(data []byte) (book *epub.Book, err error) {
	var (
		tmpFilename = util.RandStringBytesRmndr(24)
		file        *os.File
	)

	if file, err = os.Create(tmpFilename); err != nil {
		return
	}

	defer file.Close()

	defer os.Remove(tmpFilename)

	reader := bytes.NewReader(data)

	if _, err = io.Copy(file, reader); err != nil {
		return
	}

	if book, err = epub.Open(tmpFilename); err != nil {
		return
	}

	return
}
