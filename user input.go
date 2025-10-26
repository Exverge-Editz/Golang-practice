package main

import "fmt"

func main() {
	forename := ``
	surname := ``
	fmt.Println(`what is your forename?: `)
	fmt.Scan(&forename)
	fmt.Println(`what is your surname?: `)
	fmt.Scan(&surname)
	fmt.Println(`hello`, forename, surname, `.`)
}
