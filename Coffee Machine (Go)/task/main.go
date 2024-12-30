package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getNumber(prompt string) int {
	fmt.Print(prompt)
	var n int
	_, err := fmt.Scan(&n)
	for err == nil {
		fmt.Println("Enter a digit")
		fmt.Print(prompt)
		_, err = fmt.Scan(&n)
	}
	return n
}

func howManyCupsMade(water, milk, coffeeBeans int) int {
	var n int
	amountOfWater := 200
	amountOfMIlk := 50
	amountOfCoffeeBeans := 15
	for water >= amountOfWater && milk >= amountOfMIlk && coffeeBeans >= amountOfCoffeeBeans {
		water -= amountOfWater
		milk -= amountOfMIlk
		coffeeBeans -= amountOfCoffeeBeans
		n++
	}
	return n
}

func displaySupplies(water, milk, coffeeBeans, numberOfCups, amountOfMoney int) {
	fmt.Println("\nThe coffee machine has:")
	fmt.Printf("%d ml of water\n", water)
	fmt.Printf("%d ml of milk\n", milk)
	fmt.Printf("%d g of coffee beans\n", coffeeBeans)
	fmt.Printf("%d disposable cups\n", numberOfCups)
	fmt.Printf("$%d of money\n\n", amountOfMoney)
}

func makeCoffee(water, milk, coffeeBeans *int, typeOfCoffee int) int {
	var price int
	switch typeOfCoffee {
	case 1:
		// make espresso
		if *water < 250 {
			fmt.Print("Sorry, not enough water!\n\n")
			return -1
		}
		if *coffeeBeans < 16 {
			fmt.Print("Sorry, not enough coffee beans!\n\n")
			return -1
		}
		*water -= 250
		*coffeeBeans -= 16
		price = 4 // costs $4
	case 2:
		// make latte
		if *water < 350 {
			fmt.Print("Sorry, not enough water!\n\n")
			return -1
		}
		if *milk < 75 {
			fmt.Print("Sorry, not enough milk!\n\n")
			return -1
		}
		if *coffeeBeans < 20 {
			fmt.Print("Sorry, not enough coffee beans!\n\n")
			return -1
		}
		*water -= 350
		*milk -= 75
		*coffeeBeans -= 20
		price = 7 // costs $7
	case 3:
		// make cappuccino
		if *water < 200 {
			fmt.Print("Sorry, not enough water!\n\n")
			return -1
		}
		if *milk < 100 {
			fmt.Print("Sorry, not enough milk!\n\n")
			return -1
		}
		if *coffeeBeans < 12 {
			fmt.Print("Sorry, not enough coffee beans!\n\n")
			return -1
		}
		*water -= 200
		*milk -= 100
		*coffeeBeans -= 12
		price = 6 // costs $4
	}
	fmt.Print("I have enough resources, making you a coffee!\n\n")
	return price
}

func chooseAction() {
	// initial ingredients
	amountOfWater := 400
	amountOfMilk := 540
	amountOfCoffeeBeans := 120
	numberOfCups := 9
	amountOfMoney := 550

	var action string
	for {
		fmt.Print("Write action (buy, fill, take, remaining, exit): ")
		fmt.Scan(&action)
		action = strings.TrimSpace(action) // remove leading and trailing whitespace
		fmt.Print("\n")
		switch action {
		case "buy":
			fmt.Print("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu: ")
			var input string
			fmt.Scanf("%s", &input)
			fmt.Print("\n")
			input = strings.TrimSpace(input)
			if input == "back" {
				continue
			}
			coffeeType, _ := strconv.Atoi(input) // convert to int
			number := makeCoffee(&amountOfWater, &amountOfMilk, &amountOfCoffeeBeans, coffeeType)
			if number != -1 {
				amountOfMoney += number
				numberOfCups--
			}
		case "fill":
			water := getNumber("Write how many ml of water you want to add: ")
			fmt.Print("\n")
			milk := getNumber("Write how many ml of milk you want to add: ")
			fmt.Print("\n")
			coffeeBeans := getNumber("Write how many grams of coffee beans you want to add: ")
			fmt.Print("\n")
			cups := getNumber("Write how many disposable cups you want to add: ")
			fmt.Print("\n")
			amountOfWater += water
			amountOfMilk += milk
			amountOfCoffeeBeans += coffeeBeans
			numberOfCups += cups
		case "take":
			fmt.Printf("I gave you $%d\n\n", amountOfMoney)
			amountOfMoney = 0
		case "remaining":
			displaySupplies(amountOfWater, amountOfMilk, amountOfCoffeeBeans, numberOfCups, amountOfMoney)
		case "exit":
			return
		}
	}

}

func main() {
	chooseAction()
}
