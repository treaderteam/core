package getters

import "gitlab.com/alexnikita/gols/epub"

// GetByHref returns specified html page from book
func GetByHref(book *epub.Book, href string) (file []byte, err error) {
	return book.Render(href)
}
