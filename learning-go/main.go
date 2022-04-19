package main

import "log"

func main() {
	var whatToSay string
	var saySomeThingElse string
	whatToSay = saySomething("Hello, World!")
	log.Println(whatToSay)

	saySomething("GoodBye")
	log.Println(saySomeThingElse)
	log.Println(saySomething("Finally"))

}

func saySomething(s string) string {
	return s
}
