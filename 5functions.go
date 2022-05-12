/*

func function_name(Parameter-list)(Return_type){
    // function body.....
}


here func is keyword , which is used to create functions.
function_name : - it is name of your function
parameter-list : - It contains the name and the type of the function parameters.

Return_type: It is optional and it contain the types of the values that function returns.
			 If you are using return_type in your function, then it is necessary to use a return statement in your function

*/

package main

import (
	"fmt"
)

func main() {

	var a, b int
	fmt.Println("Enter A and B : ")
	fmt.Scanf("%v%v", &a, &b)

	fmt.Println("Sum of A and B is : ", add(a, b))

}

func add(a, b int) int {
	sum := a + b

	return sum




/*

There are two types of passsing argument to a functpion that is :

1. Pass by value 
2. Pass by reference

Call by value: : In this parameter passing method,
 	values of actual parameters are copied to function’s formal parameters and
  	the two types of parameters are stored in different memory locations.
	So any changes made inside functions are not reflected in actual parameters of the caller.


Call by reference: Both the actual and formal parameters refer to the same locations,
	so any changes made inside the function are actually reflected in actual parameters of the caller.




*/


