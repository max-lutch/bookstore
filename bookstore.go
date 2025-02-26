package bookstore

// Book represents information about a book.
type Book struct {
	Title  string
	Author string
	Copies int
	ID     int
}

type Catalog map[int]Book

func (c Catalog) GetAllBooks() []Book {
	result := []Book{}
	for _, b := range c {
		result = append(result, b)
	}
	return result
}
