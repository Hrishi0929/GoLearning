package main

func main() {
	zeroToTen := oddOrEven{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	zeroToTen.print()

}

// func main() {
// 	var zeroToTen []int
// 	for i := 0; i <= 10; i++ {
// 		zeroToTen = append(zeroToTen, i)
// 	}

// 	for _, n := range zeroToTen {
// 		if n%2 == 0 {
// 			fmt.Println(n, " is Even")
// 		} else {
// 			fmt.Println(n, " is Odd")
// 		}
// 	}
// }
