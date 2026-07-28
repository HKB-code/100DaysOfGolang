package manager

import (
	"errors"
	"fmt"

	"example.com/student_managment_cli/student"
)
type st = student.Student

type  Manager struct{
	students []st
	nextId int
}

func New()*Manager{
	return  &Manager{
		students: []st{},
		nextId: 0,
	}
}






func (m *Manager)Add(name,grade string, age int )error{
		for _,student := range m.students{
		if student.Name== name && student.Age== age{
			
			return  errors.New("student already existed")
		}
	}
	m.nextId++

	s,err:=student.New(m.nextId,name,age,grade)
	if  err!= nil {
		return  err
	}

	m.students = append(m.students, s)
	return  nil
}

func (m *Manager)Delete(id int)error{
	
	for i,v := range m.students{
		if v.Id ==id{
			m.students = append(m.students[:i],m.students[i+1:]... )
			return  nil
		}
		
	}
	
	return  errors.New("student not found")
}


func (m *Manager) Search(id int)(st,error){

	for _,v := range m.students{
		if v.Id == id{
			return  v,nil
		}
	}
	return st{},errors.New("student not found")
}

func (m *Manager) Update(id int,age int,grade string)error{

	for i := range m.students{
		if m.students[i].Id==id{
			m.students[i].Age = age
			m.students[i].Grade = grade
			return nil
		}
		
}
return  errors.New("student not found")

}

func (m *Manager)List(){
	if len(m.students)==0{
		fmt.Println("No student Found")
		return 
	}
	for _,v := range m.students{
		fmt.Printf("ID: %d | Student: %s | Age: %d | Grade: %s\n",v.Id,v.Name,v.Age,v.Grade)
	}
}


