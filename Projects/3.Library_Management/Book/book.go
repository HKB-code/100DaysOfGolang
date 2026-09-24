package book

import "errors"

type Book struct {
	Genre     string
	Available bool
	Author    string
	Pages     int
	Title     string
	Id        int
}
func New(genre,author,title string, pages,id int,available bool)(*Book,error){
	if (genre==""||author==""||title==""||pages==0||id==0){
		return  nil,errors.New("please enter all field")
	}
	return  &Book{
		Genre: genre,
		Available: available,
		Author: author,
		Pages: pages,
		Title: title,
		Id: id,
		
	},nil
}

func (b *Book) Borrow() (bool, string) {
	if b.Available == true {
		return true, "yes this available"
	}
	return false, "sorry it is not available"
}

func (b *Book) Return(author,title string) error {
	if b.Author == author|| b.Title==title {
		b.Available = true
		return nil
	}
	return errors.New("book not found")
}