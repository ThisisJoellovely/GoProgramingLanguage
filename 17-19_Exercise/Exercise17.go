package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Exercise 17

● create a program that uses both SEQUENTIAL and CONDITIONAL control flow. Your
program should do the following
○ create a random int between 0 and 250
○ store the value of that int in a variable with the identifier of x
■ func Intn(n int) int ● rand.Intn()
○ print out the name and value of the variable

	use an if statement to do the following

■ if the value is between 0 and 100 print between 0 and 100
■ if the value is between 101 and 200 print between 101 and 200
■ if the value is between 201 and 250 print between 201 and 250
● re: inclusive, exclusive – does rand.Intn()
○ include zero in its output?
○ include 250 in its output?
○ show this in code using the numbers 0 - 3

Exercise 18

	Modify the previous program to use one of these conditional logic statements

○ a switch statement
○ a select statement
● Which of the above conditional logic statements did you choose and why?
curriculum item # 075-hands-on-exercise-20
*/
var x int = rand.Intn(251)
var y int = rand.Intn(251)
var z int = rand.Intn(251)

func main() {

	fmt.Printf("Name of the variable: %T and the value is %v\n", x, x)

	if (x >= 0) && (x <= 100) {
		fmt.Printf("between 0 and 100\n")
	} else if (x >= 101) && (x <= 200) {
		fmt.Printf("between 101 and 200\n")
	} else if (x >= 201) && (x <= 250) {
		fmt.Printf("between 200 and 250\n")
	}

	fmt.Printf("----Exercise 18--------\n\n")

	switch {
	case y > 0 && y < 100:
		fmt.Printf("between 0 and 100\n")
	case y > 101 && y < 200:
		fmt.Printf("between 101 and 200\n")
	case y >= 201 && x <= 250:
		fmt.Printf("between 200 and 250\n")
	}

	fmt.Printf("------Select-Statement-------")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Go-Routine-Channel 1 is on\n"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Go-Routine-Channel 2 is on\n"
	}()
	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- "Go-Routine-Channel 3 is on\n"
	}()

	select {
	case z := <-ch1:
		fmt.Println(z)
	case z := <-ch2:
		fmt.Println(z)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout: No communication ready after 3 second")
	}

}
