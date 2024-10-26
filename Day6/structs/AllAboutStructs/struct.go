package main

import (
	"encoding/json"
	"fmt"
)

/*Structs in Go
Structs are one of the most important features in Go, allowing you to create complex data structures. Let's dive into the details of structs in Go:

Definition:
A struct is a composite data type that groups together zero or more fields with different data types. It's used to represent a record.

syntax:
type StructName struct {
    field1 type1
    field2 type2
    // more fields...
}

*/

//Define a struct named Person
type Person struct{
	name string
	age int
	email string 
}
/*Embedded Structs (Composition):
Go supports composition through embedding, allowing one struct to include another struct.*/
type Employee struct{
	Person
	EmployeeId string
}

/*Tags in Structs:
Tags provide metadata for struct fields, often used for JSON encoding/decoding.*/
type Alien struct{
	Planet string `json:"Planet"`
	Powers string    `json:"Powers"`
	FromSpace bool    `json:"FromSpace"`
} 


func main(){
	//creating an instance of the struct
	p:=Person{
		name: "Jhon Doe",
		age: 30,
		email: "jhon.doe@example.com",
	}



//Accessing fields
fmt.Println("name: ",p.name)
fmt.Println("age: ",p.age)
fmt.Println("email: ",p.email)

//Modifying fields
p.age = 31
fmt.Println("Updated Age: ", p.age)

//method call
fmt.Println("Greet: ",p.greet())


//Example of Embedding:
	emp:=Employee{
	Person: Person{
		name: "Jhon Doe",
		age: 30,
		email: "jhon.doe@example.com",
},
EmployeeId: "E1234",
	}

	/// Accessing fields of embedded struct
    fmt.Println("Name:", emp.name)
    fmt.Println("Age:", emp.age)
    fmt.Println("Employee ID:", emp.EmployeeId)


		//Tags: Add metadata for fields.
x:=Alien{
	Planet: "Mars",
	Powers:"Leaser",
	FromSpace: true,
}
jsonData,_ := json.Marshal(x)
fmt.Println(string(jsonData))

}
//Methods
func (p Person) greet() string{
	return "Hello my name is "+p.name
}



/*Instance in Programming:
An instance is a concrete occurrence of any object or data structure. When you define a struct in Go (or a class in OOP languages), you're creating a blueprint. An instance is an actual object created from this blueprint, with its own unique values.

Creating Instances in Go:
Even though Go is not purely object-oriented, it allows you to create instances of structs, which are similar to creating instances of classes in OOP languages

Explanation:
Struct Definition: type Person struct {...} defines the blueprint for a Person.

Creating an Instance: p := Person{...} creates an instance of the Person struct. This instance p has its own values for Name, Age, and Email.

Summary:
Instance: A specific realization of a data structure (like a struct in Go).

Creating Instances in Go: Use the blueprint defined by a struct to create objects with unique values.




Note:In Go, structs can be seen as a structure or blueprint that helps in creating instances (or samples), just like classes work in object-oriented programming (OOP) languages.

Struct as a Blueprint:
Struct: A blueprint that groups multiple fields together.

Instance: When you use this blueprint (struct) to create an actual object, it is called an instance.

Summary:
In Go, structs can be considered as an architecture or structure that helps in creating objects, similar to classes in OOP languages.

You can create instances of structs in Go and use them just like you use objects of classes in OOP languages.
*/