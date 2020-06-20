package main

import (
	"fmt"

	"github.com/kanit147/finalexam/customer"
)

func main() {
	fmt.Println("### welcome to customer management ###")
	r := customer.SetupRouter()
	r.Run(":2019")
}
