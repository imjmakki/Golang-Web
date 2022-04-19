package main

import "log"

func main() {
	var whatToSay string
	var saySomeThingElse string
	var i int

	whatToSay = saySomething("Hello, World!")
	log.Println(whatToSay)

	saySomething("GoodBye")

	log.Println(saySomeThingElse)
	log.Println(saySomething("Finally"))

	i = 7
	log.Println(i)

}

func saySomething(s string) string {
	return s
}
