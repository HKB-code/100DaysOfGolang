package student

import "errors"

type Student struct {
	Id    int
	Name  string
	Age   int
	Grade string
}

func New(id int, name string, age int, grade string) (Student, error) {
	if id == 0 || name == "" || age <= 0 || grade == "" {
		return Student{}, errors.New("please enter all the things")
	}
	return Student{
		Id: id,
		Name: name,
		Age: age,
		Grade: grade,
	}, nil
}