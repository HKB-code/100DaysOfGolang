package library

import (
	"errors"
	"fmt"

	"example.com/library_management/shelf"
)

type Library struct {
	Shelfs []*shelf.Shelf
	sn int
}

func (lb Library)AddShelf()error{
	s:= shelf.New()
	lb.sn++
	s.ShelfId = lb.sn
	
	
	lb.Shelfs = append(lb.Shelfs, s)
	
	return  nil
}


func (lb Library)FindShelf(id int)(*shelf.Shelf,error){
for _,v := range lb.Shelfs{
	if v.ShelfId == id{
        return  v,nil
	}
}
return  nil,errors.New("can not find shelf")
}

func(lb Library)BorrowBook(shelfId int, title,author string)(error){
for i := range lb.Shelfs{
	if lb.Shelfs[i].ShelfId == shelfId{
		for j:= range lb.Shelfs[i].Books{
            if lb.Shelfs[i].Books[j].Author== author && lb.Shelfs[i].Books[j].Title==title{
				y,str:=  lb.Shelfs[i].Books[j].Borrow()		
				if y{
					fmt.Println(str)
					return nil
				}
			}
		}
	}
}
	return  errors.New("not available")
}



func(lb Library)ReturnBook(shelfid int, author,title string)(error){
for i := range lb.Shelfs{
	if lb.Shelfs[i].ShelfId == shelfid{
		for j:= range lb.Shelfs[i].Books{
            if lb.Shelfs[i].Books[j].Author== author && lb.Shelfs[i].Books[j].Title==title{
			return 	lb.Shelfs[i].Books[j].Return(author,title)
			}
		}
	}
}
	return  errors.New("not available")
}
