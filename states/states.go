package states

import "fmt"

var Boat string = ""

var FarmerPosition = ""
var ChickenPosition = ""
var FoxPosition = ""
var CornPosition = ""
var BoatPosition = ""

var wSlot1 string = ""
var wSlot2 string = ""
var wSlot3 string = ""
var wSlot4 string = ""
var bSlot1 string = "_"
var bSlot2 string = "_"
var eSlot1 string = ""
var eSlot2 string = ""
var eSlot3 string = ""
var eSlot4 string = ""
var wCoastSlot string = ""
var eCoastSlot string = ""

var Status string = ""

func PlaceAll() {
	PlaceFarmer()
	PlaceChicken()
	PlaceFox()
	PlaceCorn()
	PlaceBoat()
}

func PlaceFarmer() {
	if FarmerPosition == "West" {
		goToWestSlot("Farmer")
	} else if FarmerPosition == "Boat" {
		GoToBoatSlot("Farmer")
	} else if FarmerPosition == "East" {
		goToEastSlot("Farmer")
	}
}

func PlaceChicken() {
	if ChickenPosition == "West" {
		goToWestSlot("Chicken")
	} else if ChickenPosition == "Boat" {
		GoToBoatSlot("Chicken")
	} else if ChickenPosition == "East" {
		goToEastSlot("Chicken")
	}
}

func PlaceFox() {
	if FoxPosition == "West" {
		goToWestSlot("Fox")
	} else if FoxPosition == "Boat" {
		GoToBoatSlot("Fox")
	} else if FoxPosition == "East" {
		goToEastSlot("Fox")
	}
}

func PlaceCorn() {
	if CornPosition == "West" {
		goToWestSlot("Corn")
	} else if CornPosition == "Boat" {
		GoToBoatSlot("Corn")
	} else if CornPosition == "East" {
		goToEastSlot("Corn")
	}
}

func PlaceBoat() {
	Boat = "\\_" + bSlot1 + "_" + bSlot2 + "_/"
	if BoatPosition == "West" {
		wCoastSlot = Boat
		eCoastSlot = "_"
	} else if BoatPosition == "East" {
		eCoastSlot = Boat
		wCoastSlot = "_"
	}
}

func goToWestSlot(item string) {
	if wSlot1 == "" {
		wSlot1 = item
	} else if wSlot2 == "" {
		wSlot2 = item
	} else if wSlot3 == "" {
		wSlot3 = item
	} else if wSlot4 == "" {
		wSlot4 = item
	}
}

func GoToBoatSlot(item string) {
	if bSlot1 == "_" {
		bSlot1 = item
	} else if bSlot2 == "_" {
		bSlot2 = item
	}
}

func goToEastSlot(item string) {
	if eSlot1 == "" {
		eSlot1 = item
	} else if eSlot2 == "" {
		eSlot2 = item
	} else if eSlot3 == "" {
		eSlot3 = item
	} else if eSlot4 == "" {
		eSlot4 = item
	}
}

func ShowWorld() {
	Status = wSlot1 + "-" + wSlot2 + "-" + wSlot3 + "-" + wSlot4 + "-W|" + wCoastSlot + "_____________" + eCoastSlot + "|E-" +
		eSlot1 + "-" + eSlot2 + "-" + eSlot3 + "-" + eSlot4
	fmt.Print(Status)
}
