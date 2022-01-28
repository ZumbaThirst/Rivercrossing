package main

import (
	"fmt"

	s "github.com/ZumbaThirst/Rivercrossing/states"
)

func main() {
	initializePositions()
	MoveBoat(s.BoatPosition)
	// PutInBoat("Chicken")
	// PutInBoat("Farmer")
	s.ShowWorld()
}

func initializePositions() {
	s.FarmerPosition = "West"
	s.ChickenPosition = "West"
	s.FoxPosition = "West"
	s.CornPosition = "West"
	s.BoatPosition = "West"
	s.PlaceAll()
}

func MoveBoat(item string) (result string) {
	if item == "West" {
		result = "East"
		s.BoatPosition = "East"
		s.PlaceBoat()
	} else if item == "East" {
		result = "West"
		s.BoatPosition = "West"
		s.PlaceBoat()
	}
	return result
}

func PutInBoat(item string) (result string) {
	if item == "Farmer" {
		s.FarmerPosition = "Boat"
		s.PlaceFarmer()
		s.PlaceBoat()
		result = s.Boat
	} else if item == "Chicken" {
		s.ChickenPosition = "Boat"
		s.PlaceChicken()
		s.PlaceBoat()
		result = s.Boat
	} else if item == "Fox" {
		s.FoxPosition = "Boat"
		s.PlaceFox()
		s.PlaceBoat()
		result = s.Boat
	} else if item == "Corn" {
		s.CornPosition = "Boat"
		s.PlaceCorn()
		s.PlaceBoat()
		result = s.Boat
	}
	fmt.Print(result)
	return result
}
