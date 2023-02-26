package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// creating a struct called person
type person struct {
	firstName string
	lastName  string
	// contact   contactInfo // here we are adding another field called contact whose type will be contactInfo i.e a struct
	contactInfo // this is equivalent of using contact contactInfo
}

func main() {
	//1st way to create a struct
	// alex := person{"Alex", "Anderson"} // assigning and initializing a variable called alex with a person struct here the order of assignment is what matters

	//2nd way
	// alex := person{lastName: "Anderson", firstName: "Alex"} // here we can change the order of the fields and still the struct will be defined with the correct order
	// fmt.Println(alex)

	// 3rd way
	// var alex person // we are just initializing alex as a variable of type person here that's it we are not adding the firstName and lastName of alex here
	// // so the zero value will be assigned to both firstName and lastName the zero value for a string is an empty string ("")
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex) //"%+v" will print out all the field names and their values  of the variable alex

	// jim := person{
	// 	firstName: "Jim",
	// 	lastName:  "Party",
	// 	contact: contactInfo{
	// 		zipCode: 94000,
	// 		email:   "JimPartiesHard@gmail.com",
	// 	},
	// }

	// fmt.Printf("%+v", jim)
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{ // if the fieldName for contactInfo type in the person struct is not declared then the fieldName is by default the name of the type
			zipCode: 94000,
			email:   "JimPartiesHard@gmail.com",
		},
	}

	// look at this variable and give us access to the memory address that this variable is pointing at variable in this case is jim
	// jim is pointing at the person struct in the memoory and that struct is existing at some particular acess
	// &jim we say give me access to the struct jim is pointing at
	// jimPointer := &jim // jimPointer will now point to the same address that jim is at
	// jimPointer.updateName("Dwight")

	// if we define a receiver with a type of pointer to whatever type and when we call
	// the method on top of that go will allow us to call this method with either the pointer or the normal variable
	jim.updateName("Dwight")
	jim.print()
}

// But before we do, I want to give you a quick refresher on how RAM or how memory works on your machine, on your computer right now.

// Now, as we're talking about Ram, I'm going to give you a somewhat non-scientific definition.

// So I'm going to give you just kind of a very sweeping, broad strokes overview on how RAM works, because

// I want you to think more about what Go is doing right now rather than doing a 20 minute dive on how Ram works on your local machine.

// So all I really want you to understand about Ram right now is that memory on your local machine can

// be thought of as like a bunch of little cubbies or a bunch of little slots or a bunch of little boxes.

// Each box in your computer's memory can store some data, and each one of these little boxes or these little value containers has some discrete address.

// And so whenever your program says, Oh, I want to retrieve some information from the computer's memory,

// it looks at it goes and finds some address, and then it pulls the value out of there.

// And so each of these little boxes right here can contain some amount of information.

// And that's really all I want to talk about as far as RAM goes right now.

// Just quick overview on exactly how it works.

// In the last section we discussed how go is a pass by value language.

// That means that whenever we pass a value into a function, go copies that value and then the copy is

// made available to the code inside the function.

// So the net result to that meant that whenever we pass Jim into the update name function go first, made

// a copy of that struct and then the copy was made available to the update name function.

// And so whenever we updated the person inside of update name, the change was never propagated back to Jim.

//	func (p person) updateName(newFirstName string) {
//		p.firstName = newFirstName
//	}

// *pointerToPerson says give me acess to the value sitting at the pointerToPerson memory location in this case pointerToPerson is jimPointer
// *person is a pointer that points to person
// ********************************************************************** watch pointer operations again tomorrow morning pointer operations is tricky ***************************************
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// we setup a function with a person type receiver so now we can call the print() method on the person type variables tada
func (p person) print() {
	fmt.Printf("%+v", p)
}
