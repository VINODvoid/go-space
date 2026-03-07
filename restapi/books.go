package main

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}
type Store struct {
	Books   []Book `json:"books"`
	Counter int    `json:"counter"`
}

func (store *Store) Add(title string, author string, year int) Book {
	id := len(store.Books)
	book := Book{
		ID:     id + 1,
		Title:  title,
		Author: author,
		Year:   year,
	}
	store.Books = append(store.Books, book)
	return book
}

func (store *Store) GetAll() []Book {
	return store.Books
}

func (store *Store) GetByID(id int) (Book, bool) {
	for _, book := range store.Books {
		if book.ID == id {
			return book, true
		}
	}
	return Book{}, false
}

func (store *Store) Delete(id int) bool {
	for index, book := range store.Books {
		if book.ID == id {
			store.Books = append(store.Books[:index], store.Books[index+1:]...)
			return true
		}
	}
	return false
}
