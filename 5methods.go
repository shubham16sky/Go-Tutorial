package main

import "fmt"

type students struct{
	name string
	age int
	branch string
}



//method with a receiver of students type

func (s students)show(){
	fmt.Println("Details of the students are :\n")
	fmt.Println("Name is : ",s.name)
	fmt.Println("Age is : ",s.age)
	fmt.Println("Branch is : ",s.branch)
}



func main(){

//Initializing the value of student structure	
	stud := students{
		name : "Shubham",
		age : 20,
		branch : "ISE",
	}

//calling the show method
	stud.show()


}
