package main

import (
	"fmt"
)

func main(){
	var (
		conference_name string = "Go conference"
		avaliable_ticket int = 20
		first_name string
		last_name string
		booked_ticket int 
		user_email string
		//bookings  [20]string
		
		)
	const conference_tickets = 20

	// Asking for User Email Address
	fmt.Printf("Enter your Email Address: ")  
	fmt.Scanln(&user_email)

	// Asking for User Fullname
	fmt.Printf("Enter your Fullname: ")
	fmt.Scan(&first_name)
	fmt.Scan(&last_name)         
	fmt.Printf("Hi 🖐🖐🖐🖐 %v %v, Welcome to our %v booking application \n", first_name, last_name,conference_name)
	fmt.Printf("We have total of %v tickes of %v,avaliable tickets are %v. \n", conference_tickets, conference_name,avaliable_ticket)
	fmt.Println("Get 👎🏻 your ticket here to attend!, before it finishes")

	// Getting input from the user
	fmt.Printf("How many tickets would you like to book? ")
	fmt.Scan(&booked_ticket)
	fmt.Printf("%v, you booked %v tickets \n", first_name, booked_ticket)
	avaliable_ticket -= booked_ticket

	// confirmation Email from the user
	fmt.Printf("Thank you for booking %v tickets for %v, you will receive a confirmation Email at %v soon. \n",booked_ticket,conference_name,user_email)
	fmt.Printf("we have %v tickets avaliable \n",avaliable_ticket)
}




