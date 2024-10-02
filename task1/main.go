package main

func main() {
	bookSlice := make([]*Book, 0)
	bookSlice = append(bookSlice, NewBook(532, "Сьюзен Коллинз", "Голодные игры"))
	bookSlice = append(bookSlice, NewBook(700, "Фенимор Купер", "Зверобой"))
	bookSlice = append(bookSlice, NewBook(670, "Дж.К Роулинг", "Гарри Потер и Философский камень"))

	storage := NewSliceStorage()
	library := NewLibrary(storage)

	for _, book := range bookSlice {
		library.AddBook(FirstIdGenerator, book)
	}

	firstBook := library.GetBook("Голодные игры")
	println(firstBook.GetAuthor())

	secondBook := library.GetBook("Зверобой")
	println(secondBook.GetAuthor())

	library.AddBook(SecondIdGenerator, NewBook(301, "Александр Куприн", "Гранатовый браслет"))
	thirdBook := library.GetBook("Гранатовый браслет")
	println(thirdBook.GetAuthor())

	mapStorage := NewMapStorage()
	library.ChangeStorage(mapStorage)

	library.AddBook(SecondIdGenerator, NewBook(340, "Этель Лилиан Войнич", "Овод"))
	library.AddBook(SecondIdGenerator, NewBook(700, "Фенимор Купер", "Зверобой"))
	library.AddBook(SecondIdGenerator, NewBook(333, "Редьярд Киплинг", "Сталки и компания"))

	fourthBook := library.GetBook("Овод")
	println(fourthBook.GetAuthor())

	lastBook := library.GetBook("Сталки и компания")
	println(lastBook.GetAuthor())
}
