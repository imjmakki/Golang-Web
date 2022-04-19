package main

import "log"

func main() {
	var whatToSay string
	whatToSay = saySaySomething("Hello, World!")
	log.Println(whatToSay)
}

func saySaySomething(s string) string {
	return s
}
