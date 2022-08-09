package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = conferenceTickets
	// var bookings = [50]string{"koybi", "nana", "nicole"}
	// var bookings := []string
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")

	// ask user for their name

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Print("enter your firstname : ")
		fmt.Scan(&firstName)

		if firstName == "0" {
			break
		}

		fmt.Print("enter your lastname : ")
		fmt.Scan(&lastName)

		fmt.Print("enter your email : ")
		fmt.Scan(&email)

		fmt.Print("enter number of tickets that you want : ")
		fmt.Scan(&userTickets)

		isValidNAme := len(firstName) >= 5 && len(lastName) >= 5
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if !(isValidEmail && isValidNAme && isValidTicketNumber) {
			break
		}

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets

			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}

			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of booking are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out, come back next time !")
				break
			}
		} else {
			fmt.Printf("We only have %v tikets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			continue
		}
	}
}
