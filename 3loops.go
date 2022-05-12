//Go language has just one single loop that is for loop
//we use for loop in various forms to match the usage of other loops 


package main

import "fmt"


func main(){
	//using for loop as simple for loop like c , c++ and java

	for i:=0;i<10;i++{
		fmt.Println("Hello shubham")
	}


	//using for loop as while loop

	j:=0
	for j<10{
		fmt.Println("Hello shubham")
	}

	//for loop as infinite loop like while(True) in c , c++ and java
	for{
		fmt.Println("One day you are going to be successfull")
	}

	//using for loop like range in python
	
	/*
	
	Syntax:

		for i, j:= range rvariable{
		// statement..
		}
		Here,

		i and j are the variables in which the values of the iteration are assigned. They are also known as iteration variables.
		The second variable, i.e, j is optional.
		The range expression is evaluated once before the starting of the loop.
			
	
	
	*/



	rangeVar:=[]string{"SHubham","Pranit","Sanu","Navnit"}
	for i , j := range rangeVar{
		fmt.Println(i , j)
	}


	
}