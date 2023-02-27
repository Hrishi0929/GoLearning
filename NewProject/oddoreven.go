package main

import "fmt"

type oddOrEven []int

func (o oddOrEven) print() {
	for _, o := range o {
		if o%2 == 0 {
			fmt.Println(o, " is even")
		} else {
			fmt.Println(o, " is odd")
		}
	}
}
