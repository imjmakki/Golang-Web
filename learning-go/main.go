package main

import "log"

func main() {
	var whatToSay string
	var saySomeThingElse string
	var i int

	whatToSay, _ = saySomething("Hello")
	log.Println(whatToSay)

	saySomeThingElse, _ = saySomething("GoodBye")

	log.Println(saySomeThingElse)
	log.Println(saySomething("Finally"))

	i = 7
	log.Println(i)

}

func saySomething(s string) (string, string) {
	return s, "World"
}
