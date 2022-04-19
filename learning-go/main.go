package main

import "log"

func main() {
	var whatToSay string
	var saySomeThingElse string
	whatToSay = saySaySomething("Hello, World!")
	log.Println(whatToSay)

	saySaySomething("GoodBye")
	log.Println(saySomeThingElse)

}

func saySaySomething(s string) string {
	return s
}
