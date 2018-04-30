package reader

const (
	TYPE_EPUB = "epub"
)

// Book is primary book type
type Book struct {
	Author    string
	Title     string
	Coverpage string
}

// Validate checks if file can be readed
func Validate(data []byte) (result bool) {
	return validate(data)
}

// Open opens book to reader
func (b *Book) Open(file []byte) error {
	return open(file, b)
}

// Test func
func Test() string {
	return "hello"
}
