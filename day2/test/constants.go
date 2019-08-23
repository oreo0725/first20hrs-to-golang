package test

import (
	"fmt"
)

const (
	PI = 3.14
	pp = "pp"
	//Sunday hi sunday
	Sunday = iota // expression list is omitted from rest but not the first one, Initialized to 0
	Monday        //  Now Monday is automatically defined as 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Partyday
	numberOfDays // this constant is not exported but its valued as 7
)

func main() {
	fmt.Println(Monday, Sunday, Tuesday)
}
