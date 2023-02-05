package main

import "fmt"

func main() {
	var publisher, writer, artist, title string
	title = "Mr. GoToSleep"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampdon"
	publisher = "DizzyBooks Publishing Inc."
	var year, pageNumber int16
	year = 1997
	pageNumber = 14
	var grade float32
	grade = 6.5
	fmt.Println(title, "written by", writer, "drawn by", artist)
	fmt.Println(publisher, "written by", year, "drawn by", pageNumber, "and", grade)
}
