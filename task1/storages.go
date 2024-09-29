package main

type Storage interface {
	GetBook(id int) Book
	AddBook(book *Book)
}

type SliceStorage struct {
	sliceStorage []*Book
}

func NewSliceStorage() *SliceStorage {
	storage := make([]*Book, 0)
	return &SliceStorage{
		sliceStorage: storage,
	}
}

func (s *SliceStorage) GetBook(id int) Book {
	for _, book := range s.sliceStorage {
		if book.id == id {
			return *book
		}
	}
	return Book{}
}

func (s *SliceStorage) AddBook(book *Book) {
	s.sliceStorage = append(s.sliceStorage, book)
}

type MapStorage struct {
	mapStorage map[int]*Book
}

func NewMapStorage() *MapStorage {
	storage := make(map[int]*Book)
	return &MapStorage{
		mapStorage: storage,
	}
}

func (m *MapStorage) GetBook(id int) Book {
	if value, ok := m.mapStorage[id]; ok {
		return *value
	}
	return Book{}
}

func (m *MapStorage) AddBook(book *Book) {
	m.mapStorage[book.id] = book
}
