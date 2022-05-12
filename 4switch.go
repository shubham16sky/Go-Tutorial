/*
There are two types of switch statements in Go :
		1.Expression switch
		2. Type switch

Here we will see epression switch

Expression switch is similar to switch statement in C , C++ and Java


Syntax :

		switch optstatement; optexpression{
		case expression1: Statement..
		case expression2: Statement..
		...
		default: Statement..
		}


Important points :

--> Both optstatement and optexpression in the expression switch are optional statements.
--> If both optstatementand optexpression are present, then a semi-colon(;) is required in between them.
--> If the switch does not contain any expression, then the compiler assume that the expression is true.
--> The optional statement, i.e, optstatement contains simple statements like variable declarations, increment or assignment statements, or function calls, etc.
--> If a variable present in the optional statement, then the scope of the variable is limited to that switch statement.
--> In switch statement, the case and default statement does not contain any break statement.
	 But you are allowed to use break and fallthrough statement if your program required.
--> The default statement is optional in switch statement.
--> If a case can contain multiple values and these values are separated by comma(,).
--> If a case does not contain any expression, then the compiler assume that te expression is true.
*/

package main

import "fmt"

func main() {

	//switch statement with both optional statement and expression
	switch day := 4; day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid")
	}



	var val int = 2
      
    // Switch statement without an     
    // optional statement and 
    // expression
   switch {
       case val == 1:
       fmt.Println("Hello")
       case val == 2:
       fmt.Println("Bonjour")
       case val == 3:
       fmt.Println("Namstay")
       default: 
       fmt.Println("Invalid")
   }


   var value string = "five"
      
    // Switch statement without default statement
    // Multiple values in case statement
   switch value {
       case "one":
       fmt.Println("C#")
       case "two", "three":
       fmt.Println("Go")
       case "four", "five", "six":
       fmt.Println("Java")
   }  
}
