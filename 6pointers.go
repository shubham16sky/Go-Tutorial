/*

Pointers in Go is a variable that is used to store the memory address of another variable.
Pointer in golang is aso termed as special variable.
A pointer is a special kind of variable that is not only used to store the memory address of other
variables but also points where the memory is located and porvides the ways
to find out the value stored at the memory location.


It is generally termed as a special kind of variable because it is
almost declared as a variable but with *(dereferencing operator).




--> * operator is also termed as dereferencing operator
	used to declare pointer variable and access the value stored in the address.

--> & operator is termed as address operator used to return the address of a variable
	or to access the address the address of a variable to a pointer.




//Declaring a pointer

var s *string


here var is a keyword to declare variables
s is name of pointer / variable
string is datatype.



Initializing a Pointer : To do so we need to initialize a pointer with the memory address of
another variable using the address operator.

Ex :-

//Normal variable declaration
var a = 45

//Initialization of pointer s with memory address of a


var s *int = &a

*/
package main

import "fmt"

func main() {
	var x int = 5656

	var p *int

	p = &x

	fmt.Println("value stored in x = ", x)
	fmt.Println("Address of x = ", &x)
	fmt.Println("Value of stored in variable p = ", p)

	/*

		//Dereferencing the pointer
		As we know that * operator is also termed as the dereferencing operator. It is not only used to
		declare the pointer variable but also used to access the value stored in the variable which
		the pointer points to which is generally termed as indirecting or dereferencing.
		* operator is also termed as the value as the address of.

	*/

	var y = 420
	var z = &y

	fmt.Println("Value stored in y : ", y)
	//now we will access the value stored at the pointing location
	//using the deeferencing variable

	fmt.Println("value stored in z(*y) = ", *z)

	//You can also change the value of the pointer or at the memory location
	//instead of assigning a new value to the variable.

	*z = 500

	fmt.Println("Value stored after changing :", y)
}
