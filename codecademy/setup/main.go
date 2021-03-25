package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	isHeistOn := true
	eludedGuards := rand.Intn(100)
	openedVault := rand.Intn(100)

	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time?")
	}
	if isHeistOn && openedVault >= 50 {
		fmt.Println("Grab and GO!")
	} else if isHeistOn {
		isHeistOn = false
		fmt.Println("Sorry, the vault can't be opened")
	}
	leftSafely := rand.Intn(5)

	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("Failed heist 0")
		case 1:
			isHeistOn = false
			fmt.Println("failed heist 1")
		case 2:
			isHeistOn = false
			fmt.Println("failed heist 2")
		case 3:
			isHeistOn = false
			fmt.Println("failed heist 3")
		default:
			fmt.Println("Start the getaway Car!")
		}
	}
	if isHeistOn {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Println("Amount of money stolen: ", amtStolen)
	}
	fmt.Println(isHeistOn)
}
