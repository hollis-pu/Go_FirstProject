package main

import (
	"fmt"
	"time"
)

/**
* Description:
	select使用的案例
* @Author Hollis
* @Create 2023/9/10 13:31
*/

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		for i := 1; i <= 5; i++ {
			time.Sleep(time.Second)
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		fruits := []string{"apple", "banana", "cherry", "date", "grape"}
		for _, fruit := range fruits {
			time.Sleep(time.Second)
			ch2 <- fruit
		}
		close(ch2)
	}()

	for {
		select {
		case num, ok := <-ch1:
			if !ok {
				fmt.Println("ch1 is closed.")
				ch1 = nil // Set to nil to prevent further reads
			} else {
				fmt.Println("Received from ch1:", num)
			}
		case fruit, ok := <-ch2:
			if !ok {
				fmt.Println("ch2 is closed.")
				ch2 = nil // Set to nil to prevent further reads
			} else {
				fmt.Println("Received from ch2:", fruit)
			}
		}

		// Check if both channels are closed to exit the loop
		if ch1 == nil && ch2 == nil {
			break
		}
	}
}
