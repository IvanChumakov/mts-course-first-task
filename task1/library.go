package main

type Generator func(book *Book)

type Library struct {
	storage Storage
	mapping map[string]int
}

func NewLibrary(storage Storage) *Library {
	mapping := make(map[string]int)
	return &Library{
		storage: storage,
		mapping: mapping,
	}
}

func (l *Library) ChangeStorage(newStorage Storage) {
	l.storage = newStorage
	l.mapping = make(map[string]int)
}

func FirstIdGenerator(book *Book) {
	countRunes := func(str string) int {
		sum := 0
		for i := range str {
			sum += i
		}
		return sum
	}
	book.SetId(countRunes(book.GetName()) + countRunes(book.GetAuthor()))
}

func SecondIdGenerator(book *Book) {
	countRunes := func(str string) int {
		sum := 0
		for i := range str {
			sum += i
		}
		return sum
	}
	book.SetId(countRunes(book.GetName()) + countRunes(book.GetAuthor()) + (book.GetPageNum() % 883))
}

func (l *Library) AddBook(generator Generator, book *Book) {
	generator(book)
	l.mapping[book.GetName()] = book.GetId()
	l.storage.AddBook(book)
}

func (l *Library) GetBook(name string) Book {
	return l.storage.GetBook(l.mapping[name])
}
