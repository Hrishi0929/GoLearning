package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	response, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error Occured:", err)
		os.Exit(1)
	}
	// fmt.Println(response)

	// // 	So this make function right here is a built in function in the language that takes a type of a slice.
	// // And then as a second argument, this is the number of elements or empty spaces that we want the slice to be initialized with.
	// // So remember we spoke a lot about slices just a little bit ago and how they work with underlying arrays.
	// // So yes, a byte slice can grow and shrink, but if we want to, we can create a byte slice with n number of empty elements inside of it.
	// // And so this line right here specifically says, give me an empty byte slice that has space for 99,999 elements.

	// byteSlice := make([]byte, 99999)

	// // 	So essentially the read function is just going to read data into the byte slice until the bite size is full.
	// // And so if we initialize a bite slice with like zero elements inside of it, the read function is going to say, Oh, this thing is already full.
	// // I can't read any more data into the byte slice, so I'm just going to quit early.

	// // So now right underneath this, we'll look at the response struct.
	// // We'll get the body property off of there, which is the value that implements the reader interface.
	// // And because it implements the reader interface, we know it has access to the read function and even
	// // the autocomplete is kicking in here and saying, Hey, this thing has a read function, so I'm going
	// // to click on read, we'll call it, and then we're going to pass in that byte slice right there that we just created.
	// // So now the response body should take the byte slice take it's HTML that is contained inside the body and read that data into the byte slice.
	// response.Body.Read(byteSlice)
	// fmt.Println(string(byteSlice))
	lw := logWriter{}
	// why we used io.copy here refer to the link below
	// https://www.udemy.com/course/go-the-complete-developers-guide/learn/lecture/7797408#overview
	// io.Copy(os.Stdout, response.Body)
	io.Copy(lw, response.Body)
}

func (logWriter) Write(byteSlice []byte) (int, error) {
	fmt.Println(string(byteSlice))
	fmt.Println("Just wrote this many bytes:", len(byteSlice))
	return len(byteSlice), nil
}
