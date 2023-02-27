//how do we run the code in our project..??
//go run <fileName> which is main.go  go run complies and executes one or two files.
// go buid will compile a bunch of source code files
// run the executable file on mac by doing ./<fileName>

// what does package main mean..??
// package main specifically is called to build an executable package
// if we name it anything else then we wouldn't have an executable package
// multiple files can have the package main and then all those files will be packaged into an executable file

// basically whenever we see the word package main, that means we are making an executable package.
// Any other name whatsoever means we are making a reusable or dependency type package.
// Now, the last thing I want to tell you about this last thing is that any time we make an executable
// package, it must always have a function inside of it called Main as well.
// So if we flip back over to our code editor, that's what this statement right here was.
// We had said func main.
// So we specifically created a function called Main because we had called package main right here, creating an Executable type package.
package main

// what does import "fmt" mean..??
// So saying import FMT specifically means give my package main access to all of the code and all the functionality

// that is contained inside of this other package called Fmt.

// Fmt is the name of a standard library package that is included with the go programming language by default.

// Fmt itself is a kind of a shortened form of the word format.

// The Fmt library is used to print out a lot of different information specifically to the terminal, just

// to give you a better sense of debugging and stuff like that.
import "fmt"

// what does this func thing mean
// func is short for functions
func main() {
	fmt.Println("Hi There!")
}

// How is main.go file organized..??

// So it's always going to be the exact same pattern inside of every single file that we ever create at the very top.

// We're always going to place our package declaration.

// So remember we say, Oh, this file is a part of package, blah, blah, blah. In this case, package main.

// Then right underneath that, we will list out all the other packages that we might need to import into this file.

// So import statement for FMT and then maybe for IO or O's or whatever other packages we want to get access

// to, either from the standard library list of packages which remember we just looked at 2 seconds ago,

// or we could also specify import statements for custom packages like reusable packages that you and I

// have authored ourselves.

// After the package and import statements, we then get down to the body of the file, which is where

// we add in a bunch of logic that actually kind of does something.

// So it'll be a collection of different functions, variable declarations and all that kind of other good stuff as well.

// So in general, we're going to get very used to this very same pattern of code in every last file that we put together.

// So I think that closes it out for the five big questions about our main go file.
