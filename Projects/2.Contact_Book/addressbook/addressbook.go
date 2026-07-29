package addressbook

import (
	"encoding/json"
	"errors"
	"os"
	"sort"

	"example.com/contact_book/address"
	"example.com/contact_book/person"
)


type Addressbook struct{
	Contacts [] *person.Person
	nextId int
	filePath string
}


func New()*Addressbook{
	return  &Addressbook{
		Contacts: []*person.Person{},
		nextId: 0,
		filePath: "contact.json",
	}
}

func (ad *Addressbook)Add(name,email,phone,street,city,state,country,zipcode string, )error{
for _,v:= range ad.Contacts{
	if v.Phone == phone && v.Email==email{
		return  errors.New("contact already existed")
	}
}
	add,err:=address.New(street,city,state,country,zipcode)
	if  err!= nil {
		return  err
	}


	p,err:=person.New(name,email,phone,add,ad.nextId)
	if  err!= nil {
		return  err
	}
	ad.Contacts = append(ad.Contacts, p)
	
ad.nextId++
	
	return  nil
}

func (ad *Addressbook) Delete(id int)error{
	for i,v := range ad.Contacts{
	if	v.Id == id{
		ad.Contacts = append(ad.Contacts[:i],ad.Contacts[i+1:]... )
		return  nil
	}
	}

	return  errors.New("contact not found")
}

func (ad *Addressbook)Update(id int, phone,email string,)error{
	
	
	for i := range ad.Contacts{
		if ad.Contacts[i].Id == id{
			ad.Contacts[i].Phone = phone
			ad.Contacts[i].Email = email
			
			return  nil
		}
		
	}
	return   errors.New("conatct not found")
}

func (ad *Addressbook)Search(name,phone string)(*person.Person,error){
for _,v := range ad.Contacts{
	if v.Name == name &&v.Phone==phone{
		
		return  (v),nil
	}
}
	return  nil,errors.New("conatct not found")
}

func (ad *Addressbook)Sort()[]*person.Person{
	copyContacts := make([]*person.Person,len(ad.Contacts))
	copy(copyContacts,ad.Contacts)

	sort.Slice(copyContacts,func(i, j int) bool {
		return copyContacts[i].Name < copyContacts[j].Name
	})
return copyContacts
}




func (ad *Addressbook)SaveToJson()error{
	
	file,err:= os.Create(ad.filePath)
	if  err!= nil {
		return  err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent(""," ")
	err = encoder.Encode(ad.Contacts)
	if  err!= nil {
		return err
	}

	return  nil
}




func (ad *Addressbook)LoadData()(error){
	file,err:= os.Open(ad.filePath)
	if  err!= nil {
		if os.IsNotExist(err){
			ad.Contacts= []*person.Person{}
			ad.nextId = 1
			return  nil
		}
		return err
	}
	defer file.Close()
	
	err= json.NewDecoder(file).Decode(&ad.Contacts)
if  err!= nil {
	return err
}
ad.nextId=1
  for _,p := range ad.Contacts{
	if p.Id>=ad.nextId{
		ad.nextId = p.Id+1
	}
  }
return nil
}

