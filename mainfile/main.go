package main

import (
	"fmt"
	"log"
	"project3/calendar"
)
func main(){
	date := calendar.Event{}
	// date.Title = "mommmy birthday"
	err := date.SetEvent("Mommy Birthday and my cousin from uk")
	 if err != nil{
		log.Fatal(err)
	}
	err = date.SetYear(2014)
	if err != nil{
		log.Fatal(err)
	}
	err = date.SetMonth(12)
	if err != nil{
		log.Fatal(err)
	}
	err = date.SetDay(14)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(date)
	fmt.Println(date.Title())
	fmt.Println(date.Year())
	fmt.Println(date.Month())
	fmt.Println(date.Day())
	// // var now time.Time = time.Now()
	// fmt.Println(time.Now()) // prints the current time
	// fmt.Println(now.Date())// prints the current date
	


}