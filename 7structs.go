/*
A structure of struct in golang is a user - defined type that allows to group / combine
items of possibely different types into a single type.
Any real world object entity which has some set of
properties can be repreented as struct.
This concept is generally compared with classes in object oriented programming.
It can be termed as a lightweight class that does nit support inheritance
but supports composition.
*/

/*
For example let's an Address has name , street , city , state , Pincode.
It makes sense to group these three properties into a single structure
Address as shown below



	type Address struct {
		name string
		street string
		city string
		state string
		pincode string
	}

	In the above the type keyword introduces a new type.
	It is followed by the name of the type that is Address
	and the keyword struct to illustrate that we are defining a struct.
	The struct contains a list of various fields inside the curly braces.
	Each field has a name and a type.


	//Initializing a variable of struct type

	var add = Address{"Pranit","Yalhanka","Bangalore","karntaka","56064"}
*/

package main

import "fmt"

type Address struct {
	name    string
	street  string
	city    string
	state   string
	pincode int
}

func main() {
	var add Address

	fmt.Println(add)

	add = Address{"Pranit", "yalhanka", "Bangalore", "Karnataka", 560064}
	fmt.Println(add)

	//Accessing individual field of a struct
	//Dot (.) operator can be used to access the single fields

	fmt.Println(add.name)
}
