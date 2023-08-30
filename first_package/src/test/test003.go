package main

import "fmt"
/**
* Description:
* @Author Hollis
* @Create 2023/8/29 10:14
 */


func main() {
	source := [5]int{1, 2, 3, 4, 5}
	destination := [5]int{}

	numCopied := copy(destination[:], source[:]) // Using slices to copy arrays

	fmt.Println("Copied", numCopied, "elements")
	fmt.Println("Source:", source)
	fmt.Println("Destination:", destination)
}

