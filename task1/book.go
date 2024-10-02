package main

type Book struct {
	pageNum int
	author  string
	name    string
	id      int
}

func NewBook(pageNum int, author, name string) *Book {
	return &Book{
		author:  author,
		name:    name,
		pageNum: pageNum,
	}
}

func (b *Book) GetAuthor() string {
	return b.author
}

func (b *Book) GetName() string {
	return b.name
}

func (b *Book) GetPageNum() int {
	return b.pageNum
}

func (b *Book) SetId(id int) {
	b.id = id
}

func (b *Book) GetId() int {
	return b.id
}
