package loadData

import (
	"encoding/json"
	"errors"
	"os"

	"example.com/contact_book/addressbook"
	"example.com/contact_book/person"
)
const File = "contact.json"

func SaveToJson()error{
	var users addressbook.Addressbook
	file,err:= os.Create(File)
	if  err!= nil {
		return  err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent(""," ")
	err = encoder.Encode(users)
	if  err!= nil {
		return err
	}

	return  nil
}



func UnloadData()([]*person.Person,error){
	file,err:= os.Open(File)
	if  err!= nil {
		return nil,errors.New("file couldnt found")
	}
	defer file.Close()
	var users []*person.Person
	err= json.NewDecoder(file).Decode(&users)
if  err!= nil {
	return nil,errors.New("failed to decode json file")
}
return users,nil
}


