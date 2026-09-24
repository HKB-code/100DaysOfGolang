package shelf

import (
	"errors"

	book "example.com/library_management/Book"
)

type Shelf struct {
	Books []*book.Book
	nextId int
	available bool
	ShelfId int
}

func New()*Shelf{
	return  &Shelf{
		Books: []*book.Book{},
		nextId: 0,
		available: true,
		ShelfId: 0,
	}
}

func (s Shelf)AddBook(genre,title,author string,pages,id int)error{

	for i := range s.Books{
		if s.Books[i].Author==author && s.Books[i].Title==title{
			return  errors.New("book already existed")
		}
	}
	s.available = true
	s.nextId++
	b1,err:=book.New(genre,author,title,pages,s.nextId,s.available)
	if  err!= nil {
		return  err
	}
	s.Books = append(s.Books, b1)
	return  nil

}



func (s Shelf)RemoveBook(author,title string)error{
for i := range s.Books{
	if s.Books[i].Author==author&&s.Books[i].Title==title{
		s.Books = append(s.Books[:i],s.Books[i+1:]... )
		return nil
	}
}

	return  errors.New("book not found")
}


func (s Shelf)FindBook(author,title string)(*book.Book,error){
for _ ,v:= range s.Books{
	if v.Author==author&&v.Title==title{
		
		return v,nil
	}
}

	return   nil,errors.New("couldn't find the book")
}